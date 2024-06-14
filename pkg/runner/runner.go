package runner

import (
	"io"
	"log/slog"
	"os"
	"os/exec"
	"strings"
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
