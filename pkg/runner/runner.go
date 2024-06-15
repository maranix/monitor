package runner

import (
	"errors"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"strings"

	"github.com/maranix/monitor/pkg/fsutil"
)

func Run(cmd string) {
	executable, args := splitCmd(cmd)

	if ok := executableExists(executable); !ok {
		slog.Error("Executable not found for the given command")
		os.Exit(1)
	}

	p := exec.Command(executable, args...)
	stdout, err := p.StdoutPipe()
	if err != nil {
		slog.Error(err.Error())
	}

	if err := p.Start(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	io.Copy(os.Stdout, stdout)

	if err := p.Wait(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func splitCmd(cmd string) (exec string, args []string) {
	c := strings.Split(cmd, " ")
	executable := c[0]
	arguments := c[1:]

	return executable, arguments
}

func executableExists(e string) bool {
	_, err := exec.LookPath(e)

	return err == nil
}

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
