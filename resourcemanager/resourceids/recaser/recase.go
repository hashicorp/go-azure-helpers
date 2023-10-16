package recaser

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var knownResourceIds = make(map[string]resourceids.ResourceIdWithParser)

var resourceIdsWriteLock = &sync.Mutex{}

func RegisterResourceId(id resourceids.ResourceIdWithParser) {
	fullName := fmt.Sprintf("%s/%s", reflect.TypeOf(id).PkgPath(), reflect.TypeOf(id).Name())

	resourceIdsWriteLock.Lock()
	if _, ok := knownResourceIds[fullName]; !ok {
		knownResourceIds[fullName] = id
	}
	resourceIdsWriteLock.Unlock()
}

// ReCase tries to determine the type of Resource ID defined in `input` to be able to re-case it from
func ReCase(input string) string {
	output := input

	// TODO: to be useful this would need a Feature Flag within `go-azure-helpers` which conditionally re-cases
	// based on the flag - and would likely be useful to call when a Struct Tag is defined in the Typed SDK

	for _, id := range knownResourceIds {
		// we need to take a local copy of id to work against else we're mutating the original
		localId := id

		parser := resourceids.NewParserFromResourceIdType(localId)
		parsed, err := parser.Parse(input, true)
		if err != nil {
			continue
		}

		if err := id.FromParseResult(*parsed); err != nil {
			continue
		}
		output = id.ID()
		break
	}

	return output
}
