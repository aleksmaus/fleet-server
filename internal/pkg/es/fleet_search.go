// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated from specification version 7.x: DO NOT EDIT

// This is a copy of api.search.go file from go-elasticsearch library
// It was modified for /_fleet/_fleet_search experimental API,
// implemented by the custom fleet plugin https://github.com/elastic/elasticsearch/pull/73134
// This file can be removed and replaced with the official client library wrapper once it is available

package es

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/fleet-server/v7/internal/pkg/sqn"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// ----- API Definition -------------------------------------------------------

// FleetSearch returns results matching a query.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/search-search.html.
//
type FleetSearch func(o ...func(*FleetSearchRequest)) (*Response, error)

// FleetSearchRequest configures the FleetSearch API request.
//
type FleetSearchRequest struct {
	Index        []string
	DocumentType []string

	Body io.Reader

	AllowNoIndices             *bool
	AllowPartialSearchResults  *bool
	Analyzer                   string
	AnalyzeWildcard            *bool
	BatchedReduceSize          *int
	CcsMinimizeRoundtrips      *bool
	DefaultOperator            string
	Df                         string
	DocvalueFields             []string
	ExpandWildcards            string
	Explain                    *bool
	From                       *int
	IgnoreThrottled            *bool
	IgnoreUnavailable          *bool
	Lenient                    *bool
	MaxConcurrentShardRequests *int
	MinCompatibleShardNode     string
	Preference                 string
	PreFilterShardSize         *int
	Query                      string
	RequestCache               *bool
	RestTotalHitsAsInt         *bool
	Routing                    []string
	Scroll                     time.Duration
	SearchType                 string
	SeqNoPrimaryTerm           *bool
	Size                       *int
	Sort                       []string
	Source                     []string
	SourceExcludes             []string
	SourceIncludes             []string
	Stats                      []string
	StoredFields               []string
	SuggestField               string
	SuggestMode                string
	SuggestSize                *int
	SuggestText                string
	TerminateAfter             *int
	Timeout                    time.Duration
	TrackScores                *bool
	TrackTotalHits             interface{}
	TypedKeys                  *bool
	Version                    *bool
	WaitForCheckpoints         []int64

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r FleetSearchRequest) Do(ctx context.Context, transport esapi.Transport) (*esapi.Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len(strings.Join(r.Index, ",")) + 1 + len(strings.Join(r.DocumentType, ",")) + 1 + len("_fleet/_fleet_search"))
	if len(r.Index) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.Index, ","))
	}
	if len(r.DocumentType) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.DocumentType, ","))
	}
	path.WriteString("/")
	path.WriteString("_fleet/_fleet_search")

	params = make(map[string]string)

	if len(r.WaitForCheckpoints) > 0 {
		params["wait_for_checkpoints"] = sqn.SeqNo(r.WaitForCheckpoints).String()
	}

	if r.AllowNoIndices != nil {
		params["allow_no_indices"] = strconv.FormatBool(*r.AllowNoIndices)
	}

	if r.AllowPartialSearchResults != nil {
		params["allow_partial_search_results"] = strconv.FormatBool(*r.AllowPartialSearchResults)
	}

	if r.Analyzer != "" {
		params["analyzer"] = r.Analyzer
	}

	if r.AnalyzeWildcard != nil {
		params["analyze_wildcard"] = strconv.FormatBool(*r.AnalyzeWildcard)
	}

	if r.BatchedReduceSize != nil {
		params["batched_reduce_size"] = strconv.FormatInt(int64(*r.BatchedReduceSize), 10)
	}

	if r.CcsMinimizeRoundtrips != nil {
		params["ccs_minimize_roundtrips"] = strconv.FormatBool(*r.CcsMinimizeRoundtrips)
	}

	if r.DefaultOperator != "" {
		params["default_operator"] = r.DefaultOperator
	}

	if r.Df != "" {
		params["df"] = r.Df
	}

	if len(r.DocvalueFields) > 0 {
		params["docvalue_fields"] = strings.Join(r.DocvalueFields, ",")
	}

	if r.ExpandWildcards != "" {
		params["expand_wildcards"] = r.ExpandWildcards
	}

	if r.Explain != nil {
		params["explain"] = strconv.FormatBool(*r.Explain)
	}

	if r.From != nil {
		params["from"] = strconv.FormatInt(int64(*r.From), 10)
	}

	if r.IgnoreThrottled != nil {
		params["ignore_throttled"] = strconv.FormatBool(*r.IgnoreThrottled)
	}

	if r.IgnoreUnavailable != nil {
		params["ignore_unavailable"] = strconv.FormatBool(*r.IgnoreUnavailable)
	}

	if r.Lenient != nil {
		params["lenient"] = strconv.FormatBool(*r.Lenient)
	}

	if r.MaxConcurrentShardRequests != nil {
		params["max_concurrent_shard_requests"] = strconv.FormatInt(int64(*r.MaxConcurrentShardRequests), 10)
	}

	if r.MinCompatibleShardNode != "" {
		params["min_compatible_shard_node"] = r.MinCompatibleShardNode
	}

	if r.Preference != "" {
		params["preference"] = r.Preference
	}

	if r.PreFilterShardSize != nil {
		params["pre_filter_shard_size"] = strconv.FormatInt(int64(*r.PreFilterShardSize), 10)
	}

	if r.Query != "" {
		params["q"] = r.Query
	}

	if r.RequestCache != nil {
		params["request_cache"] = strconv.FormatBool(*r.RequestCache)
	}

	if r.RestTotalHitsAsInt != nil {
		params["rest_total_hits_as_int"] = strconv.FormatBool(*r.RestTotalHitsAsInt)
	}

	if len(r.Routing) > 0 {
		params["routing"] = strings.Join(r.Routing, ",")
	}

	if r.Scroll != 0 {
		params["scroll"] = formatDuration(r.Scroll)
	}

	if r.SearchType != "" {
		params["search_type"] = r.SearchType
	}

	if r.SeqNoPrimaryTerm != nil {
		params["seq_no_primary_term"] = strconv.FormatBool(*r.SeqNoPrimaryTerm)
	}

	if r.Size != nil {
		params["size"] = strconv.FormatInt(int64(*r.Size), 10)
	}

	if len(r.Sort) > 0 {
		params["sort"] = strings.Join(r.Sort, ",")
	}

	if len(r.Source) > 0 {
		params["_source"] = strings.Join(r.Source, ",")
	}

	if len(r.SourceExcludes) > 0 {
		params["_source_excludes"] = strings.Join(r.SourceExcludes, ",")
	}

	if len(r.SourceIncludes) > 0 {
		params["_source_includes"] = strings.Join(r.SourceIncludes, ",")
	}

	if len(r.Stats) > 0 {
		params["stats"] = strings.Join(r.Stats, ",")
	}

	if len(r.StoredFields) > 0 {
		params["stored_fields"] = strings.Join(r.StoredFields, ",")
	}

	if r.SuggestField != "" {
		params["suggest_field"] = r.SuggestField
	}

	if r.SuggestMode != "" {
		params["suggest_mode"] = r.SuggestMode
	}

	if r.SuggestSize != nil {
		params["suggest_size"] = strconv.FormatInt(int64(*r.SuggestSize), 10)
	}

	if r.SuggestText != "" {
		params["suggest_text"] = r.SuggestText
	}

	if r.TerminateAfter != nil {
		params["terminate_after"] = strconv.FormatInt(int64(*r.TerminateAfter), 10)
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
	}

	if r.TrackScores != nil {
		params["track_scores"] = strconv.FormatBool(*r.TrackScores)
	}

	if r.TrackTotalHits != nil {
		params["track_total_hits"] = fmt.Sprintf("%v", r.TrackTotalHits)
	}

	if r.TypedKeys != nil {
		params["typed_keys"] = strconv.FormatBool(*r.TypedKeys)
	}

	if r.Version != nil {
		params["version"] = strconv.FormatBool(*r.Version)
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), r.Body)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := esapi.Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f FleetSearch) WithContext(v context.Context) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.ctx = v
	}
}

