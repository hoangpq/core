// +build real_integration

package cli_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeployFromURL(t *testing.T) {
	r := MESG(t, "service", "deploy", "https://github.com/mesg-foundation/service-webhook")
	out := strings.Split(r.StdoutAll(), "\n")
	require.Regexp(t, `Service downloaded with success`, out[0])
	require.Regexp(t, `Service context received with success`, out[1])
	require.Regexp(t, `Image built with success`, out[2])
	require.Regexp(t, `Service deployed with ID: [a-z0-9]{40}`, out[3])
	require.Regexp(t, `To start it, run the command:`, out[4])
	require.Regexp(t, `mesg-core service start [a-z0-9]{40}`, out[5])
}
