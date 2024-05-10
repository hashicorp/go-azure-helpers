package recaser

import "fmt"

// ValidateID checks that 'input' can be parsed as a valid ID
func ValidateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := reCaseWithIds(v, knownResourceIds, true); err != nil {
		errors = append(errors, err)
	}

	return
}
