package service

import (
	"context"
	"encoding/json"

	"github.com/mesg-foundation/core/grpcclient"
)

// SubmitResult submits results of an execution.
func (s *Server) SubmitResult(context context.Context, request *grpcclient.SubmitResultRequest) (*grpcclient.SubmitResultReply, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(request.OutputData), &data); err != nil {
		return nil, err
	}
	return &grpcclient.SubmitResultReply{}, s.api.SubmitResult(request.ExecutionID, request.OutputKey, data)
}
