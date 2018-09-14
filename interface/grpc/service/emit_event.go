package service

import (
	"context"
	"encoding/json"

	"github.com/mesg-foundation/core/grpcclient"
)

// EmitEvent permits to send and event to anyone who subscribed to it.
func (s *Server) EmitEvent(context context.Context, request *grpcclient.EmitEventRequest) (*grpcclient.EmitEventReply, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(request.EventData), &data); err != nil {
		return nil, err
	}
	return &grpcclient.EmitEventReply{}, s.api.EmitEvent(request.Token, request.EventKey, data)
}
