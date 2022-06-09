package tfdata

import (
	"strconv"
	"strings"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/types"
)

func ToSlice[T types.Primary](input []interface{}) *[]T {
	result := make([]T, 0)
	for _, item := range input {
		if item != nil {
			result = append(result, item.(T))
		} else {
			var v T
			result = append(result, v)
		}
	}
	return &result
}

func ToRangeSlice[T types.Primary](input []interface{}) *[][]T {
	result := make([][]T, 0)
	for _, item := range input {
		if item != nil {
			result = append(result, *ToSlice[T](item.([]interface{})))
		}
	}
	return &result
}

func ToMap[K types.Primary, V types.Primary](input map[K]interface{}) map[K]V {
	result := make(map[K]V)
	for k, v := range input {
		result[k] = v.(V)
	}
	return result
}

func ToMapOfPtr[K types.Primary, V types.Primary](input map[K]interface{}) map[K]*V {
	result := make(map[K]*V)
	for k, v := range input {
		result[k] = pointer.From(v.(V))
	}
	return result
}

func ToStringWithDelimiter(input []interface{}, delimiter string) *string {
	result := make([]string, 0)
	for _, item := range input {
		if item != nil {
			switch item.(type) {
			case string:
				result = append(result, item.(string))
			case int:
				result = append(result, strconv.Itoa(item.(int)))
			default:
				// TODO: should we handle other types?
				result = append(result, "")
			}
		} else {
			result = append(result, "")
		}
	}
	return pointer.FromString(strings.Join(result, delimiter))
}
