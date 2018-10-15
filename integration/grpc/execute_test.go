// +build real_integration

package grpc_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/mesg-foundation/core/protobuf/coreapi"
	"github.com/stretchr/testify/require"
)

// TestExecute deploys and starts service-webhook from Github repo
// and executes a task on it and makes assertions about task results.
func TestExecute(t *testing.T) {
	client := newClient(t)
	c, err := client.DeployService(context.Background())
	require.NoError(t, err)

	require.NoError(t, c.Send(&coreapi.DeployServiceRequest{
		Value: &coreapi.DeployServiceRequest_Url{
			Url: "https://github.com/mesg-foundation/service-webhook",
		},
	}))
	require.NoError(t, c.CloseSend())

	var serviceID string
	for {
		reply, err := c.Recv()
		require.NoError(t, err)
		serviceID = reply.GetServiceID()
		if serviceID != "" {
			break
		}
	}

	_, err = client.StartService(context.Background(), &coreapi.StartServiceRequest{
		ServiceID: serviceID,
	})
	require.NoError(t, err)

	resultStream, err := client.ListenResult(context.Background(), &coreapi.ListenResultRequest{
		ServiceID: serviceID,
	})
	require.NoError(t, err)

	time.Sleep(time.Second * 3)

	executeReply, err := client.ExecuteTask(context.Background(), &coreapi.ExecuteTaskRequest{
		ServiceID: serviceID,
		TaskKey:   "call",
		InputData: `{"url": "http://ip-api.com/json", "data": {}, "headers": {}}`,
	})
	require.NoError(t, err)

	reply, err := resultStream.Recv()
	require.NoError(t, err)
	require.Equal(t, executeReply.ExecutionID, reply.ExecutionID)

	var data map[string]interface{}
	require.NoError(t, json.Unmarshal([]byte(reply.OutputData), &data))
	require.Equal(t, "success", data["status"].(string))
}
