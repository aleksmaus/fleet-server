// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package config

// Instrumentation configures APM Tracing for the `fleet-server`.
type Instrumentation struct {
	Environment string   `config:"environment"`
	APIKey      string   `config:"api_key"`
	SecretToken string   `config:"secret_token"`
	Hosts       []string `config:"hosts"`
	Enabled     bool     `config:"enabled"`
}