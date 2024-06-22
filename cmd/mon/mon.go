package main

import (
	"fmt"
	"os"

	"github.com/maranix/monitor/internal/app/mon"
)

func main() {
	if err := mon.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
