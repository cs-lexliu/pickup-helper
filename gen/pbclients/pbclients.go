package pbclients

import (
	"context"
	"fmt"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"

	ointerceptors "github.com/carousell/Orion/interceptors"
	"github.com/carousell/messaging/messaging/client/asynccall"
)

type Config struct {
	Address     string
	Retries     uint
	DialTimeout int // in millisecond
}

func GetConfig(address string, retries uint, dialTimeout int) Config {
	if dialTimeout == 0 {
		dialTimeout = 500
	}
	return Config{
		Address:     address,
		Retries:     retries,
		DialTimeout: dialTimeout,
	}
}

func GetGrpcClientConn(cfg Config, interceptors ...grpc.UnaryClientInterceptor) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.DialTimeout)*time.Millisecond)
	defer cancel()
	// With blocking dial, the ContextWithTimeout ensures the establish of connection won't take long
	return grpc.DialContext(
		ctx,
		cfg.Address,
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(
			getInterceptors(cfg, interceptors...)...,
		),
	)
}

// GetAsyncGrpcConn establishes grpc connection for asynchronous calling with default interceptors
func GetAsyncGrpcConn(serverCfg Config, messagingCfg Config, callerID string) (*grpc.ClientConn, *grpc.ClientConn, error) {
	msgConn, err := getMessagingGrpcClientConn(context.Background(), messagingCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("getMessagingGrpcClientConn: %w", err)
	}

	clientConn, err := GetGrpcClientConn(serverCfg, getAsyncCallUnaryClientInterceptor(msgConn, callerID))
	return msgConn, clientConn, err
}

func getMessagingGrpcClientConn(ctx context.Context, cfg Config) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx,
		cfg.Address,
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(
			getInterceptors(cfg)...,
		),
	)
}

func getInterceptors(cfg Config, interceptors ...grpc.UnaryClientInterceptor) []grpc.UnaryClientInterceptor {
	res := []grpc.UnaryClientInterceptor{
		ointerceptors.DefaultClientInterceptor(cfg.Address),
		grpc_prometheus.UnaryClientInterceptor,
	}
	res = append(res, interceptors...)
	return res
}

func getAsyncCallUnaryClientInterceptor(msgConn *grpc.ClientConn, callerId string) grpc.UnaryClientInterceptor {
	return asynccall.UnaryClientInterceptor(
		asynccall.NewAsynCallCallerFromClientConn(msgConn, callerId, asynccall.NewConfig()),
		asynccall.WithSyncCall(),
	)
}
