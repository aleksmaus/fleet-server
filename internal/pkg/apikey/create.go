// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package apikey

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

const fleetAgent = "fleet-agent"

type Type int

const (
	TypeAccess Type = iota
	TypeOutput
)

func (t Type) String() string {
	return []string{"access", "output"}[t]
}

type Metadata struct {
	Application string `json:"application"`
	AgentId     string `json:"agent_id"`
	Type        string `json:"type"`
}

func Create(ctx context.Context, client *elasticsearch.Client, keyType Type, agentId, name, ttl string, roles []byte) (*ApiKey, error) {

	payload := struct {
		Name       string          `json:"name,omitempty"`
		Expiration string          `json:"expiration,omitempty"`
		Roles      json.RawMessage `json:"role_descriptors,omitempty"`
		Metadata   Metadata        `json:"metadata"`
	}{
		Name:       name,
		Expiration: ttl,
		Roles:      roles,
		Metadata: Metadata{
			Application: fleetAgent,
			AgentId:     agentId,
			Type:        keyType.String(),
		},
	}

	body, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

	opts := []func(*esapi.SecurityCreateAPIKeyRequest){
		client.Security.CreateAPIKey.WithContext(ctx),
		client.Security.CreateAPIKey.WithRefresh("true"),
	}

	res, err := client.Security.CreateAPIKey(
		bytes.NewReader(body),
		opts...,
	)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("fail CreateAPIKey: %s", res.String())
	}

	type APIKeyResponse struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		Expiration uint64 `json:"expiration"`
		ApiKey     string `json:"api_key"`
	}

	var resp APIKeyResponse
	d := json.NewDecoder(res.Body)
	if err = d.Decode(&resp); err != nil {
		return nil, err
	}

	key := ApiKey{
		Id:  resp.Id,
		Key: resp.ApiKey,
	}

	return &key, err
}
