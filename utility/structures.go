package utility

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	"github.com/jmuldoon/utility"
)

// StructToSlice takes a struct and transforms it to a slice of the elements
// contained within it.
// TODO: Move this into the utilities package
func StructToSlice(data interface{}) (xdata []string, err error) {
	inspected := utility.InspectStruct(data)
	// TODO: make this work later for something more complex.
	if len(inspected) != 1 {
		return nil, fmt.Errorf("StructToSlice: expected interface is nested too deep or empty.")
	}

	// Allocate the memory needed for the slice.
	xdata = make([]string, len(inspected[0]))

	// Compile regex to find the number for the column.
	re := regexp.MustCompile("[0-9]+")
	// Loop over all values in the structure that was inspected and pull the fields
	for i, d := range inspected[0] {
		v := re.FindAllString(i, 1)
		index, _ := strconv.Atoi(v[0])
		xdata[index-1] = reflect.ValueOf(d).Field(0).Interface().(string)
	}
	return xdata, nil
}