// WithBody - The search definition using the Query DSL.
//
func (f FleetSearch) WithBody(v io.Reader) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Body = v
	}
}

// WithIndex - a list of index names to search; use _all to perform the operation on all indices.
//
func (f FleetSearch) WithIndex(v ...string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Index = v
	}
}

// WithDocumentType - a list of document types to search; leave empty to perform the operation on all types.
//
func (f FleetSearch) WithDocumentType(v ...string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.DocumentType = v
	}
}

// WithAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (this includes `_all` string or when no indices have been specified).
//
func (f FleetSearch) WithAllowNoIndices(v bool) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.AllowNoIndices = &v
	}
}

// WithAllowPartialSearchResults - indicate if an error should be returned if there is a partial search failure or timeout.
//
func (f FleetSearch) WithAllowPartialSearchResults(v bool) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.AllowPartialSearchResults = &v
	}
}

// WithAnalyzer - the analyzer to use for the query string.
//
func (f FleetSearch) WithAnalyzer(v string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Analyzer = v
	}
}

// WithAnalyzeWildcard - specify whether wildcard and prefix queries should be analyzed (default: false).
//
func (f FleetSearch) WithAnalyzeWildcard(v bool) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.AnalyzeWildcard = &v
	}
}

// WithBatchedReduceSize - the number of shard results that should be reduced at once on the coordinating node. this value should be used as a protection mechanism to reduce the memory overhead per search request if the potential number of shards in the request can be large..
//
func (f FleetSearch) WithBatchedReduceSize(v int) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.BatchedReduceSize = &v
	}
}

