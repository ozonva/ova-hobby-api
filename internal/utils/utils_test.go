package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceIntsIntoBatches(t *testing.T) {
	tests := []struct {
		inputSlice         []int
		inputBatchSize     int
		expected           [][]int
		expectedErrMessage string
	}{
		{[]int{0}, 1, [][]int{{0}}, ""},
		{[]int{0, 1}, 2, [][]int{{0, 1}}, ""},
		{[]int{0, 1, 2, 3}, 2, [][]int{{0, 1}, {2, 3}}, ""},
		{[]int{0, 1, 2, 3}, 3, [][]int{{0, 1, 2}, {3}}, ""},
		{[]int{0, 1, 2, 3}, 4, [][]int{{0, 1, 2, 3}}, ""},

		{nil, 1, nil, "target slice is not initialized"},
		{[]int{}, 1, nil, "target slice must contain at least 1 element, yours has 0"},
		{[]int{0}, 2, nil, "batchSize cannot exceed the len of target, you put 2 - length of the slice is 1"},
		{[]int{0}, 0, nil, "batchSize cannot be less than 1, you put 0"},
		{[]int{0}, -5, nil, "batchSize cannot be less than 1, you put -5"},
	}

	for _, testCase := range tests {
		testName := fmt.Sprintf("%d,%d --> %d", testCase.inputSlice, testCase.inputBatchSize, testCase.expected)

		t.Run(testName, func(t *testing.T) {
			actual, err := SliceIntsIntoBatches(testCase.inputSlice, testCase.inputBatchSize)

			if testCase.expectedErrMessage != "" {
				require.EqualError(t, err, testCase.expectedErrMessage)
			}
			require.Equal(t, testCase.expected, actual, "Slices must be the same")
		})
	}
}

func TestSwapKeysValues(t *testing.T) {
	tests := []struct {
		inputMap           map[string]int
		expected           map[int]string
		expectedErrMessage string
	}{
		{map[string]int{"some": 1}, map[int]string{1: "some"}, ""},
		{map[string]int{"some": 1, "Any": 33}, map[int]string{1: "some", 33: "Any"}, ""},

		{nil, nil, "target map is not initialized"},
		{map[string]int{"some": 1, "any": 1}, nil, "more than one key {1} in the resulting map"},
	}

	for _, testCase := range tests {
		testName := fmt.Sprintf("%v --> %v", testCase.inputMap, testCase.expected)

		t.Run(testName, func(t *testing.T) {
			actual, err := SwapKeysValues(testCase.inputMap)

			if testCase.expectedErrMessage != "" {
				require.EqualError(t, err, testCase.expectedErrMessage)
			}
			require.Equal(t, testCase.expected, actual, "Maps must be the same")
		})
	}
}

func TestFilterSliceInts(t *testing.T) {
	tests := []struct {
		inputSlice         []int
		inputFilterSlice   []int
		expected           []int
		expectedErrMessage string
	}{
		{[]int{0}, []int{}, []int{0}, ""},
		{[]int{0}, []int{1}, []int{0}, ""},
		{[]int{0}, []int{1, 1}, []int{0}, ""},
		{[]int{0, 1}, []int{1, 1}, []int{0}, ""},
		{[]int{0, 1, 2, 3}, []int{1, 3}, []int{0, 2}, ""},

		{nil, []int{}, nil, "target slice is not initialized"},
		{[]int{}, nil, nil, "filterSlice is not initialized"},
	}

	for _, testCase := range tests {
		testName := fmt.Sprintf("%d,%d --> %d", testCase.inputSlice, testCase.inputFilterSlice, testCase.expected)

		t.Run(testName, func(t *testing.T) {
			actual, err := FilterSliceInts(testCase.inputSlice, testCase.inputFilterSlice)

			if testCase.expectedErrMessage != "" {
				require.EqualError(t, err, testCase.expectedErrMessage)
			}

			require.Equal(t, testCase.expected, actual, "Slices must be the same")
		})
	}
}
