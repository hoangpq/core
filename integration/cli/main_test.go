// +build real_integration

package cli_test

import (
	"log"
	"os"
	"testing"

	"github.com/mesg-foundation/core/integration"
)

var options = integration.RunOptions{
	CoreName:      "core-i-c",
	DBPath:        "/tmp/.mesg-i-c",
	ServerAddress: ":30052",
	ClientAddress: "localhost:30052",
}

func TestMain(m *testing.M) {
	// start an integration testing environment.
	if err := integration.Run(options); err != nil {
		log.Fatalln("err cli integration run:", err)
	}

	// run tests.
	status := m.Run()

	integration.Cleanup(options)
	os.Exit(status)
}