// WithCcsMinimizeRoundtrips - indicates whether network round-trips should be minimized as part of cross-cluster search requests execution.
//
func (f FleetSearch) WithCcsMinimizeRoundtrips(v bool) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.CcsMinimizeRoundtrips = &v
	}
}

// WithDefaultOperator - the default operator for query string query (and or or).
//
func (f FleetSearch) WithDefaultOperator(v string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.DefaultOperator = v
	}
}

// WithDf - the field to use as default where no field prefix is given in the query string.
//
func (f FleetSearch) WithDf(v string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Df = v
	}
}

// WithDocvalueFields - a list of fields to return as the docvalue representation of a field for each hit.
//
func (f FleetSearch) WithDocvalueFields(v ...string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.DocvalueFields = v
	}
}

// WithExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both..
//
func (f FleetSearch) WithExpandWildcards(v string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.ExpandWildcards = v
	}
}

// WithExplain - specify whether to return detailed information about score computation as part of a hit.
//
func (f FleetSearch) WithExplain(v bool) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Explain = &v
	}
}

// WithFrom - starting offset (default: 0).
//
func (f FleetSearch) WithFrom(v int) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.From = &v
	}
}

// WithIgnoreThrottled - whether specified concrete, expanded or aliased indices should be ignored when throttled.
//
func (f FleetSearch) WithIgnoreThrottled(v bool) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.IgnoreThrottled = &v
	}
}

// WithIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
//
func (f FleetSearch) WithIgnoreUnavailable(v bool) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.IgnoreUnavailable = &v
	}
}

// WithLenient - specify whether format-based query failures (such as providing text to a numeric field) should be ignored.
//
func (f FleetSearch) WithLenient(v bool) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Lenient = &v
	}
}

// WithMaxConcurrentShardRequests - the number of concurrent shard requests per node this search executes concurrently. this value should be used to limit the impact of the search on the cluster in order to limit the number of concurrent shard requests.
//
func (f FleetSearch) WithMaxConcurrentShardRequests(v int) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.MaxConcurrentShardRequests = &v
	}
}

// WithMinCompatibleShardNode - the minimum compatible version that all shards involved in search should have for this request to be successful.
//
func (f FleetSearch) WithMinCompatibleShardNode(v string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.MinCompatibleShardNode = v
	}
}

// WithPreference - specify the node or shard the operation should be performed on (default: random).
//
func (f FleetSearch) WithPreference(v string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Preference = v
	}
}

// WithPreFilterShardSize - a threshold that enforces a pre-filter roundtrip to prefilter search shards based on query rewriting if the number of shards the search request expands to exceeds the threshold. this filter roundtrip can limit the number of shards significantly if for instance a shard can not match any documents based on its rewrite method ie. if date filters are mandatory to match but the shard bounds and the query are disjoint..
//
func (f FleetSearch) WithPreFilterShardSize(v int) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.PreFilterShardSize = &v
	}
}

// WithQuery - query in the lucene query string syntax.
//
func (f FleetSearch) WithQuery(v string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Query = v
	}
}

// WithRequestCache - specify if request cache should be used for this request or not, defaults to index level setting.
//
func (f FleetSearch) WithRequestCache(v bool) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.RequestCache = &v
	}
}

