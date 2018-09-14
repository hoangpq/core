package core

import (
	"context"

	"github.com/mesg-foundation/core/grpcclient"
)

// StopService stops a service.
func (s *Server) StopService(ctx context.Context, request *grpcclient.StopServiceRequest) (*grpcclient.StopServiceReply, error) {
	return &grpcclient.StopServiceReply{}, s.api.StopService(request.ServiceID)
}
