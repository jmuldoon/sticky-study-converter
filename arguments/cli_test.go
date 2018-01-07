package arguments

import (
	"testing"
)

const (
	targetTestVersion = 0
)

func TestVersionValidation(t *testing.T) {
	t.Parallel() // indicator that it can be tested in parallel
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v\n", testVersion, targetTestVersion)
	}
}

// TODO: actually cleanup these two functions and increase coverage.
func (a *MockArgs) setFlagArguments() {}

func (a *MockArgs) validateRequiredArguments() error {
	if *a.Input == "" {
		return errRequiredArgumentsNotSet
	}
	return nil
}

func TestParse(t *testing.T) {
	t.Parallel() // indicator that it can be tested in parallel
	for _, test := range testParse {
		// commandLine := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		// observed := Parse(test.Tested.Value, commandLine)
		observed := Parse(test.Tested.Value)
		t.Logf("Running test for `%s`\n", test.Description)
		if observed != test.Expected.Error {
			t.Fatalf("%s(%v):\n"+
				"Brief (%s)\n"+
				"Observed: %t\n"+
				"Expected: %t\n",
				"TestParse", test.Tested.Value,
				test.Description, observed, test.Expected.Error)
		}
	}
}
