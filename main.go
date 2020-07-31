package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"

	goflags "github.com/jessevdk/go-flags"

	"foobar/marsh"
)

type ApplicationOptions struct {
	Verbose    []bool `short:"v" long:"verbose" description:"verbose output"`
	StartIndex uint32 `short:"si" long:"start-index" description:"starting index" default:0`
}

func run(w io.Writer, args []string) (return_value int) {
	return_value = 0

	applicationConfig := ApplicationOptions{}

	parser := goflags.NewParser(&applicationConfig, goflags.Default)

	if _, err := parser.ParseArgs(args); err != nil {
		return_value = 1
		return
	}

	fields := marsh.New(applicationConfig.StartIndex)
	fields.AddAttribute("Marker1")
	fields.AddAttribute("Marker2")
	fields.AddAttribute("Marker3")
	fields.AddAttribute("Silly Sarah Squats Salad Squash Sasquatches")
	fields.AddAttribute("Timmy's Toys Take Talent Totally Tangential To Tiger Tails")

	fmt.Sprintf("Index: %d - Random: %d", fields.Index, fields.Random)

	// verify the Comparison function fails against an uninitiated object
	if true == fields.Compare(marsh.New(rand.Uint32())) {
		fmt.Sprintf("Bad Comparison")
		return_value = 2
		return
	}

	marshalledData := fields.Marshal()

	unmarshalled := marsh.Unmarshal(marshalledData)

	if false == fields.Compare(unmarshalled) {
		fmt.Sprintf("Mis-match")
		return_value = 3
		return
	}

	streamMarshalled, umErr := fields.StreamMarshal()
	if umErr != nil {
		streamUnmarshalled, unmErr := fields.StreamUnmarshal(streamMarshalled)

		if unmErr != nil {
			if false == fields.Compare(streamUnmarshalled) {
				fmt.Sprintf("Mis-match")
				return_value = 4
				return
			}
		} else {
			fmt.Sprintf("Failed to stream unmarshall the data: %s", unmErr)
			return_value = 5
		}
	} else {
		fmt.Sprintf("Failed to stream marshall the data: %s", umErr)
		return_value = 6
	}

	// reached the end
	return
}

func main() {
	os.Exit(run(os.Stdout, os.Args))
}
