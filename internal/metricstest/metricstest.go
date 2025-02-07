// Copyright (c) 2022 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metricstest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ExpectedMetric contains metrics under test.
type ExpectedMetric struct {
	Name  string
	Tags  map[string]string
	Value int
}

// TODO do something similar for Timers

// AssertCounterMetrics checks if counter metrics exist.
func (f *Factory) AssertCounterMetrics(t *testing.T, expectedMetrics ...ExpectedMetric) {
	counters, _ := f.Snapshot()
	assertMetrics(t, counters, expectedMetrics...)
}

// AssertGaugeMetrics checks if gauge metrics exist.
func (f *Factory) AssertGaugeMetrics(t *testing.T, expectedMetrics ...ExpectedMetric) {
	_, gauges := f.Snapshot()
	assertMetrics(t, gauges, expectedMetrics...)
}

func assertMetrics(t *testing.T, actualMetrics map[string]int64, expectedMetrics ...ExpectedMetric) {
	for _, expected := range expectedMetrics {
		key := GetKey(expected.Name, expected.Tags, "|", "=")
		assert.EqualValues(t,
			expected.Value,
			actualMetrics[key],
			"expected metric name: %s, tags: %+v", expected.Name, expected.Tags,
		)
	}
}
