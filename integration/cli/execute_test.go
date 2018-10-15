// +build real_integration

package cli_test

import (
	"encoding/json"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var serviceIDFromDeployTar = regexp.MustCompile(`Service deployed with ID: ([a-z0-9]{40})`)

// TestExecute deploys and starts service-webhook from filesystem
// and executes a task on it and makes assertions about task results.
func TestExecute(t *testing.T) {
	r := MESG(t, "service", "deploy", "../testdata/service-webhook")
	out := strings.Split(r.StdoutAll(), "\n")
	ss := serviceIDFromDeployTar.FindStringSubmatch(out[2])
	require.Len(t, ss, 2)
	serviceID := ss[1]

	MESG(t, "service", "start", serviceID).StdoutAll()

	r = MESG(t, "service", "execute", serviceID,
		"-t", "call",
		"-d", "url=http://ip-api.com/json",
		"-d", "data={}",
		"-d", "headers={}")
	out = strings.Split(r.StdoutAll(), "\n")
	require.Regexp(t, `Task "call" executed`, out[0])
	require.Regexp(t, `Task call returned output result with data:`, out[1])

	jsonData := strings.Join(out[2:], "\n")
	var data map[string]interface{}
	require.NoError(t, json.Unmarshal([]byte(jsonData), &data))
	require.Equal(t, "success", data["status"].(string))
}
