// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package dl

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/elastic/fleet-server/v7/internal/pkg/bulk"
	"github.com/elastic/fleet-server/v7/internal/pkg/model"

	"github.com/elastic/fleet-server/v7/internal/pkg/dsl"
)

var (
	tmplQueryLatestPolicies = prepareQueryLatestPolicies()

	queryPolicyByID = preparePolicyFindByID()
)

var ErrPolicyLeaderNotFound = errors.New("policy has no leader")
var ErrMissingAggregations = errors.New("missing expected aggregation result")

func prepareQueryLatestPolicies() []byte {
	root := dsl.NewRoot()
	root.Size(0)
	policyId := root.Aggs().Agg(FieldPolicyId)
	policyId.Terms("field", FieldPolicyId, nil)
	revisionIdx := policyId.Aggs().Agg(FieldRevisionIdx).TopHits()
	revisionIdx.Size(1)
	rSort := revisionIdx.Sort()
	rSort.SortOrder(FieldRevisionIdx, dsl.SortDescend)
	rSort.SortOrder(FieldCoordinatorIdx, dsl.SortDescend)
	return root.MustMarshalJSON()
}

func preparePolicyFindByID() *dsl.Tmpl {
	tmpl := dsl.NewTmpl()
	root := dsl.NewRoot()

	root.Size(1)
	root.Query().Bool().Filter().Term(FieldPolicyId, tmpl.Bind(FieldPolicyId), nil)
	sort := root.Sort()
	sort.SortOrder(FieldRevisionIdx, dsl.SortDescend)
	sort.SortOrder(FieldCoordinatorIdx, dsl.SortDescend)

	tmpl.MustResolve(root)
	return tmpl
}

// QueryLatestPolices gets the latest revision for a policy
func QueryLatestPolicies(ctx context.Context, bulker bulk.Bulk, opt ...Option) ([]model.Policy, error) {
	o := newOption(FleetPolicies, opt...)
	res, err := bulker.Search(ctx, []string{o.indexName}, tmplQueryLatestPolicies)
	if err != nil {
		return nil, err
	}

	policyId, ok := res.Aggregations[FieldPolicyId]
	if !ok {
		return nil, ErrMissingAggregations
	}
	if len(policyId.Buckets) == 0 {
		return []model.Policy{}, nil
	}
	policies := make([]model.Policy, len(policyId.Buckets))
	for i, bucket := range policyId.Buckets {
		revisionIdx, ok := bucket.Aggregations[FieldRevisionIdx]
		if !ok || len(revisionIdx.Hits) != 1 {
			return nil, ErrMissingAggregations
		}
		hit := revisionIdx.Hits[0]
		err = hit.Unmarshal(&policies[i])
		if err != nil {
			return nil, err
		}
	}
	return policies, nil
}

// CreatePolicy creates a new policy in the index
func CreatePolicy(ctx context.Context, bulker bulk.Bulk, policy model.Policy, opt ...Option) (string, error) {
	o := newOption(FleetPolicies, opt...)
	data, err := json.Marshal(&policy)
	if err != nil {
		return "", err
	}
	return bulker.Create(ctx, o.indexName, "", data)
}

// FindPolicyByID find policy by ID
func FindPolicyByID(ctx context.Context, bulker bulk.Bulk, policyID string) (policy model.Policy, err error) {
	res, err := SearchWithOneParam(ctx, bulker, queryPolicyByID, FleetPolicies, FieldPolicyId, policyID)
	if err != nil {
		return
	}

	if len(res.Hits) == 0 {
		return policy, ErrNotFound
	}

	err = res.Hits[0].Unmarshal(&policy)
	return policy, err
}
