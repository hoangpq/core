package service

import (
	"os"
	"strings"
	"testing"

	"github.com/stvp/assert"
)

func TestStopRunningService(t *testing.T) {
	if os.Getenv("CI") == "true" {
		return
	}
	service := &Service{
		Name: "TestStopRunningService",
		Dependencies: map[string]Dependency{
			"test": Dependency{
				Image: "nginx",
			},
		},
	}
	service.Start()
	err := service.Stop()
	assert.Nil(t, err)
	assert.Equal(t, service.IsStopped(), true)
}

func TestStopNonRunningService(t *testing.T) {
	if os.Getenv("CI") == "true" {
		return
	}
	service := &Service{
		Name: "TestStopNonRunningService",
		Dependencies: map[string]Dependency{
			"test": Dependency{
				Image: "nginx",
			},
		},
	}
	err := service.Stop()
	assert.Nil(t, err)
	assert.Equal(t, service.IsStopped(), true)
}

func TestStopDependency(t *testing.T) {
	if os.Getenv("CI") == "true" {
		return
	}
	namespace := strings.Join([]string{NAMESPACE, "TestStopDependency"}, "_")
	name := "test"
	dependency := Dependency{Image: "nginx"}
	dependency.Start(namespace, name)
	err := dependency.Stop(namespace, name)
	assert.Nil(t, err)
	assert.Equal(t, dependency.IsStopped(namespace, name), true)
	assert.Equal(t, dependency.IsRunning(namespace, name), false)
}