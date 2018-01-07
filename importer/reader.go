package importer

import (
	"encoding/csv"
	"io"
	"os"
)

// type SSFormat struct {
// 	Question string
// 	OnYomi   string
// 	KunYomi  string
// 	Answer   string
// }
type SSFormat struct {
	Column1 string
	Column2 string
	Column3 string
	Column4 string
}

// NewReader returns a new CSV reader that deliniates on tabs
func NewReader(r io.Reader) (reader *csv.Reader) {
	reader = csv.NewReader(r)
	reader.Comma = '\t'

	return
}

// Parse takes an input file path and attempts to parse the data into the SSFormat.
// A list of all rows parsed are returned, and an error if one occurred.
func Parse(inputPath string) (data []SSFormat, err error) {
	// Load a file.
	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Create a new reader.
	r := NewReader(f)

	// var data []SSFormat
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		// Store record.
		item := SSFormat{
			Column1: record[0],
			Column2: record[1],
			Column3: record[2],
		}
		// May not have already been through the application and have history saved.
		if len(record) > 3 {
			item.Column4 = record[3]
		}
		// // Display record.
		// fmt.Printf("%v\n", item)
		// Append the next item to the list
		data = append(data, item)
	}
	return data, nil
}
