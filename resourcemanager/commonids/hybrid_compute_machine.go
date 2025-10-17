// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package commonids

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = &HybridComputeMachineId{}

// HybridComputeMachineId is a struct representing the Resource ID for a Virtual Machine
type HybridComputeMachineId struct {
	SubscriptionId           string
	ResourceGroupName        string
	HybridComputeMachineName string
}

// NewHybridComputeMachineID returns a new HybridComputeMachineId struct
func NewHybridComputeMachineID(subscriptionId string, resourceGroupName string, hybridComputeMachineName string) HybridComputeMachineId {
	return HybridComputeMachineId{
		SubscriptionId:           subscriptionId,
		ResourceGroupName:        resourceGroupName,
		HybridComputeMachineName: hybridComputeMachineName,
	}
}

// ParseHybridComputeMachineID parses 'input' into a HybridComputeMachineId
func ParseHybridComputeMachineID(input string) (*HybridComputeMachineId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HybridComputeMachineId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HybridComputeMachineId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseHybridComputeMachineIDInsensitively parses 'input' case-insensitively into a HybridComputeMachineId
// note: this method should only be used for API response data and not user input
func ParseHybridComputeMachineIDInsensitively(input string) (*HybridComputeMachineId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HybridComputeMachineId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HybridComputeMachineId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *HybridComputeMachineId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.HybridComputeMachineName, ok = input.Parsed["machineName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "machineName", input)
	}

	return nil
}

// ValidateHybridComputeMachineID checks that 'input' can be parsed as a HybridCompute Machine ID
func ValidateHybridComputeMachineID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseHybridComputeMachineID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted HybridCompute Machine ID
func (id HybridComputeMachineId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.HybridCompute/machines/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.HybridComputeMachineName)
}

// Segments returns a slice of Resource ID Segments which comprise this HybridCompute Machine ID
func (id HybridComputeMachineId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticHybridComputeMachines", "machines", "machines"),
		resourceids.UserSpecifiedSegment("machineName", "machineValue"),
	}
}

// String returns a human-readable description of this HybridCompute Machine ID
func (id HybridComputeMachineId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("HybridCompute Machine Name: %q", id.HybridComputeMachineName),
	}
	return fmt.Sprintf("HybridCompute Machine (%s)", strings.Join(components, "\n"))
}
