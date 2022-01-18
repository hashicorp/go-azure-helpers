package tags

func Flatten(input *map[string]string) map[string]interface{} {
	output := make(map[string]interface{})
	if input == nil {
		return output
	}

	for k, v := range *input {
		tagKey := k
		tagValue := v
		output[tagKey] = tagValue
	}

	return output
}
