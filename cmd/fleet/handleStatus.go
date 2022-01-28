// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package fleet

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/elastic/fleet-server/v7/internal/pkg/apikey"
	"github.com/elastic/fleet-server/v7/internal/pkg/bulk"
	"github.com/elastic/fleet-server/v7/internal/pkg/cache"
	"github.com/elastic/fleet-server/v7/internal/pkg/config"
	"github.com/elastic/fleet-server/v7/internal/pkg/limit"
	"github.com/elastic/fleet-server/v7/internal/pkg/logger"
	"github.com/julienschmidt/httprouter"

	"github.com/elastic/elastic-agent-client/v7/pkg/proto"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	kStatusMod = "status"
)

type AuthFunc func(*http.Request) (*apikey.ApiKey, error)

type StatusT struct {
	cfg    *config.Server
	limit  *limit.Limiter
	bulk   bulk.Bulk
	cache  cache.Cache
	authfn AuthFunc
}

type OptFunc func(*StatusT)

func NewStatusT(cfg *config.Server, bulker bulk.Bulk, cache cache.Cache, opts ...OptFunc) *StatusT {
	log.Info().
		Interface("limits", cfg.Limits.StatusLimit).
		Msg("Setting config status_limits")

	st := &StatusT{
		cfg:   cfg,
		bulk:  bulker,
		cache: cache,
		limit: limit.NewLimiter(&cfg.Limits.StatusLimit),
	}
	st.authfn = st.authenticate

	for _, opt := range opts {
		opt(st)
	}
	return st
}

func withAuthFunc(authfn AuthFunc) OptFunc {
	return func(st *StatusT) {
		if authfn != nil {
			st.authfn = authfn
		}
	}
}

func (st StatusT) authenticate(r *http.Request) (*apikey.ApiKey, error) {
	return authApiKey(r, st.bulk, st.cache)
}

func (st StatusT) handleStatus(zlog *zerolog.Logger, r *http.Request, rt *Router) (resp StatusResponse, status proto.StateObserved_Status, err error) {
	limitF, err := st.limit.Acquire()
	// When failing to acquire a limiter send an error response.
	if err != nil {
		return
	}
	defer limitF()

	authed := true
	if _, aerr := st.authfn(r); aerr != nil {
		log.Debug().Err(aerr).Msg("unauthenticated status request, return short status response only")
		authed = false
	}

	status = rt.sm.Status()
	resp = StatusResponse{
		Name:   kServiceName,
		Status: status.String(),
	}

	if authed {
		resp.Version = &StatusResponseVersion{
			Number:    rt.bi.Version,
			BuildHash: rt.bi.Commit,
			BuildTime: rt.bi.BuildTime.Format(time.RFC3339),
		}
	}

	return resp, status, nil

}

func (rt Router) handleStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	start := time.Now()

	dfunc := cntStatus.IncStart()
	defer dfunc()

	reqId := r.Header.Get(logger.HeaderRequestID)

	zlog := log.With().
		Str(EcsHttpRequestId, reqId).
		Str("mod", kStatusMod).
		Logger()

	resp, status, err := rt.st.handleStatus(&zlog, r, &rt)
	if err != nil {
		cntStatus.IncError(err)
		resp := NewErrorResp(err)

		zlog.WithLevel(resp.Level).
			Err(err).
			Int(EcsHttpResponseCode, resp.StatusCode).
			Int64(EcsEventDuration, time.Since(start).Nanoseconds()).
			Msg("fail status")

		if rerr := resp.Write(w); rerr != nil {
			zlog.Error().Err(rerr).Msg("fail writing error response")
		}
		return
	}

	data, err := json.Marshal(&resp)
	if err != nil {
		code := http.StatusInternalServerError
		zlog.Error().Err(err).Int(EcsHttpResponseCode, code).
			Int64(EcsEventDuration, time.Since(start).Nanoseconds()).Msg("fail status")
		http.Error(w, "", code)
		return
	}

	code := http.StatusServiceUnavailable
	if status == proto.StateObserved_DEGRADED || status == proto.StateObserved_HEALTHY {
		code = http.StatusOK
	}
	w.WriteHeader(code)

	var nWritten int
	if nWritten, err = w.Write(data); err != nil {
		if err != context.Canceled {
			zlog.Error().Err(err).Int(EcsHttpResponseCode, code).
				Int64(EcsEventDuration, time.Since(start).Nanoseconds()).Msg("fail status")
		}
	}

	cntStatus.bodyOut.Add(uint64(nWritten))
	zlog.Debug().
		Int(EcsHttpResponseBodyBytes, nWritten).
		Int64(EcsEventDuration, time.Since(start).Nanoseconds()).
		Msg("ok status")
}
