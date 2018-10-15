package integration

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

var (
	// CliExecutable is the integration cli executable of core.
	CliExecutable = os.Getenv("MESG_CLI_EXECUTABLE")
)

var (
	removalInProgress  = regexp.MustCompile(`Error response from daemon: removal of container (\w+) is already in progress`)
	containerNotExists = regexp.MustCompile(`Error: No such container: (\w+)`)
)

// RunOptions of integration env.
type RunOptions struct {
	CoreName      string
	DBPath        string
	ServerAddress string
	ClientAddress string
}

// Run starts integration core daemon.
func Run(o RunOptions) error {
	if err := Cleanup(o); err != nil {
		return err
	}

	cmd := exec.Command(CliExecutable, "start")
	cmd.Env = append(os.Environ(), []string{
		fmt.Sprintf("MESG_CORE_NAME=%s", o.CoreName),
		fmt.Sprintf("MESG_CORE_PATH=%s", o.DBPath),
		fmt.Sprintf("MESG_SERVER_ADDRESS=%s", o.ServerAddress),
	}...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("start integration core err: %s", stderr.String())
	}

	// wait until gRPC server is ready to accept connections.
	for {
		_, err := http.Get(fmt.Sprintf("http://%s", o.ClientAddress))
		if strings.Contains(err.Error(), "malformed HTTP response") {
			break
		}
		time.Sleep(time.Second)
	}

	return nil
}

// Cleanup stops integration core daemon and cleanups artifacts.
func Cleanup(o RunOptions) error {
	// cmd := exec.Command(CliExecutable, "stop")
	// cmd.Env = os.Environ()

	// var stderr bytes.Buffer
	// cmd.Stderr = &stderr

	// if err := cmd.Run(); err != nil {
	// 	return fmt.Errorf("stop integration core err: %s", stderr.String())
	// }

	return manualCleanup(o)
}

// manualCleanup removes all artifacts of integration core server
// to make sure we have a clean environment each time we run the
// integration tests.
func manualCleanup(o RunOptions) error {
	if strings.HasPrefix(o.DBPath, "/tmp/") {
		if err := os.RemoveAll(o.DBPath); err != nil {
			return err
		}
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	filters := filters.NewArgs()
	filters.Add("label", fmt.Sprintf("com.mesg.core.name=%s", o.CoreName))

	// rm integration services
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{
		Filters: filters,
	})
	if err != nil {
		return err
	}
	for _, service := range services {
		if err := cli.ServiceRemove(context.Background(), service.ID); err != nil {
			return err
		}
	}

	// rm integration containers.
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All:     true,
		Filters: filters,
	})
	if err != nil {
		return err
	}
	for _, container := range containers {
		options := types.ContainerRemoveOptions{Force: true}
		for {
			if err := cli.ContainerRemove(context.Background(), container.ID, options); err != nil {
				if removalInProgress.MatchString(err.Error()) {
					time.Sleep(time.Second)
					continue
				}
				if containerNotExists.MatchString(err.Error()) {
					break
				}
				return err
			}
		}
	}

	return nil
}
