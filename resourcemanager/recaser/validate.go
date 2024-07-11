package recaser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// ValidateID checks that 'input' can be parsed as a valid ID
func ValidateID(input string) []error {
	return validateId(input, knownResourceIds)
}

func validateId(input string, ids map[string]resourceids.ResourceId) []error {
	var errors []error
	var id resourceids.ResourceId
	key, ok := buildInputKey(input)
	if ok {
		id = ids[*key]
		if id != nil {
			_, _, err := parseId(id, input, false)
			if err != nil {
				errors = append(errors, err)
			}
		}
	}

	// if we can't find a matching id, validate case of known segments instead
	if id == nil {
		err := validateKnownSegments(input)
		if err != nil {
			errors = append(errors, err...)
		}
	}

	return errors
}

func validateKnownSegments(input string) []error {
	var errors []error
	knownSegments := []string{
		"/subscriptions/",
		"/resourceGroups/",
		"/managementGroups/",
		"/tenants/",
	}

	for _, segment := range knownSegments {
		err := validateSegment(input, segment)
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

func validateSegment(input string, segment string) error {
	if strings.Contains(strings.ToLower(input), strings.ToLower(segment)) {
		if !strings.Contains(input, segment) {
			re := regexp.MustCompile(fmt.Sprintf("(?i)%s", segment))
			return fmt.Errorf("expected %s but got %s", segment, re.FindAllStringSubmatch(input, 0)[0])
		}
	}
	return nil
}
