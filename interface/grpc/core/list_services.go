package core

import (
	"context"

	"github.com/mesg-foundation/core/grpcclient"
)

// ListServices lists services.
func (s *Server) ListServices(ctx context.Context, request *grpcclient.ListServicesRequest) (*grpcclient.ListServicesReply, error) {
	services, err := s.api.ListServices()
	if err != nil {
		return nil, err
	}
	return &grpcclient.ListServicesReply{Services: toProtoServices(services)}, nil
}
