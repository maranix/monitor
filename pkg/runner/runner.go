package runner

import (
	"context"
	"errors"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/maranix/monitor/pkg/fsutil"
)

func Validate(run string) error {
	executable, _ := splitCmd(run)

	_, err := exec.LookPath(executable)
	if err != nil {
		_, err = fsutil.IsPathValidAndAccessible(executable)
		if err != nil {
			return err
		}

		return errors.New("**Invalid Runner Target:**\nCould not find the runner target either in environment or provided path.")
	}

	return nil
}

func RunContext(ctx context.Context, cmd string) (context.CancelFunc, error) {
	executable, args := splitCmd(cmd)
	context, cancel := context.WithCancel(ctx)

	process := exec.CommandContext(context, executable, args...)
	err := startAndWait(process)
	if err != nil {
		cancel()
		return nil, err
	}

	return cancel, nil
}

func splitCmd(cmd string) (exec string, args []string) {
	c := strings.Split(cmd, " ")
	executable := c[0]
	arguments := c[1:]

	return executable, arguments
}

func startAndWait(process *exec.Cmd) error {
	stdout, err := process.StdoutPipe()
	if err != nil {
		return err
	}

	if err := process.Start(); err != nil {
		return err
	}

	io.Copy(os.Stdout, stdout)

	if err := process.Wait(); err != nil {
		return err
	}

	return nil
}
