package arguments

import (
	"flag"
	"fmt"
)

const testVersion = 0

var errRequiredArgumentsNotSet = fmt.Errorf("All required arguments were not set")

// Parser is the interface by which the flags are controlled.
type Parser interface {
	// TODO: fully controlled version
	// setFlagArguments(*flag.FlagSet)
	setFlagArguments()
	validateRequiredArguments() error
}

// Args are arguments that are passed in over the command line
type Args struct {
	Input *string
	Ouput *string
}

func (a *Args) validateRequiredArguments() error {
	if *a.Input == "" {
		flag.PrintDefaults()
		return errRequiredArgumentsNotSet
	}
	return nil
}

func (a *Args) setFlagArguments() {
	a.Input = flag.String("input", "", "Input path. (Required)")
	a.Ouput = flag.String("output", "", "Output path")
}

// Parse will take the user's commandline args and parse them out into a usable
// object
// func Parse(p Parser, cli *flag.FlagSet) error {
func Parse(p Parser) error {
	// p.setFlagArguments(cli)
	p.setFlagArguments()
	flag.Parse()

	return p.validateRequiredArguments()
}
