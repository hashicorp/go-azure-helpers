// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package commonids

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = &TenantId{}

// TenantId is a struct representing the Resource ID for a Tenant
type TenantId struct {
	TenantId string
}

// NewTenantID returns a new TenantId struct
func NewTenantID(tenantId string) TenantId {
	return TenantId{
		TenantId: tenantId,
	}
}

// ParseTenantID parses 'input' into a TenantId
func ParseTenantID(input string) (*TenantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TenantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TenantId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTenantIDInsensitively parses 'input' case-insensitively into a TenantId
// note: this method should only be used for API response data and not user input
func ParseTenantIDInsensitively(input string) (*TenantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TenantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TenantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TenantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.TenantId, ok = input.Parsed["tenantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "tenantId", input)
	}

	return nil
}

// ValidateTenantID checks that 'input' can be parsed as a Tenant ID
func ValidateTenantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTenantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Tenant ID
func (id TenantId) ID() string {
	fmtString := "/tenants/%s"
	return fmt.Sprintf(fmtString, id.TenantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Tenant ID
func (id TenantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("tenants", "tenants", "tenants"),
		resourceids.UserSpecifiedSegment("tenantId", "12345678-1234-9876-4563-123456789012"),
	}
}

// String returns a human-readable description of this Tenant ID
func (id TenantId) String() string {
	components := []string{
		fmt.Sprintf("Tenant: %q", id.TenantId),
	}
	return fmt.Sprintf("Tenant (%s)", strings.Join(components, "\n"))
}
