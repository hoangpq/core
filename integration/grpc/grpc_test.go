// +build real_integration

package grpc_test

import (
	"context"
	"testing"

	"github.com/mesg-foundation/core/protobuf/coreapi"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func newClient(t *testing.T) coreapi.CoreClient {
	conn, err := grpc.Dial(options.ClientAddress, grpc.WithInsecure())
	if err != nil {
		require.NoError(t, err)
	}
	return coreapi.NewCoreClient(conn)
}

func ctxTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), callTimeout)
}
