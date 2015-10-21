package csvmap

import (
	"reflect"
	"testing"
)

func TestHeaderSliceCreation(t *testing.T) {
	result := New("test_csvs/test.csv").Headers

	t.Logf("Generated Result: \n%v", result)

	expectedResult := []string{"first", "second", "third"}
	if result == nil {
		t.Errorf("Test CSV headers returned nil\n", expectedResult)
	} else {
		t.Log("Test CSV headers generated and are not nil")
	}

	if len(expectedResult) > len(result) {
		t.Errorf("Resulting header slice length is %v, which is %v less than expected\n", len(result), len(expectedResult)-len(result))
	} else if len(expectedResult) < len(result) {
		t.Errorf("Resulting header slice length is %v, which is %v more than expected\n", len(result), len(result)-len(expectedResult))
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Resulting header slice does not equal the expected result:\n\n%v\n\t!=\n%v\n", result, expectedResult)
	}
}

func TestHeaderMapCreation(t *testing.T) {
	result := New("test_csvs/test.csv").HeaderIndexMap
	expectedResult := map[string]int{
		"first":  0,
		"second": 1,
		"third":  2,
	}

	if result == nil {
		t.Errorf("Test CSV header map returned nil\n", expectedResult)
	}

	for key, value := range expectedResult {
		if _, ok := result[key]; !ok {
			t.Errorf("The following key is not located in the resulting header map: %v\n", key)
		} else if result[key] != value {
			t.Errorf("The generated header map has incorrect values for this key: %v\n", key)
		}
	}
}

func BenchmarkSmallFileHandling(b *testing.B) {
	New("test_csvs/test.csv")
}

func BenchmarkLargeFileHandling(b *testing.B) {
	New("test_csvs/FakeNameGeneratorFile.csv")
}
