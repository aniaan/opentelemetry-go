// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package view // import "go.opentelemetry.io/otel/sdk/metric/view"

// InstrumentKind describes the kind of instrument a Meter can create.
//
// Deprecated: Use InstrumentKind in go.opentelemetry.io/otel/sdk/metric
// instead.
type InstrumentKind uint8

// These are all the instrument kinds supported by the SDK.
const (
	// undefinedInstrument is an uninitialized instrument kind, should not be used.
	//nolint:deadcode,varcheck
	undefinedInstrument InstrumentKind = iota
	// SyncCounter is an instrument kind that records increasing values
	// synchronously in application code.
	//
	// Deprecated: Use InstrumentKindSyncCounter in
	// go.opentelemetry.io/otel/sdk/metric instead.
	SyncCounter
	// SyncUpDownCounter is an instrument kind that records increasing and
	// decreasing values synchronously in application code.
	//
	// Deprecated: Use InstrumentKindSyncUpDownCounter in
	// go.opentelemetry.io/otel/sdk/metric instead.
	SyncUpDownCounter
	// SyncHistogram is an instrument kind that records a distribution of
	// values synchronously in application code.
	//
	// Deprecated: Use InstrumentKindSyncHistogram in
	// go.opentelemetry.io/otel/sdk/metric instead.
	SyncHistogram
	// AsyncCounter is an instrument kind that records increasing values in an
	// asynchronous callback.
	//
	// Deprecated: Use InstrumentKindAsyncCounter in
	// go.opentelemetry.io/otel/sdk/metric instead.
	AsyncCounter
	// AsyncUpDownCounter is an instrument kind that records increasing and
	// decreasing values in an asynchronous callback.
	//
	// Deprecated: Use InstrumentKindAsyncUpDownCounter in
	// go.opentelemetry.io/otel/sdk/metric instead.
	AsyncUpDownCounter
	// AsyncGauge is an instrument kind that records current values in an
	// asynchronous callback.
	//
	// Deprecated: Use InstrumentKindAsyncGauge in
	// go.opentelemetry.io/otel/sdk/metric instead.
	AsyncGauge
)
