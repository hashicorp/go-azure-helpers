package pointer

import "github.com/hashicorp/go-azure-helpers/lang/types"

// ToBool turns a pointer to a bool into a bool, returning the default value for a bool if it's nil
func ToBool(input *bool) bool {
	if input != nil {
		return *input
	}

	return false
}

// ToFloat64 turns a pointer to a float64 into a float64, returning the default value for a float64 if it's nil
func ToFloat64(input *float64) float64 {
	if input != nil {
		return *input
	}

	return 0.0
}

// ToInt turns a pointer to an int into an int, returning the default value for an int if it's nil
func ToInt(input *int) int {
	if input != nil {
		return *input
	}

	return 0
}

// ToInt64 turns a pointer to an int64 into an int64, returning the default value for an int64 if it's nil
func ToInt64(input *int64) int64 {
	if input != nil {
		return *input
	}

	return 0
}

// ToMapOfStringInterfaces turns a pointer to a map[string]interface{} into a map[string]interface{}
// returning an empty map[string]interface{} if it's nil
func ToMapOfStringInterfaces(input *map[string]interface{}) map[string]interface{} {
	if input != nil {
		return *input
	}

	return map[string]interface{}{}
}

// ToMapOfStringStrings turns a pointer to a map[string]string into a map[string]string returning
// an empty map[string]string if it's nil
func ToMapOfStringStrings(input *map[string]string) map[string]string {
	if input != nil {
		return *input
	}

	return map[string]string{}
}

// ToSliceOfStrings turns a pointer to a slice of strings into a slice of strings returning
// an empty slice of strings if it's nil
func ToSliceOfStrings(input *[]string) []string {
	if input != nil {
		return *input
	}

	return []string{}
}

// ToString turns a pointer to a string into a string, returning an empty string if it's nil
func ToString(input *string) string {
	if input != nil {
		return *input
	}

	return ""
}

// ToPrimary turns a pointer to a primary into a primary, returning its zero value if it's nul
func ToPrimary[T types.Primary](input *T) T {
	if input != nil {
		return *input
	}
	var v T
	return v
}

// ToMap turns a pointer to a map into a map, returning its zero value if it's nul
func ToMap[K types.Primary, T types.Primary | interface{}](input *map[K]T) map[K]T {
	if input != nil {
		return *input
	}
	return map[K]T{}
}

// ToSlice turns a pointer to a slice into a slice, returning its zero value if it's nul
func FromSlice[T types.Primary | interface{}](input *[]T) []T {
	if input != nil {
		return *input
	}
	return []T{}
}
