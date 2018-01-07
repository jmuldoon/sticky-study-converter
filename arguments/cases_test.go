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
}

type Expected struct {
	Value interface{}
	Error error
}

type Tested struct {
	Value Parser
	Error error
}

var testParse = []struct {
	Description string
	Expected
	Tested
}{
	{
		Description: `CLI Arguments defined, Successful`,
		Expected:    Expected{Error: nil},
		Tested: Tested{Value: &MockArgs{
			Input:  flag.String("input", TestInput, usage),
			Output: flag.String("output", TestOutput, usage),
		}},
	},
}
