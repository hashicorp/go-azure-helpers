package pointer

func ToBoolPointer(input bool) *bool {
	return &input
}

func ToFloatPointer(input float64) *float64 {
	return &input
}

func ToIntPointer(input int) *int {
	return &input
}

func ToInt64Pointer(input int64) *int64 {
	return &input
}

func ToMapOfStringInterfacesPointer(input map[string]interface{}) *map[string]interface{} {
	return &input
}

func ToMapOfStringStringsPointer(input map[string]string) *map[string]string {
	return &input
}

func ToSliceOfStringsPointer(input []string) *[]string {
	return &input
}

func ToStringPointer(input string) *string {
	return &input
}
