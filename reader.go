package cartogopher

import (
	"encoding/csv"
	"fmt"
	"io"
)

// MapReader contains all our necessary data for the various methods to function
type MapReader struct {
	Headers        []string
	HeaderIndexMap map[string]int
	Reader         *csv.Reader
}

// CreateHeaderIndexMap creates a map of header strings to their indices in the array generated by
// encoding/csv's reader. For instance, if your CSV file looks something like
// this:
// 	---------------------
// 	| one | two | three |
// 	---------------------
// 	|  A  |  B  |   C   |
// 	---------------------
// Go's generated array for the header row will be [ "one", "two", "three" ].
// Cartogopher's generated map for the header row will be { "one": 1, "two": 2, "three": 3 }
func (m *MapReader) CreateHeaderIndexMap() {
	headerIndexMap := make(map[string]int, len(m.Headers))

	for index, header := range m.Headers {
		headerIndexMap[header] = index
	}

	m.HeaderIndexMap = headerIndexMap
}

// CreateRowMap takes a given CSV array and returns a map of column names to the values contained therein.
// For instance, if your CSV file looks something like this:
// 	---------------------
// 	| one | two | three |
// 	---------------------
// 	|  A  |  B  |   C   |
// 	---------------------
// The return result will be:
// 	{
// 		"one": "A",
// 		"two": "B",
// 		"three": "C",
// 	}
// Note that this requires the HeaderIndexMap to be created and not a null value.
func (m MapReader) CreateRowMap(csvRow []string) map[string]string {
	result := map[string]string{}
	for header, index := range m.HeaderIndexMap {
		result[header] = csvRow[index]
	}

	return result
}

// Read mimics the built-in CSV reader Read method, returning one row of
// the CSV. The only difference here being that obviously we return a
// map instead of a slice.
func (m MapReader) Read() (map[string]string, error) {
	csvRow, err := m.Reader.Read()
	if err != nil {
		return nil, err
	}
	return m.CreateRowMap(csvRow), nil
}

// ReadAll mimics the built-in CSV reader ReadAll method
func (m MapReader) ReadAll() ([]map[string]string, error) {
	records, err := m.Reader.ReadAll()
	if err != nil {
		return nil, err
	}
	results := []map[string]string{}
	for _, record := range records {
		results = append(results, m.CreateRowMap(record))
	}
	return results, nil
}

// NewReader returns a new MapReader struct. It can be created the same way
// a regular CSV file is created, by providing it with a reference to a file
// reader, ideally one that points to a CSV file.  I'm using an interface
// here so that, should the need arise, you can provide your CSV to the
// package in a variety of non-file based ways. Note that here we read the
// first row of the file without setting any non-standard values for the
// CSV package's Reader struct. If it becomes apparent that the ability to
// change these parameters is vital, then I'm more than happy to figure out
// an idiomatic way to accomplish that task.
func NewReader(file io.Reader) (*MapReader, error) {
	// Create our reader
	reader := csv.NewReader(file)

	// Create our resulting struct
	output := &MapReader{}

	inputHeaders, err := reader.Read()
	if err != nil {
		return nil, err
	}

	// Use our methods (defined above) to populate our struct fields
	output.Headers = inputHeaders
	output.Reader = reader
	output.CreateHeaderIndexMap()

	if output.HeaderIndexMap == nil {
		return nil, fmt.Errorf("error assigning header to index map for CSV file")
	}

	return output, nil
}
