package main

import (
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

// constructDialOptions constructs a list of grpc dial options
func constructDialOptions(
	maxCallRecvMsgSize int,
	grpcRetries uint,
	grpcRetryDelay time.Duration,
	extraOpts ...grpc.DialOption,
) []grpc.DialOption {
	var transportSecurity grpc.DialOption
	transportSecurity = grpc.WithTransportCredentials(insecure.NewCredentials())

	if maxCallRecvMsgSize == 0 {
		maxCallRecvMsgSize = 10 * 5 << 20 // Default 50Mb
	}

	defaultCallOptions := grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(maxCallRecvMsgSize),
		grpc_retry.WithMax(grpcRetries),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(grpcRetryDelay)),
	)

	dialOpts := []grpc.DialOption{
		transportSecurity,
		defaultCallOptions,
	}

	dialOpts = append(dialOpts, extraOpts...)

	return dialOpts
}
