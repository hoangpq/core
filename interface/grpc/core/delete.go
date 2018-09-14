package core

import (
	"context"

	"github.com/mesg-foundation/core/grpcclient"
)

// DeleteService stops and deletes service serviceID.
func (s *Server) DeleteService(ctx context.Context, request *grpcclient.DeleteServiceRequest) (*grpcclient.DeleteServiceReply, error) {
	return &grpcclient.DeleteServiceReply{}, s.api.DeleteService(request.ServiceID)
}
