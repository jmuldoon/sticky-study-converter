package main

import (
	"fmt"
	"os"

	"github.com/jmuldoon/sticky-study-converter/arguments"
	"github.com/jmuldoon/sticky-study-converter/exporter"
	"github.com/jmuldoon/sticky-study-converter/importer"
)

// Exit Codes
const (
	ExitArgParseErr = 0 + iota
	ExitWriteErr
	ExitReadErr
)

func main() {
	args := &arguments.Args{}
	// commandLine := flag.NewFlagSet(os.Args[0], ExitOnError)
	// if err := arguments.Parse(args, commandLine); err != nil {
	// TODO: follow up with the adjustment to fully controlled system as detailed above.
	if err := arguments.Parse(args); err != nil {
		fmt.Println(err)
		os.Exit(ExitArgParseErr)
	}

	data, err := importer.Parse(*args.Input, *args.New)
	if err != nil {
		fmt.Println(err)
		os.Exit(ExitReadErr)
	}
	err = exporter.OutputToFile(*args.Ouput, data)
	if err != nil {
		fmt.Println(err)
		os.Exit(ExitWriteErr)
	}
}
