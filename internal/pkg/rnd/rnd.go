// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package rnd

import (
	"crypto/rand"
	"math/big"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func Int(min, max int) (int, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return 0, err
	}
	return int(r.Int64() + int64(min)), nil
}

func Bool() (bool, error) {
	n, err := Int(0, 2)
	if err != nil {
		return false, err
	}
	return n != 0, nil
}

func String(sz int) (string, error) {
	b := make([]byte, sz)
	for i := range b {
		n, err := Int(0, len(charset))
		if err != nil {
			return "", err
		}
		b[i] = charset[n]
	}
	return string(b), nil
}

type OffsetDirection int

const (
	TimeBefore = iota
	TimeAfter
)

func (d OffsetDirection) String() string {
	return []string{"Before", "After"}[d]
}

func Time(tm time.Time, min, max int, units time.Duration, direction OffsetDirection) (time.Time, error) {
	n, err := Int(min, max)
	if err != nil {
		return time.Time{}, err
	}

	dur := time.Duration(n) * units

	if direction == TimeBefore {
		return tm.Add(-dur), nil
	}
	return tm.Add(dur), nil
}
