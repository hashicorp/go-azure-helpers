package recaser

import (
	"regexp"
	"strings"
	"sync"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var knownResourceIds = make(map[string]resourceids.ResourceId)

var resourceIdsWriteLock = &sync.Mutex{}

func RegisterResourceId(id resourceids.ResourceId) {
	//fullName := fmt.Sprintf("%s/%s", reflect.TypeOf(id).PkgPath(), reflect.TypeOf(id).Name())
	key := id.ID()

	resourceIdsWriteLock.Lock()
	if _, ok := knownResourceIds[key]; !ok {
		knownResourceIds[key] = id
	}
	resourceIdsWriteLock.Unlock()
}

// ReCase tries to determine the type of Resource ID defined in `input` to be able to re-case it from
func ReCase(input string) string {
	output := input
	recased := false

	for _, id := range knownResourceIds {
		output, recased = parseId(id, input)
		if recased {
			break
		}
	}

	// recase just "subscriptions" and "resourceGroups" segments if we can't find a matching id
	if !recased {
		subscriptions := "/subscriptions/"
		resourceGroups := "/resourceGroups/"

		if strings.Contains(strings.ToLower(input), strings.ToLower(subscriptions)) {
			re := regexp.MustCompile(`(?i)/subscriptions/`)
			output = re.ReplaceAllString(output, subscriptions)
			recased = true
		}

		if strings.Contains(strings.ToLower(input), strings.ToLower(resourceGroups)) {
			re := regexp.MustCompile(`(?i)/resourceGroups/`)
			output = re.ReplaceAllString(output, resourceGroups)
			recased = true
		}
	}

	// TODO if !recased check scope? possible import cycle error?

	return output
}

func parseId(id resourceids.ResourceId, input string) (string, bool) {

	// we need to take a local copy of id to work against else we're mutating the original
	localId := id

	parser := resourceids.NewParserFromResourceIdType(localId)
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return input, false
	}

	if err = id.FromParseResult(*parsed); err != nil {
		return input, false
	}
	input = id.ID()

	return input, true
}
