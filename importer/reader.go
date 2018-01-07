package importer

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
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
func NewReader(r io.Reader, delim rune) (reader *csv.Reader) {
	reader = csv.NewReader(r)
	reader.Comma = delim

	return
}

// Parse takes an input file path and attempts to parse the data into the SSFormat.
// A list of all rows parsed are returned, and an error if one occurred.
func Parse(inputPath string, newToSS bool) (data []SSFormat, err error) {
	// Load a file.
	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Create a new reader.
	var r *csv.Reader
	if newToSS {
		r = NewReader(f, ' ')
	} else {
		r = NewReader(f, '\t')
	}

	// remove formatting from the file to make things cleaner.
	replacer := strings.NewReplacer("[", "", "]", "", "/", "")
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		// Store record, if new swap kanji to second, else keep
		var item SSFormat

		if newToSS {
			item = SSFormat{
				Column1: replacer.Replace(record[1]),
				Column2: record[0] + " " + replacer.Replace(strings.Join(record[2:], " ")),
			}
		} else {
			item = SSFormat{
				Column1: record[0],
				Column2: record[1],
			}
			// May not have already been through the application and have history saved.
			if len(record) > 2 {
				item.Column3 = record[2]
			}
		}

		// Append the next item to the list
		data = append(data, item)
	}

	return data, nil
}
