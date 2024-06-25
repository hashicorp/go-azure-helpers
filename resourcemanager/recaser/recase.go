// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package recaser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// ReCase tries to determine the type of Resource ID defined in `input` to be able to re-case it from
func ReCase(input string) string {
	return reCaseWithIds(input, knownResourceIds)
}

// reCaseWithIds tries to determine the type of Resource ID defined in `input` to be able to re-case it based on an input list of Resource IDs
func reCaseWithIds(input string, ids map[string]resourceids.ResourceId) string {
	output := input
	reCased := false
	var parseError error

	key, ok := buildInputKey(input)
	if ok {
		id := ids[*key]
		if id != nil {
			output, parseError = parseId(id, input)
			if parseError == nil {
				reCased = true
			}
		}
	}

	// if we can't find a matching id re-case these known segments
	if !reCased {

		segmentsToFix := []string{
			"/subscriptions/",
			"/resourceGroups/",
			"/managementGroups/",
			"/tenants/",
		}

		for _, segment := range segmentsToFix {
			output = fixSegment(output, segment)
		}
	}

	return output
}

func RecaseKnownId(input string) (*string, error) {
	return reCaseKnownId(input, knownResourceIds)
}

func reCaseKnownId(input string, ids map[string]resourceids.ResourceId) (*string, error) {
	output := input

	key, ok := buildInputKey(input)
	if ok {
		id := ids[*key]
		if id != nil {
			var parseError error
			output, parseError = parseId(id, input)
			if parseError != nil {
				return &output, fmt.Errorf("fixing case for ID '%s': %+v", input, parseError)
			}
		}
	} else {
		return nil, fmt.Errorf("could not determine ID type for '%s', or ID type not supported", input)
	}

	return &output, nil
}

// parseId uses the specified ResourceId to parse the input and returns the id string with correct casing
func parseId(id resourceids.ResourceId, input string) (string, error) {

	// we need to take a local copy of id to work against else we're mutating the original
	localId := id

	parser := resourceids.NewParserFromResourceIdType(localId)
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return input, err
	}

	if err = id.FromParseResult(*parsed); err != nil {
		return input, err
	}
	input = id.ID()

	return input, err
}

// fixSegment searches the input id string for a specified segment case-insensitively
// and returns the input string with the casing corrected on the segment
func fixSegment(input, segment string) string {
	if strings.Contains(strings.ToLower(input), strings.ToLower(segment)) {
		re := regexp.MustCompile(fmt.Sprintf("(?i)%s", segment))
		input = re.ReplaceAllString(input, segment)
	}
	return input
}

// buildInputKey takes an input id string and removes user-specified values from it
// so it can be used as a key to extract the correct id from knownResourceIds
func buildInputKey(input string) (*string, bool) {

	// Attempt to determine if this is just missing a leading slash and prepend it if it seems to be
	if !strings.HasPrefix(input, "/") {
		if len(input) == 0 || !strings.Contains(input, "/") {
			return nil, false
		}

		input = "/" + input
	}

	output := ""

	segments := strings.Split(input, "/")
	// iterate through the segments extracting any that are not user inputs
	// and append them together to make a key
	// eg "/subscriptions/1111/resourceGroups/group1/providers/Microsoft.BotService/botServices/botServiceValue" will become:
	// "/subscriptions//resourceGroups//providers/Microsoft.BotService/botServices/"
	if len(segments)%2 != 0 {
		for i := 1; len(segments) > i; i++ {
			if i%2 != 0 {
				key := segments[i]
				output = fmt.Sprintf("%s/%s/", output, key)

				// if the current segment is a providers segment, then we should append the next segment to the key
				// as this is not a user input segment
				if strings.EqualFold(key, "providers") && len(segments) >= i+2 {
					value := segments[i+1]
					output = fmt.Sprintf("%s%s", output, value)
				}
			}
		}
	}
	output = strings.ToLower(output)
	return &output, true
}
