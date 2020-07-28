package main

import (
	"fmt"
	"io"
	"os"

	 goflags "github.com/jessevdk/go-flags"

	 "foobar/marsh"
)

type ApplicationOptions struct {
	Verbose []bool `short:"v" long:"verbose" description:"verbose output"`
	StartIndex uint32 `short:"si" long:"start-index" description:"starting index" default:0`
}


func run(w io.Writer, args []string) (return_value int) {
	return_value = 0

	applicationConfig := ApplicationOptions{}

	parser := goflags.NewParser(&applicationConfig, goflags.Default)

	if _, err := parser.ParseArgs(args); err != nil {
		return 1
	}

	fields := marsh.New(applicationConfig.StartIndex)
	fmt.Sprintf("Index: %d - Random: %d", fields.Index, fields.Random)

	return
}


func main() {
	os.Exit(run(os.Stdout, os.Args))
}
