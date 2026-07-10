package commonids

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = &CosmosDBAccountId{}

// CosmosDBAccountId is a struct representing the Resource ID for a CosmosDB Account
type CosmosDBAccountId struct {
	SubscriptionId      string
	ResourceGroupName   string
	DatabaseAccountName string
}

// NewCosmosDBAccountID returns a new CosmosDBAccountId struct
func NewCosmosDBAccountID(subscriptionId string, resourceGroupName string, databaseAccountName string) CosmosDBAccountId {
	return CosmosDBAccountId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		DatabaseAccountName: databaseAccountName,
	}
}

// ParseCosmosDBAccountID parses 'input' into a CosmosDBAccountId
func ParseCosmosDBAccountID(input string) (*CosmosDBAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CosmosDBAccountId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CosmosDBAccountId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCosmosDBAccountIDInsensitively parses 'input' case-insensitively into a CosmosDBAccountId
// note: this method should only be used for API response data and not user input
func ParseCosmosDBAccountIDInsensitively(input string) (*CosmosDBAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CosmosDBAccountId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CosmosDBAccountId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CosmosDBAccountId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.DatabaseAccountName, ok = input.Parsed["databaseAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "databaseAccountName", input)
	}

	return nil
}

// ValidateCosmosDBAccountID checks that 'input' can be parsed as a Database Account ID
func ValidateCosmosDBAccountID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCosmosDBAccountID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Database Account ID
func (id CosmosDBAccountId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DatabaseAccountName)
}

// Segments returns a slice of Resource ID Segments which comprise this Database Account ID
func (id CosmosDBAccountId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDocumentDB", "Microsoft.DocumentDB", "Microsoft.DocumentDB"),
		resourceids.StaticSegment("staticDatabaseAccounts", "databaseAccounts", "databaseAccounts"),
		resourceids.UserSpecifiedSegment("databaseAccountName", "databaseAccountName"),
	}
}

// String returns a human-readable description of this Database Account ID
func (id CosmosDBAccountId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("CosmosDB Account Name: %q", id.DatabaseAccountName),
	}
	return fmt.Sprintf("CosmosDB Account (%s)", strings.Join(components, "\n"))
}
