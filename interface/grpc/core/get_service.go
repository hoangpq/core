package core

import (
	"context"

	"github.com/mesg-foundation/core/grpcclient"
)

// GetService returns service serviceID.
func (s *Server) GetService(ctx context.Context, request *grpcclient.GetServiceRequest) (*grpcclient.GetServiceReply, error) {
	srv, err := s.api.GetService(request.ServiceID)
	if err != nil {
		return nil, err
	}
	return &grpcclient.GetServiceReply{Service: toProtoService(srv)}, nil
}
