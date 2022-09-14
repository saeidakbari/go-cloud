// Copyright 2018 The Go Cloud Development Kit Authors
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

// Package driver defines an interface for custom Net listeners.
// Application code should use package grpcserver.
package driver // import "gocloud.dev/server/driver"

import (
	"context"
	"net"
)

// Server connect grpc server instance to net listener.
type Server interface {
	// ListenAndServe listens on the TCP network address addr and then
	// calls Serve with listener to listen to incoming connections.
	// The addr argument will be a non-empty string specifying "host:port".
	// The net.Listener will always be non-nil.
	// Drivers must block until serving is done (or
	// return an error if serving can't occur for some reason)
	// and be interruptible by Shutdown.
	// Drivers should use the given address if they serve using TCP directly.
	ListenAndServe(addr string, l net.Listener) error

	// Shutdown gracefully shuts down the server without interrupting
	// any active connections. If the provided context expires before
	// the shutdown is complete, Shutdown returns the context's error,
	// otherwise it returns any error returned from closing the Server's
	// underlying Listener(s).
	Shutdown(ctx context.Context) error
}

// TLSServer is an optional interface for Server drivers, that adds support
// for serving TLS.
type TLSServer interface {
	// ListenAndServeTLS is similar to Server.ListenAndServe, but should
	// serve using TLS.
	// See http://go/godoc/net/http/#Server.ListenAndServeTLS.
	ListenAndServeTLS(addr, certFile, keyFile string, l net.Listener) error
}
