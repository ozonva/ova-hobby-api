package utils

import (
	"errors"
	"fmt"
)

// SliceIntsIntoBatches slices a target slice into a slice of slices with the demanded size
func SliceIntsIntoBatches(target []int, batchSize int) ([][]int, error) {
	if target == nil {
		return nil, errors.New("target slice is not initialized")
	}
	targetLen := len(target)

	if targetLen < 1 {
		return nil, errors.New("target slice must contain at least 1 element, yours has 0")
	}
	if batchSize < 1 {
		return nil, fmt.Errorf("batchSize cannot be less than 1, you put %v", batchSize)
	}
	if batchSize > targetLen {
		return nil, fmt.Errorf(
			"batchSize cannot exceed the len of target, you put %v - length of the slice is %v", batchSize, targetLen,
		)
	}

	batchedCapacity := (len(target) + batchSize - 1) / batchSize
	batchedSlice := make([][]int, 0, batchedCapacity)

	for currentIndex := 0; currentIndex < targetLen; currentIndex += batchSize {
		if currentIndex+batchSize > targetLen-1 {
			batchedSlice = append(batchedSlice, target[currentIndex:])
			return batchedSlice, nil
		}
		batchedSlice = append(batchedSlice, target[currentIndex:currentIndex+batchSize])
	}
	return batchedSlice, nil
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
		return nil, errors.New("target slice is not initialized")
	}

	if filterSlice == nil {
		return nil, errors.New("filterSlice is not initialized")
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
