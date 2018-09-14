package core

import (
	"context"

	"github.com/mesg-foundation/core/grpcclient"
)

// StartService starts a service.
func (s *Server) StartService(ctx context.Context, request *grpcclient.StartServiceRequest) (*grpcclient.StartServiceReply, error) {
	return &grpcclient.StartServiceReply{}, s.api.StartService(request.ServiceID)
}
