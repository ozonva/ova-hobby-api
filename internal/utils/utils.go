package utils

import (
	"fmt"
)

// SliceIntsIntoBatches slices a target slice into a slice of slices with the demanded size
func SliceIntsIntoBatches(target []int, batchSize int) ([][]int, error) {
	if target == nil {
		return nil, fmt.Errorf("target slice is not initialized")
	}
	targetLen := len(target)

	if targetLen < 1 {
		return nil, fmt.Errorf("target slice must contain at least 1 element, yours has %v", batchSize)
	}
	if batchSize < 1 {
		return nil, fmt.Errorf("batchSize cannot be less than 1, you put %v", batchSize)
	}
	if batchSize > targetLen {
		return nil, fmt.Errorf(
			"batchSize cannot exceed the len of target, you put %v - length of the slice is %v", batchSize, targetLen,
		)
	}

	maxSize := targetLen / batchSize
	if targetLen%batchSize != 0 {
		maxSize += 1
	}

	newSlice := make([][]int, 0, maxSize)
	for currentIndex := 0; currentIndex < targetLen; currentIndex += batchSize {
		if currentIndex+batchSize > targetLen-1 {
			newSlice = append(newSlice, target[currentIndex:])
			return newSlice, nil
		}
		newSlice = append(newSlice, target[currentIndex:currentIndex+batchSize])
	}
	return newSlice, nil
}

// SwapKeysValues swaps keys with values in a map, error if there are multiple the same values
func SwapKeysValues(target map[string]int) (map[int]string, error) {
	if target == nil {
		return nil, fmt.Errorf("target map is not initialized")
	}

	output := make(map[int]string, len(target))

	for originalKey, originalValue := range target {
		if _, ok := output[originalValue]; ok {
			return nil, fmt.Errorf("more than one key {%v} in the resulting map", originalValue)
		}
		output[originalValue] = originalKey
	}
	return output, nil
}

var emptyValue struct{}

func convertSliceToMap(slice []int) map[int]struct{} {
	stopValuesSet := make(map[int]struct{}, len(slice))
	for _, stopValue := range slice {
		stopValuesSet[stopValue] = emptyValue
	}
	return stopValuesSet
}

// FilterSliceInts filters a slice against values of an another slice and returns a new slice
func FilterSliceInts(target []int, filterSlice []int) ([]int, error) {
	if target == nil {
		return nil, fmt.Errorf("target slice is not initialized")
	}

	stopValuesSet := convertSliceToMap(filterSlice)

	filteredOutput := make([]int, 0)
	for _, number := range target {
		if _, ok := stopValuesSet[number]; !ok {
			filteredOutput = append(filteredOutput, number)
		}
	}
	return filteredOutput, nil
}
