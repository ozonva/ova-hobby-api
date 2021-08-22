package utils_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/ozonva/ova-hobby-api/internal/utils"
	"github.com/ozonva/ova-hobby-api/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestHobbiesToMap(t *testing.T) {
	h1 := models.NewHobby("SUP", 0, models.Sports)
	h2 := models.NewHobby("Collecting stamps", 1, models.Collecting)

	h3 := h2
	h3.ID = h1.ID

	tests := []struct {
		inputSlice         []models.Hobby
		expected           map[uuid.UUID]models.Hobby
		expectedErrMessage string
	}{
		{[]models.Hobby{}, map[uuid.UUID]models.Hobby{}, ""},
		{[]models.Hobby{h1}, map[uuid.UUID]models.Hobby{h1.ID: h1}, ""},
		{[]models.Hobby{h1, h2}, map[uuid.UUID]models.Hobby{h1.ID: h1, h2.ID: h2}, ""},

		{nil, nil, "hobbies slice is not initialized"},
		{[]models.Hobby{h1, h3}, nil, fmt.Sprintf("more than one key {%s} in the resulting map", h1.ID)},
	}

	for _, testCase := range tests {
		testName := fmt.Sprintf("%v --> %v", testCase.inputSlice, testCase.expected)

		t.Run(testName, func(t *testing.T) {
			actual, err := utils.HobbiesToMap(testCase.inputSlice)

			if testCase.expectedErrMessage != "" {
				require.EqualError(t, err, testCase.expectedErrMessage)
			}
			require.Equal(t, testCase.expected, actual, "Maps must be the same")
		})
	}
}

func TestSliceHobbiesIntoBatches(t *testing.T) {
	h1 := models.NewHobby("SUP", 0, models.Sports)
	h2 := models.NewHobby("Collecting stamps", 1, models.Collecting)
	h3 := models.NewHobby("Cooking", 612, models.Domestic)
	h4 := models.NewHobby("Hikking", 98643, models.Outdoors)

	tests := []struct {
		inputSlice         []models.Hobby
		inputBatchSize     int
		expected           [][]models.Hobby
		expectedErrMessage string
	}{
		{[]models.Hobby{h1}, 1, [][]models.Hobby{{h1}}, ""},
		{[]models.Hobby{h1, h2}, 1, [][]models.Hobby{{h1}, {h2}}, ""},
		{[]models.Hobby{h1, h2}, 2, [][]models.Hobby{{h1, h2}}, ""},
		{[]models.Hobby{h1, h2, h3, h4}, 2, [][]models.Hobby{{h1, h2}, {h3, h4}}, ""},
		{[]models.Hobby{h1, h2, h3, h4}, 3, [][]models.Hobby{{h1, h2, h3}, {h4}}, ""},

		{nil, 1, nil, "target slice is not initialized"},
		{[]models.Hobby{}, 1, nil, "target slice must contain at least 1 element, yours has 0"},
		{[]models.Hobby{{}}, 2, nil, "batchSize cannot exceed the len of target, you put 2 - length of the slice is 1"},
		{[]models.Hobby{{}}, 0, nil, "batchSize cannot be less than 1, you put 0"},
		{[]models.Hobby{{}}, -5, nil, "batchSize cannot be less than 1, you put -5"},
	}

	for _, testCase := range tests {
		testName := fmt.Sprintf("%v,%v --> %v", testCase.inputSlice, testCase.inputBatchSize, testCase.expected)

		t.Run(testName, func(t *testing.T) {
			actual, err := utils.SliceHobbiesIntoBatches(testCase.inputSlice, testCase.inputBatchSize)

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
			actual, err := utils.SliceIntsIntoBatches(testCase.inputSlice, testCase.inputBatchSize)

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
			actual, err := utils.SwapKeysValues(testCase.inputMap)

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
			actual, err := utils.FilterSliceInts(testCase.inputSlice, testCase.inputFilterSlice)

			if testCase.expectedErrMessage != "" {
				require.EqualError(t, err, testCase.expectedErrMessage)
			}

			require.Equal(t, testCase.expected, actual, "Slices must be the same")
		})
	}
}
