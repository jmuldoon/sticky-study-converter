package arguments

import "flag"

const (
	TestInput  = "C:/test/input/path/file.txt"
	TestOutput = "C:/test/output/path"
	usage      = "TestUsageStatement"
)

type MockArgs struct {
	Input  *string
	Output *string
	New    *bool
}

type Expected struct {
	Value interface{}
	Error error
}

type Tested struct {
	Value Parser
	Error error
}

// TODO: make sure to test for a false flag as well. probably will be better to
// rewrite how I get the flags setup before wasting time here.
var testParse = []struct {
	Description string
	Expected
	Tested
}{
	{
		Description: `CLI Arguments defined, Successful New:=true`,
		Expected:    Expected{Error: nil},
		Tested: Tested{Value: &MockArgs{
			Input:  flag.String("input", TestInput, usage),
			Output: flag.String("output", TestOutput, usage),
			New:    flag.Bool("new", true, usage),
		}},
	},
}
