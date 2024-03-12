package main

import (
	"log"
	"os"

	"github.com/eblechschmidt/mindmap/internal/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "mm",
		Usage:    "Declarative mind mapping using yaml files",
		Commands: []*cli.Command{cmd.Serve()},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
