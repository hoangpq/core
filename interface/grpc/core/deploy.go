package core

import (
	"io"
	"sync"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/mesg-foundation/core/api"
	"github.com/mesg-foundation/core/protobuf/coreapi"
	service "github.com/mesg-foundation/core/service"
	"github.com/mesg-foundation/core/service/importer"
)

// DeployService deploys a service from Git URL or service.tar.gz file. It'll send status
// events during the process and finish with sending service id or validation error.
// TODO(ilgooz): sync `stream.Send()`s by doing it in a single goroutine.
func (s *Server) DeployService(stream coreapi.Core_DeployServiceServer) error {
	var (
		sr       = newDeployServiceStreamReader(stream)
		statuses = make(chan api.DeployStatus)
		wg       sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		forwardDeployStatuses(statuses, stream)
	}()

	var (
		service         *service.Service
		validationError *importer.ValidationError
		err             error
	)

	// receive the first message in the stream.
	if err := sr.RecvMessage(); err != nil {
		return err
	}

	deployOptions := []api.DeployServiceOption{
		api.DeployServiceStatusOption(statuses),
	}

	// receive confirmation for a force deploy.
	// if it's not provided deployment will start anyway.
	if sr.Confirmation != nil {
		confirmation := sr.Confirmation.GetValue()
		deployOptions = append(deployOptions, api.DeployServiceConfirmationOption(func(sid string) bool {
			return confirmation
		}))
	} else {
		deployOptions = append(deployOptions, api.DeployServiceConfirmationOption(func(sid string) bool {
			// request for confirmation.
			if err := stream.Send(&coreapi.DeployServiceReply{
				Value: &coreapi.DeployServiceReply_RequestConfirmation{RequestConfirmation: sid},
			}); err != nil {
				return false
			}

			// receive the confirmation result.
			// TODO(ilgooz) add timeout.
			if err := sr.RecvMessage(); err != nil {
				return false
			}
			if sr.Confirmation == nil {
				return false
			}
			return sr.Confirmation.GetValue()
		}))
	}

	if sr.Confirmation != nil {
		if err := sr.RecvMessage(); err != nil {
			return err
		}
	}

	if sr.URL != "" {
		service, validationError, err = s.api.DeployServiceFromURL(sr.URL, deployOptions...)
	} else {
		service, validationError, err = s.api.DeployService(sr, deployOptions...)
	}
	wg.Wait()

	if err != nil {
		return err
	}
	if validationError != nil {
		return stream.Send(&coreapi.DeployServiceReply{
			Value: &coreapi.DeployServiceReply_ValidationError{ValidationError: validationError.Error()},
		})
	}

	return stream.Send(&coreapi.DeployServiceReply{
		Value: &coreapi.DeployServiceReply_ServiceID{ServiceID: service.Hash},
	})
}

func forwardDeployStatuses(statuses chan api.DeployStatus, stream coreapi.Core_DeployServiceServer) {
	for status := range statuses {
		var typ coreapi.DeployServiceReply_Status_Type
		switch status.Type {
		case api.Running:
			typ = coreapi.DeployServiceReply_Status_RUNNING
		case api.DonePositive:
			typ = coreapi.DeployServiceReply_Status_DONE_POSITIVE
		case api.DoneNegative:
			typ = coreapi.DeployServiceReply_Status_DONE_NEGATIVE
		}
		stream.Send(&coreapi.DeployServiceReply{
			Value: &coreapi.DeployServiceReply_Status_{
				Status: &coreapi.DeployServiceReply_Status{
					Message: status.Message,
					Type:    typ,
				},
			},
		})
	}
}

type deployServiceStreamReader struct {
	stream coreapi.Core_DeployServiceServer

	URL          string
	Confirmation *wrappers.BoolValue

	chunk     []byte
	chunkDone bool
	i         int64
}

func newDeployServiceStreamReader(stream coreapi.Core_DeployServiceServer) *deployServiceStreamReader {
	return &deployServiceStreamReader{
		stream: stream,
	}
}

// RecvMessage receives the next message in gRPC stream.
func (r *deployServiceStreamReader) RecvMessage() error {
	message, err := r.stream.Recv()
	if err != nil {
		return err
	}
	r.Confirmation = message.GetConfirmation()
	r.URL = message.GetUrl()
	r.chunk = message.GetChunk()
	r.chunkDone = message.GetChunkDone()
	return nil
}

// Read reads service chunks to deploy.
func (r *deployServiceStreamReader) Read(p []byte) (n int, err error) {
	if r.i >= int64(len(r.chunk)) {
		if err := r.RecvMessage(); err != nil {
			return 0, err
		}
		if r.chunkDone {
			return 0, io.EOF
		}
		r.i = 0
		return r.Read(p)
	}
	n = copy(p, r.chunk[r.i:])
	r.i += int64(n)
	return n, nil
}
