package cmd

import (
	"errors"
	"fmt"
	"os"

	e "github.com/eblechschmidt/mindmap/internal/error"
	"github.com/urfave/cli/v2"
)

func Serve() *cli.Command {
	return &cli.Command{
		Name:        "serve",
		Usage:       "serve a file and open a browser",
		Description: "Convert a file into a mind map and serve it via a web browser",
		Action:      runServe,
	}
}
func runServe(ctx *cli.Context) error {
	if ctx.Args().Len() != 1 {
		return e.NewUserErr("Please specify exactly one file to be served", nil)
	}
	fname := ctx.Args().Get(0)
	if err := fileExists(fname); err != nil {
		return userErr(err)
	}
	return nil
}

func userErr(err error) error {
	var e e.UserErr
	if errors.As(err, &e) {
		return fmt.Errorf(e.UserMsg())
	}
	return err
}

// function to check if file exists
func fileExists(fileName string) error {
	_, err := os.Stat(fileName)

	// check if error is "file not exists"
	if os.IsNotExist(err) {
		return e.NewUserErr(fmt.Sprintf("The file '%s' does not exist", fileName), err)
	} else {
		return nil
	}
}
