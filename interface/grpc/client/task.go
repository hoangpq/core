package client

import (
	"context"
	"encoding/json"

	"github.com/mesg-foundation/core/protobuf/coreapi"
)

func (task *Task) processEvent(wf *Workflow, data *coreapi.EventData) (err error) {
	return task.process(wf, data.EventData)
}

func (task *Task) processResult(wf *Workflow, data *coreapi.ResultData) (err error) {
	return task.process(wf, data.OutputData)
}

func (task *Task) process(wf *Workflow, data string) (err error) {
	var d interface{}
	err = json.Unmarshal([]byte(data), &d)
	if err != nil {
		return
	}
	inputData, err := task.convertData(d)
	if err != nil {
		return
	}
	_, err = wf.client.ExecuteTask(context.Background(), &coreapi.ExecuteTaskRequest{
		ServiceID: task.ServiceID,
		TaskKey:   task.Name,
		InputData: inputData,
	})
	return
}

func (task *Task) convertData(data interface{}) (res string, err error) {
	inputData := task.Inputs(data)
	var inputDataJSON []byte
	inputDataJSON, err = json.Marshal(inputData)
	if err != nil {
		return
	}
	res = string(inputDataJSON)
	return
}
