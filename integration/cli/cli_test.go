// +build real_integration

package cli_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/mesg-foundation/core/integration"
	"github.com/stretchr/testify/require"
)

// CmdResult keeps the command result of cli.
type CmdResult struct {
	t      *testing.T
	cmd    *exec.Cmd
	stdout io.ReadCloser
	stderr io.ReadCloser
}

// MESG executes the specified MESG command
func MESG(t *testing.T, args ...string) *CmdResult {
	r := &CmdResult{t: t}
	r.cmd = exec.Command(integration.CliExecutable, args...)
	r.cmd.Env = append(os.Environ(), []string{
		fmt.Sprintf("MESG_CLIENT_ADDRESS=%s", options.ClientAddress),
	}...)

	var err error
	r.stdout, err = r.cmd.StdoutPipe()
	require.NoError(t, err)

	r.stderr, err = r.cmd.StderrPipe()
	require.NoError(t, err)
	require.NoError(t, r.cmd.Start())
	return r
}

// StdoutReader returns stdout reader.
func (r *CmdResult) StdoutReader() io.ReadCloser {
	return r.stdout
}

// StderrReader returns stderr reader.
func (r *CmdResult) StderrReader() io.ReadCloser {
	return r.stderr
}

// StdoutAll returns all stdout as string.
func (r *CmdResult) StdoutAll() string {
	b, err := ioutil.ReadAll(r.stdout)
	require.NoError(r.t, err)
	require.NoError(r.t, r.cmd.Wait())
	return string(b)
}

// StderrAll returns all stderr as string.
func (r *CmdResult) StderrAll() string {
	b, err := ioutil.ReadAll(r.stderr)
	require.NoError(r.t, err)
	require.Error(r.t, r.cmd.Wait())
	return string(b)
}
