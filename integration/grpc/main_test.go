// +build real_integration

package grpc_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/mesg-foundation/core/integration"
)

var callTimeout = time.Second
var options = integration.RunOptions{
	CoreName:      "core-i-g",
	DBPath:        "/tmp/.mesg-i-g",
	ServerAddress: ":40052",
	ClientAddress: "localhost:40052",
}

func TestMain(m *testing.M) {
	// start an integration testing environment.
	if err := integration.Run(options); err != nil {
		log.Fatalln("err grpc integration run:", err)
	}

	// run tests.
	status := m.Run()

	integration.Cleanup(options)
	os.Exit(status)
}
