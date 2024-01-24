package commonids

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type CompositeResourceID[T1 resourceids.ResourceId, T2 resourceids.ResourceId] struct {
	First  T1
	Second T2
}

func (id CompositeResourceID[T1, T2]) ID() string {
	fmtString := "%s|%s"
	return fmt.Sprintf(fmtString, id.First.ID(), id.Second.ID())
}

func (id CompositeResourceID[T1, T2]) String() string {
	fmtString := "%s\n%s"
	return fmt.Sprintf(fmtString, id.First.String(), id.Second.String())
}

func ParseCompositeResourceID[T1 resourceids.ResourceId, T2 resourceids.ResourceId](input string, first T1, second T2) (*CompositeResourceID[T1, T2], error) {
	return Parse(input, first, second, false)
}

func ParseCompositeResourceIDInsensitively[T1 resourceids.ResourceId, T2 resourceids.ResourceId](input string, first T1, second T2) (*CompositeResourceID[T1, T2], error) {
	return Parse(input, first, second, true)
}

func Parse[T1 resourceids.ResourceId, T2 resourceids.ResourceId](input string, first T1, second T2, insensitively bool) (*CompositeResourceID[T1, T2], error) {

	components := strings.Split(input, "|")
	if len(components) != 2 {
		return nil, fmt.Errorf("expected 2 resourceids but got %d", len(components))
	}

	output := CompositeResourceID[T1, T2]{
		First:  first,
		Second: second,
	}

	// Parse the first of the two Resource IDs from the components
	firstParser := resourceids.NewParserFromResourceIdType(output.First)
	firstParseResult, err := firstParser.Parse(components[0], insensitively)
	if err != nil {
		return nil, fmt.Errorf("parsing first id of CompositeResourceID: %v", err)
	}
	err = output.First.FromParseResult(*firstParseResult)
	if err != nil {
		return nil, fmt.Errorf("populating first id of CompositeResourceID: %v", err)
	}

	// Parse the second of the two Resource IDs from the components
	secondParser := resourceids.NewParserFromResourceIdType(output.Second)
	secondParseResult, err := secondParser.Parse(components[1], insensitively)
	if err != nil {
		return nil, fmt.Errorf("parsing second id of CompositeResourceID: %v", err)
	}
	err = output.Second.FromParseResult(*secondParseResult)
	if err != nil {
		return nil, fmt.Errorf("populating second id of CompositeResourceID: %v", err)
	}

	return &output, nil
}
