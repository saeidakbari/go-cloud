// Copyright 2019 The Go Cloud Development Kit Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package trace provides support for OpenCensus tracing.
package trace

import (
	"context"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"gocloud.dev/gcerrors"
)

// StartSpan adds a span to the trace with the given name.
func StartSpan(ctx context.Context, name string) context.Context {
	span := trace.SpanFromContext(ctx)
	span.SetName(name)
	return trace.ContextWithSpan(ctx, span)
}

// EndSpan ends a span with the given error.
func EndSpan(ctx context.Context, err error) {
	span := trace.SpanFromContext(ctx)
	if err != nil {
		span.SetStatus(toStatus(err), err.Error())
	}
	span.End()
}

// toStatus interrogates an error and converts it to an appropriate
// OpenCensus status.
func toStatus(err error) codes.Code {
	return codes.Code(gcerrors.Code(err))
}
