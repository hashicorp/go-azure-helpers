package recaser

import (
	"strings"
	"sync"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var knownResourceIds = make(map[string]resourceids.ResourceId)

var resourceIdsWriteLock = &sync.Mutex{}

func init() {
	//register common ids
	resourceIdsWriteLock.Lock()
	for _, id := range commonids.CommonIds() {
		key := strings.ToLower(id.ID())
		if _, ok := knownResourceIds[key]; !ok {
			knownResourceIds[key] = id
		}
	}
	resourceIdsWriteLock.Unlock()
}

// RegisterResourceId adds ResourceIds to a list of known ids
func RegisterResourceId(id resourceids.ResourceId) {
	key := strings.ToLower(id.ID())

	resourceIdsWriteLock.Lock()
	if _, ok := knownResourceIds[key]; !ok {
		knownResourceIds[key] = id
	}
	resourceIdsWriteLock.Unlock()
}
