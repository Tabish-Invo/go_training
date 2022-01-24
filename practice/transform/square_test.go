package transform

import (
	"reflect"
	"testing"
)

func TestTransformSquare(t *testing.T) {

	testSlice := []int{1, 2, 3, 4, 5}
	expectedResult := []int{1, 4, 9, 16, 25}

	result := SquareSlice(testSlice)

	if reflect.DeepEqual(expectedResult, result) {
		t.Log("SquareSlice Passed")
	} else {
		t.Error("Failed")
	}
}
