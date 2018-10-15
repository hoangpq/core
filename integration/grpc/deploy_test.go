// +build real_integration

package grpc_test

import (
	"context"
	"testing"

	"github.com/mesg-foundation/core/protobuf/coreapi"
	"github.com/stretchr/testify/require"
)

func TestDeployFromURL(t *testing.T) {
	client := newClient(t)
	c, err := client.DeployService(context.Background())
	require.NoError(t, err)

	require.NoError(t, c.Send(&coreapi.DeployServiceRequest{
		Value: &coreapi.DeployServiceRequest_Url{
			Url: "https://github.com/mesg-foundation/service-webhook",
		},
	}))
	require.NoError(t, c.CloseSend())

	message, err := c.Recv()
	require.NoError(t, err)
	require.Empty(t, message.GetServiceID())
	require.Empty(t, message.GetValidationError())
	require.Equal(t, coreapi.DeployServiceReply_Status_RUNNING, message.GetStatus().Type)
	require.Equal(t, "Downloading service...", message.GetStatus().Message)

	message, err = c.Recv()
	require.NoError(t, err)
	require.Empty(t, message.GetServiceID())
	require.Empty(t, message.GetValidationError())
	require.Equal(t, coreapi.DeployServiceReply_Status_DONE_POSITIVE, message.GetStatus().Type)
	require.Equal(t, "Service downloaded with success", message.GetStatus().Message)

	message, err = c.Recv()
	require.NoError(t, err)
	require.Empty(t, message.GetServiceID())
	require.Empty(t, message.GetValidationError())
	require.Equal(t, coreapi.DeployServiceReply_Status_RUNNING, message.GetStatus().Type)
	require.Equal(t, "Receiving service context...", message.GetStatus().Message)

	message, err = c.Recv()
	require.NoError(t, err)
	require.Empty(t, message.GetServiceID())
	require.Empty(t, message.GetValidationError())
	require.Equal(t, coreapi.DeployServiceReply_Status_DONE_POSITIVE, message.GetStatus().Type)
	require.Equal(t, "Service context received with success", message.GetStatus().Message)

	message, err = c.Recv()
	require.NoError(t, err)
	require.Empty(t, message.GetServiceID())
	require.Empty(t, message.GetValidationError())
	require.Equal(t, coreapi.DeployServiceReply_Status_RUNNING, message.GetStatus().Type)
	require.Equal(t, "Building Docker image...", message.GetStatus().Message)

	message, err = c.Recv()
	require.NoError(t, err)
	require.Empty(t, message.GetServiceID())
	require.Empty(t, message.GetValidationError())
	require.Equal(t, coreapi.DeployServiceReply_Status_DONE_POSITIVE, message.GetStatus().Type)
	require.Equal(t, "Image built with success", message.GetStatus().Message)

	message, err = c.Recv()
	require.NoError(t, err)
	require.NotEmpty(t, message.GetServiceID())
	require.Empty(t, message.GetValidationError())
	require.Zero(t, message.GetStatus())
}
