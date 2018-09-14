package core

import (
	"github.com/mesg-foundation/core/grpcclient"
	service "github.com/mesg-foundation/core/service"
)

func toProtoServices(ss []*service.Service) []*grpcclient.Service {
	services := make([]*grpcclient.Service, 0)
	for _, s := range ss {
		services = append(services, toProtoService(s))
	}
	return services
}

func toProtoService(s *service.Service) *grpcclient.Service {
	return &grpcclient.Service{
		ID:           s.ID,
		Name:         s.Name,
		Description:  s.Description,
		Repository:   s.Repository,
		Tasks:        toProtoTasks(s.Tasks),
		Events:       toProtoEvents(s.Events),
		Dependencies: toProtoDependencies(s.Dependencies),
	}
}

func toProtoTasks(tasks []*service.Task) []*grpcclient.Task {
	ts := make([]*grpcclient.Task, 0)
	for _, task := range tasks {
		t := &grpcclient.Task{
			Key:         task.Key,
			Name:        task.Name,
			Description: task.Description,
			Inputs:      toProtoParameters(task.Inputs),
			Outputs:     []*grpcclient.Output{},
		}
		for _, output := range task.Outputs {
			o := &grpcclient.Output{
				Key:         output.Key,
				Name:        output.Name,
				Description: output.Description,
				Data:        toProtoParameters(output.Data),
			}
			t.Outputs = append(t.Outputs, o)
		}
		ts = append(ts, t)
	}
	return ts
}

func toProtoEvents(events []*service.Event) []*grpcclient.Event {
	es := make([]*grpcclient.Event, 0)
	for _, event := range events {
		e := &grpcclient.Event{
			Key:         event.Key,
			Name:        event.Name,
			Description: event.Description,
			Data:        toProtoParameters(event.Data),
		}
		es = append(es, e)
	}
	return es
}

func toProtoParameters(params []*service.Parameter) []*grpcclient.Parameter {
	ps := make([]*grpcclient.Parameter, 0)
	for _, param := range params {
		p := &grpcclient.Parameter{
			Key:         param.Key,
			Name:        param.Name,
			Description: param.Description,
			Type:        param.Type,
			Optional:    param.Optional,
		}
		ps = append(ps, p)
	}
	return ps
}

func toProtoDependency(dep *service.Dependency) *grpcclient.Dependency {
	if dep == nil {
		return nil
	}
	return &grpcclient.Dependency{
		Key:         dep.Key,
		Image:       dep.Image,
		Volumes:     dep.Volumes,
		Volumesfrom: dep.VolumesFrom,
		Ports:       dep.Ports,
		Command:     dep.Command,
	}
}

func toProtoDependencies(deps []*service.Dependency) []*grpcclient.Dependency {
	ds := make([]*grpcclient.Dependency, 0)
	for _, dep := range deps {
		ds = append(ds, toProtoDependency(dep))
	}
	return ds
}
