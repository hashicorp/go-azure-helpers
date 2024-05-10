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
func ReCase(input string, insensitively bool) string {
	recasedId, _ := reCaseWithIds(input, knownResourceIds, insensitively)
	return recasedId
}

// reCaseWithIds tries to determine the type of Resource ID defined in `input` to be able to re-case it based on an input list of Resource IDs
func reCaseWithIds(input string, ids map[string]resourceids.ResourceId, insensitively bool) (string, error) {
	output := input
	recased := false
	var err error

	key, ok := buildInputKey(input)
	if ok {
		id := ids[*key]
		if id != nil {
			output, recased, err = parseId(id, input, insensitively)
			if err != nil {
				return output, err
			}
		}
	}

	// if we can't find a matching id recase these known segments
	if !recased {

		segmentsToFix := []string{
			"/subscriptions/",
			"/resourceGroups/",
			"/managementGroups/",
			"/tenants/",
		}

		for _, segment := range segmentsToFix {
			fixedSegment, err := fixSegment(output, segment, insensitively)
			output = fixedSegment
			if !insensitively {
				return output, err
			}
		}
	}

	return output, nil
}

// parseId uses the specified ResourceId to parse the input and returns the id string with correct casing
func parseId(id resourceids.ResourceId, input string, insensitively bool) (string, bool, error) {

	// we need to take a local copy of id to work against else we're mutating the original
	localId := id

	parser := resourceids.NewParserFromResourceIdType(localId)
	parsed, err := parser.Parse(input, insensitively)
	if err != nil {
		return input, false, err
	}

	if err = id.FromParseResult(*parsed); err != nil {
		return input, false, err
	}
	input = id.ID()

	return input, true, err
}

// fixSegment searches the input id string for a specified segment case-insensitively
// and returns the input string with the casing corrected on the segment
func fixSegment(input, segment string, insensitively bool) (string, error) {
	if strings.Contains(strings.ToLower(input), strings.ToLower(segment)) {
		re := regexp.MustCompile(fmt.Sprintf("(?i)%s", segment))
		if !insensitively && !strings.Contains(input, segment) {
			return input, fmt.Errorf("expected %s but got %s", segment, re.FindAllStringSubmatch(input, 0)[0])
		}
		input = re.ReplaceAllString(input, segment)
	}
	return input, nil
}

// buildInputKey takes an input id string and removes user-specified values from it
// so it can be used as a key to extract the correct id from knownResourceIds
func buildInputKey(input string) (*string, bool) {

	// don't attempt to build a key if this isn't a standard resource id
	if !strings.HasPrefix(input, "/") {
		return nil, false
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