// WithRestTotalHitsAsInt - indicates whether hits.total should be rendered as an integer or an object in the rest search response.
//
func (f FleetSearch) WithRestTotalHitsAsInt(v bool) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.RestTotalHitsAsInt = &v
	}
}

// WithRouting - a list of specific routing values.
//
func (f FleetSearch) WithRouting(v ...string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Routing = v
	}
}

// WithScroll - specify how long a consistent view of the index should be maintained for scrolled search.
//
func (f FleetSearch) WithScroll(v time.Duration) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Scroll = v
	}
}

// WithSearchType - search operation type.
//
func (f FleetSearch) WithSearchType(v string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.SearchType = v
	}
}

// WithSeqNoPrimaryTerm - specify whether to return sequence number and primary term of the last modification of each hit.
//
func (f FleetSearch) WithSeqNoPrimaryTerm(v bool) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.SeqNoPrimaryTerm = &v
	}
}

// WithSize - number of hits to return (default: 10).
//
func (f FleetSearch) WithSize(v int) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Size = &v
	}
}

// WithSort - a list of <field>:<direction> pairs.
//
func (f FleetSearch) WithSort(v ...string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Sort = v
	}
}

// WithSource - true or false to return the _source field or not, or a list of fields to return.
//
func (f FleetSearch) WithSource(v ...string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Source = v
	}
}

// WithSourceExcludes - a list of fields to exclude from the returned _source field.
//
func (f FleetSearch) WithSourceExcludes(v ...string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.SourceExcludes = v
	}
}

// WithSourceIncludes - a list of fields to extract and return from the _source field.
//
func (f FleetSearch) WithSourceIncludes(v ...string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.SourceIncludes = v
	}
}

// WithStats - specific 'tag' of the request for logging and statistical purposes.
//
func (f FleetSearch) WithStats(v ...string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Stats = v
	}
}

// WithStoredFields - a list of stored fields to return as part of a hit.
//
func (f FleetSearch) WithStoredFields(v ...string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.StoredFields = v
	}
}

// WithSuggestField - specify which field to use for suggestions.
//
func (f FleetSearch) WithSuggestField(v string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.SuggestField = v
	}
}

// WithSuggestMode - specify suggest mode.
//
func (f FleetSearch) WithSuggestMode(v string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.SuggestMode = v
	}
}

// WithSuggestSize - how many suggestions to return in response.
//
func (f FleetSearch) WithSuggestSize(v int) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.SuggestSize = &v
	}
}

// WithSuggestText - the source text for which the suggestions should be returned.
//
func (f FleetSearch) WithSuggestText(v string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.SuggestText = v
	}
}

// WithTerminateAfter - the maximum number of documents to collect for each shard, upon reaching which the query execution will terminate early..
//
func (f FleetSearch) WithTerminateAfter(v int) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.TerminateAfter = &v
	}
}

// WithTimeout - explicit operation timeout.
//
func (f FleetSearch) WithTimeout(v time.Duration) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Timeout = v
	}
}

// WithTrackScores - whether to calculate and return scores even if they are not used for sorting.
//
func (f FleetSearch) WithTrackScores(v bool) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.TrackScores = &v
	}
}

// WithTrackTotalHits - indicate if the number of documents that match the query should be tracked.
//
func (f FleetSearch) WithTrackTotalHits(v interface{}) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.TrackTotalHits = v
	}
}

// WithTypedKeys - specify whether aggregation and suggester names should be prefixed by their respective types in the response.
//
func (f FleetSearch) WithTypedKeys(v bool) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.TypedKeys = &v
	}
}

// WithVersion - specify whether to return document version as part of a hit.
//
func (f FleetSearch) WithVersion(v bool) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Version = &v
	}
}

// WithWaitForCheckpoints - specify the list of checkpoints to wait for, https://github.com/elastic/elasticsearch/pull/73134
//
func (f FleetSearch) WithWaitForCheckpoints(checkpoints []int64) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.WaitForCheckpoints = checkpoints
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f FleetSearch) WithPretty() func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f FleetSearch) WithHuman() func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f FleetSearch) WithErrorTrace() func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f FleetSearch) WithFilterPath(v ...string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f FleetSearch) WithHeader(h map[string]string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
//
func (f FleetSearch) WithOpaqueID(s string) func(*FleetSearchRequest) {
	return func(r *FleetSearchRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
