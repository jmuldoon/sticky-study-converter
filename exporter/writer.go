package exporter

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/jmuldoon/sticky-study-converter/importer"
	"github.com/jmuldoon/sticky-study-converter/utility"
)

// NewWriter returns a new CSV writer that deliniates on tabs
func NewWriter(w io.Writer) (writer *csv.Writer) {
	writer = csv.NewWriter(w)
	writer.Comma = '\t'

	return
}

// OutputToFile will create/truncate a file and write the csv data to it
// returning an error if one occurs.
func OutputToFile(outputPath string, data []importer.SSFormat) (err error) {
	// Open file to output to.
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	w := NewWriter(f)

	for _, record := range data {
		d, err := utility.StructToSlice(record)
		if err != nil {
			return err
		}
		w.Write(d)
		w.Flush()
	}

	return nil
}
