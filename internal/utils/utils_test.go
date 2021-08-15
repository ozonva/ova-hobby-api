package utils

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/ozonva/ova-hobby-api/pkg/hobby"
	"github.com/stretchr/testify/require"
)

func TestHobbiesToMap(t *testing.T) {
	h1 := hobby.NewHobby("SUP", 0, hobby.Sports)
	h2 := hobby.NewHobby("Collecting stamps", 1, hobby.Collecting)

	h3 := h2
	h3.ID = h1.ID

	tests := []struct {
		inputSlice         []hobby.Hobby
		expected           map[uuid.UUID]hobby.Hobby
		expectedErrMessage string
	}{
		{[]hobby.Hobby{}, map[uuid.UUID]hobby.Hobby{}, ""},
		{[]hobby.Hobby{h1}, map[uuid.UUID]hobby.Hobby{h1.ID: h1}, ""},
		{[]hobby.Hobby{h1, h2}, map[uuid.UUID]hobby.Hobby{h1.ID: h1, h2.ID: h2}, ""},

		{nil, nil, "hobbies slice is not initialized"},
		{[]hobby.Hobby{h1, h3}, nil, fmt.Sprintf("more than one key {%s} in the resulting map", h1.ID)},
	}

	for _, testCase := range tests {
		testName := fmt.Sprintf("%v --> %v", testCase.inputSlice, testCase.expected)

		t.Run(testName, func(t *testing.T) {
			actual, err := HobbiesToMap(testCase.inputSlice)

			if testCase.expectedErrMessage != "" {
				require.EqualError(t, err, testCase.expectedErrMessage)
			}
			require.Equal(t, testCase.expected, actual, "Maps must be the same")
		})
	}
}

func TestSliceHobbiesIntoBatches(t *testing.T) {
	h1 := hobby.NewHobby("SUP", 0, hobby.Sports)
	h2 := hobby.NewHobby("Collecting stamps", 1, hobby.Collecting)
	h3 := hobby.NewHobby("Cooking", 612, hobby.Domestic)
	h4 := hobby.NewHobby("Hikking", 98643, hobby.Outdoors)

	tests := []struct {
		inputSlice         []hobby.Hobby
		inputBatchSize     int
		expected           [][]hobby.Hobby
		expectedErrMessage string
	}{
		{[]hobby.Hobby{h1}, 1, [][]hobby.Hobby{{h1}}, ""},
		{[]hobby.Hobby{h1, h2}, 1, [][]hobby.Hobby{{h1}, {h2}}, ""},
		{[]hobby.Hobby{h1, h2}, 2, [][]hobby.Hobby{{h1, h2}}, ""},
		{[]hobby.Hobby{h1, h2, h3, h4}, 2, [][]hobby.Hobby{{h1, h2}, {h3, h4}}, ""},
		{[]hobby.Hobby{h1, h2, h3, h4}, 3, [][]hobby.Hobby{{h1, h2, h3}, {h4}}, ""},

		{nil, 1, nil, "target slice is not initialized"},
		{[]hobby.Hobby{}, 1, nil, "target slice must contain at least 1 element, yours has 0"},
		{[]hobby.Hobby{{}}, 2, nil, "batchSize cannot exceed the len of target, you put 2 - length of the slice is 1"},
		{[]hobby.Hobby{{}}, 0, nil, "batchSize cannot be less than 1, you put 0"},
		{[]hobby.Hobby{{}}, -5, nil, "batchSize cannot be less than 1, you put -5"},
	}

	for _, testCase := range tests {
		testName := fmt.Sprintf("%v,%v --> %v", testCase.inputSlice, testCase.inputBatchSize, testCase.expected)

		t.Run(testName, func(t *testing.T) {
			actual, err := SliceHobbiesIntoBatches(testCase.inputSlice, testCase.inputBatchSize)

			if testCase.expectedErrMessage != "" {
				require.EqualError(t, err, testCase.expectedErrMessage)
			}
			require.Equal(t, testCase.expected, actual, "Slices must be the same")
		})
	}
}

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
