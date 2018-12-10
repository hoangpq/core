package api

import (
	"encoding/json"

	"github.com/mesg-foundation/core/execution"

	"github.com/mesg-foundation/core/pubsub"
	"github.com/mesg-foundation/core/service"
)

// SubmitResult submits results for executionID.
func (a *API) SubmitResult(executionID string, outputKey string, outputData string) error {
	return newResultSubmitter(a).Submit(executionID, outputKey, outputData)
}

// resultSubmitter provides functionalities to submit a MESG task result.
type resultSubmitter struct {
	api *API
}

// newResultSubmitter creates a new resultSubmitter with given api.
func newResultSubmitter(api *API) *resultSubmitter {
	return &resultSubmitter{
		api: api,
	}
}

// Submit submits results for executionID.
func (s *resultSubmitter) Submit(executionID string, outputKey string, outputData string) error {
	exec, err := s.processExecution(executionID, outputKey, outputData)
	if err != nil {
		return err
	}
	go pubsub.Publish(exec.Service.ResultSubscriptionChannel(), exec)
	return nil
}

func (s *resultSubmitter) processExecution(executionID string, outputKey string, outputData string) (*execution.Execution, error) {
	exec, err := s.api.execDB.Find(executionID)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(outputData), &data); err != nil {
		exec.Fail(err)
		return exec, nil
	}
	exec.Service, err = service.FromService(exec.Service, service.ContainerOption(s.api.container))
	if err != nil {
		exec.Fail(err)
		return exec, nil
	}
	if err := exec.Complete(outputKey, data); err != nil {
		exec.Fail(err)
		return exec, nil
	}
	if err = s.api.execDB.Save(exec); err != nil {
		exec.Fail(err)
		return exec, nil
	}
	return exec, nil
}
