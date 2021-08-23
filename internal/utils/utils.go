package utils

import (
	"errors"
	"fmt"

	"github.com/ozonva/ova-hobby-api/pkg/models"

	"github.com/google/uuid"
)

// HobbiesToMap creates a map with a UUID key and a Hobby as a value
func HobbiesToMap(hobbies []models.Hobby) (map[uuid.UUID]models.Hobby, error) {
	if hobbies == nil {
		return nil, errors.New("hobbies slice is not initialized")
	}
	mapHobbies := make(map[uuid.UUID]models.Hobby)

	for _, hobby := range hobbies {
		if _, ok := mapHobbies[hobby.ID]; ok {
			return nil, fmt.Errorf("more than one key {%s} in the resulting map", hobby.ID)
		}
		mapHobbies[hobby.ID] = hobby
	}
	return mapHobbies, nil
}

// SliceHobbiesIntoBatches slices a target slice into a slice of slices with the demanded size
func SliceHobbiesIntoBatches(target []models.Hobby, batchSize int) ([][]models.Hobby, error) {
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
	batchedSlice := make([][]models.Hobby, 0, batchedCapacity)

	for currentIndex := 0; currentIndex < targetLen; currentIndex += batchSize {
		if currentIndex+batchSize > targetLen-1 {
			batchedSlice = append(batchedSlice, target[currentIndex:])
			return batchedSlice, nil
		}
		batchedSlice = append(batchedSlice, target[currentIndex:currentIndex+batchSize])
	}
	return batchedSlice, nil
}

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

func convertSliceToMap(slice []int) map[int]struct{} {
	stopValuesSet := make(map[int]struct{}, len(slice))
	for _, stopValue := range slice {
		stopValuesSet[stopValue] = struct{}{}
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
