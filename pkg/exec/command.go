// Package exec wraps os/exec to provide command line error logging
package exec

import (
	"context"
	osexec "os/exec"
	"strings"

	"github.com/go-logr/glogr"
)

var log = glogr.New()

// Command wraps os/exec.Command to provide verbosity controlled error logging.
type Cmd struct {
	*osexec.Cmd
}

// CommandContext is a drop-in replacement for os/exec.CommandContext
func CommandContext(ctx context.Context, name string, arg ...string) *Cmd {
	return &Cmd{osexec.CommandContext(ctx, name, arg...)}
}

// Command is a drop-in replacement for os/exec.Command
func Command(name string, arg ...string) *Cmd {
	return &Cmd{osexec.Command(name, arg...)}
}

// Run wraps os/exec.Command.Run to provide verbosity controlled error logging.
func (c *Cmd) Run() error {
	err := c.Cmd.Run()
	if err != nil {
		// Not logging the error message itself: this is left as an exercise
		// for the caller.
		log.V(5).Info("error while running: %v %v\n\t%v",
			c.Cmd.Path, strings.Join(c.Cmd.Args, " "))
	}
	return err
}
