package tfdata

import (
	"strings"

	"github.com/hashicorp/go-azure-helpers/lang/types"
)

func FromSlice[T types.Primary](input *[]T) []interface{} {
	result := make([]interface{}, 0)
	if input != nil {
		for _, item := range *input {
			result = append(result, item)
		}
	}
	return result
}

func FromRangeSlice[T types.Primary](input *[][]T) [][]interface{} {
	result := make([][]interface{}, 0)
	if input != nil {
		for _, item := range *input {
			result = append(result, FromSlice(&item))
		}
	}
	return result
}

func FromMapOfPtr[K types.Primary, V types.Primary](input map[K]*V) map[K]interface{} {
	result := make(map[K]interface{})
	for k, v := range input {
		if v == nil {
			var value V
			result[k] = value
		} else {
			result[k] = *v
		}
	}
	return result
}

func FromMap[K types.Primary, V types.Primary](input map[K]V) map[K]interface{} {
	result := make(map[K]interface{})
	for k, v := range input {
		result[k] = v
	}
	return result
}

func FromStringWithDelimiter(input *string, delimiter string) []interface{} {
	result := make([]interface{}, 0)
	if input != nil {
		inputStrings := strings.Split(*input, delimiter)
		for _, item := range inputStrings {
			result = append(result, item)
		}
	}
	return result
}
