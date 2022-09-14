package grpcserver

import (
	"net"
	"sync"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	once sync.Once

	serverOpts         []grpc.ServerOption
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor

	wrappedGRPCServer *grpc.Server
}

// Options is the set of optional parameters.
type Options struct {
	// ServerOpts is a list of options to be passed to grpc.Server
	ServerOpts []grpc.ServerOption

	// UnaryInterceptors are interceptors for unary requests
	UnaryInterceptors []grpc.UnaryServerInterceptor

	// StreamInterceptors are interceptors for stream connections
	StreamInterceptors []grpc.StreamServerInterceptor
}

func New(opts Options) *GRPCServer {
	return &GRPCServer{
		serverOpts:         opts.ServerOpts,
		unaryInterceptors:  opts.UnaryInterceptors,
		streamInterceptors: opts.StreamInterceptors,
	}
}

func (grv *GRPCServer) init() {
	grv.once.Do(func() {
		var opt []grpc.ServerOption

		if grv.serverOpts != nil {
			opt = append(opt, grv.serverOpts...)
		}

		if len(grv.unaryInterceptors) > 0 {
			opt = append(opt, grpc.ChainUnaryInterceptor(grv.unaryInterceptors...))
		}

		if len(grv.streamInterceptors) > 0 {
			opt = append(opt, grpc.ChainStreamInterceptor())
		}

		grv.wrappedGRPCServer = grpc.NewServer(opt...)
	})
}

func (grv *GRPCServer) ListenAndServe(addr string, l net.Listener) error {
	grv.wrappedGRPCServer.ServeHTTP()
}

// grpc.NewServer(
// 	grpc.ChainUnaryInterceptor(opts.unaryInterceptors...),
// 	grpc.ChainStreamInterceptor(opts.streamInterceptors...),
// 	opts.ServerOpts)
