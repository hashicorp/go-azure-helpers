package resourceids

import (
	"reflect"
	"testing"
)

func TestParseEmptyId(t *testing.T) {
	t.Logf("Sensitively")
	actual, err := NewParser([]Segment{}).Parse("", false)
	if err == nil {
		t.Fatalf("expected an error but didn't get one")
	}
	if actual != nil {
		t.Fatalf("got a parse result but didn't expect one: %+v", *actual)
	}

	t.Logf("Insensitively")
	actual, err = NewParser([]Segment{}).Parse("", true)
	if err == nil {
		t.Fatalf("expected an error but didn't get one")
	}
	if actual != nil {
		t.Fatalf("got a parse result but didn't expect one: %+v", *actual)
	}
}

func TestParseStaticId(t *testing.T) {
	testData := []struct {
		name          string
		segments      []Segment
		expected      *ParseResult
		input         string
		insensitively bool
	}{
		{
			name: "single segment sensitive",
			segments: []Segment{
				StaticSegment("hello", "hello", "example"),
			},
			expected: &ParseResult{
				map[string]string{
					"hello": "hello",
				},
			},
			input:         "/hello",
			insensitively: false,
		},
		{
			name: "single segment insensitive",
			segments: []Segment{
				StaticSegment("hello", "hello", "example"),
			},
			expected: &ParseResult{
				map[string]string{
					"hello": "hello",
				},
			},
			input:         "/Hello",
			insensitively: true,
		},
		{
			name: "multiple segments sensitive",
			segments: []Segment{
				StaticSegment("hello", "hello", "example"),
				StaticSegment("there", "there", "example"),
			},
			expected: &ParseResult{
				map[string]string{
					"hello": "hello",
					"there": "there",
				},
			},
			input:         "/hello/there",
			insensitively: false,
		},
		{
			name: "multiple segments sensitive",
			segments: []Segment{
				StaticSegment("hello", "hello", "example"),
				StaticSegment("there", "there", "example"),
			},
			expected: &ParseResult{
				map[string]string{
					"hello": "hello",
					"there": "there",
				},
			},
			input:         "/Hello/tHere",
			insensitively: true,
		},
	}
	for _, test := range testData {
		t.Logf("Test %q..", test.name)
		parser := NewParser(test.segments)
		actual, err := parser.Parse(test.input, test.insensitively)
		validateResult(t, actual, test.expected, err)
	}
}

func TestParseResourceGroupId(t *testing.T) {
	segments := []Segment{
		StaticSegment("subscriptions", "subscriptions", "example"),
		SubscriptionIdSegment("subscriptionId", "example"),
		StaticSegment("resourceGroups", "resourceGroups", "example"),
		ResourceGroupSegment("resourceGroupName", "example"),
	}
	testData := []struct {
		name        string
		input       string
		expected    *ParseResult
		insensitive bool
	}{
		{
			name:        "empty id - sensitive",
			input:       "",
			insensitive: false,
		},
		{
			name:        "empty id - insensitive",
			input:       "",
			insensitive: true,
		},
		{
			name:        "empty slash - sensitive",
			input:       "/",
			insensitive: false,
		},
		{
			name:        "empty slash - insensitive",
			input:       "/",
			insensitive: true,
		},
		{
			name:        "subscription id - sensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777",
			insensitive: false,
		},
		{
			name:        "subscription id - insensitive",
			input:       "/subscRiptions/11112222-3333-4444-555566667777",
			insensitive: true,
		},
		{
			name:        "resource groups list - sensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourceGroups",
			insensitive: false,
		},
		{
			name:        "resource groups list - insensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourcegroups",
			insensitive: true,
		},
		{
			name:        "resource groups id - empty name - sensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourceGroups/",
			insensitive: false,
		},
		{
			name:        "resource groups id - empty name - insensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourcegroups/",
			insensitive: true,
		},
		{
			name:        "resource group id - sensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourceGroups/BoB",
			insensitive: false,
			expected: &ParseResult{
				map[string]string{
					"subscriptions":     "subscriptions",
					"subscriptionId":    "11112222-3333-4444-555566667777",
					"resourceGroups":    "resourceGroups",
					"resourceGroupName": "BoB",
				},
			},
		},
		{
			name:        "resource groups id - insensitive",
			input:       "/subscRiptions/11112222-3333-4444-555566667777/resourcegroups/BoB",
			insensitive: true,
			expected: &ParseResult{
				map[string]string{
					"subscriptions":     "subscriptions",
					"subscriptionId":    "11112222-3333-4444-555566667777",
					"resourceGroups":    "resourceGroups",
					"resourceGroupName": "BoB",
				},
			},
		},
	}
	for _, test := range testData {
		t.Logf("Test %q..", test.name)
		parser := NewParser(segments)
		actual, err := parser.Parse(test.input, test.insensitive)
		validateResult(t, actual, test.expected, err)
	}
}

func TestParseVirtualMachineId(t *testing.T) {
	segments := []Segment{
		StaticSegment("subscriptions", "subscriptions", "example"),
		SubscriptionIdSegment("subscriptionId", "example"),
		StaticSegment("resourceGroups", "resourceGroups", "example"),
		ResourceGroupSegment("resourceGroupName", "example"),
		StaticSegment("providers", "providers", "example"),
		ResourceProviderSegment("provider", "Microsoft.Compute", "example"),
		StaticSegment("virtualMachines", "virtualMachines", "example"),
		UserSpecifiedSegment("virtualMachineName", "example"),
	}
	testData := []struct {
		name        string
		input       string
		expected    *ParseResult
		insensitive bool
	}{
		{
			name:        "resource group id - sensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourceGroups/BoB",
			insensitive: false,
		},
		{
			name:        "resource groups id - insensitive",
			input:       "/subscRiptions/11112222-3333-4444-555566667777/resourcegroups/BoB",
			insensitive: true,
		},
		{
			name:        "resource providers list - sensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourceGroups/BoB/providers",
			insensitive: false,
		},
		{
			name:        "resource providers list - insensitive",
			input:       "/subscRiptions/11112222-3333-4444-555566667777/resourcegroups/BoB/proViders",
			insensitive: true,
		},
		{
			name:        "resource provider name - sensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourceGroups/BoB/providers/Microsoft.Compute",
			insensitive: false,
		},
		{
			name:        "resource provider name - insensitive",
			input:       "/subscRiptions/11112222-3333-4444-555566667777/resourcegroups/BoB/proViders/Microsoft.compute",
			insensitive: true,
		},
		{
			name:        "virtual machine list - sensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourceGroups/BoB/providers/Microsoft.Compute/virtualMachines",
			insensitive: false,
		},
		{
			name:        "virtual machine list - insensitive",
			input:       "/subscRiptions/11112222-3333-4444-555566667777/resourcegroups/BoB/proViders/Microsoft.compute/virtualmachines",
			insensitive: true,
		},
		{
			name:        "virtual machine id - empty name - sensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourceGroups/BoB/providers/Microsoft.Compute/virtualMachines/",
			insensitive: false,
		},
		{
			name:        "virtual machine id - empty name - insensitive",
			input:       "/subscRiptions/11112222-3333-4444-555566667777/resourcegroups/BoB/proViders/Microsoft.compute/virtualmachines/",
			insensitive: true,
		},
		{
			name:        "virtual machine id - sensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourceGroups/BoB/providers/Microsoft.Compute/virtualMachines/machine1",
			insensitive: false,
			expected: &ParseResult{
				map[string]string{
					"subscriptions":      "subscriptions",
					"subscriptionId":     "11112222-3333-4444-555566667777",
					"resourceGroups":     "resourceGroups",
					"resourceGroupName":  "BoB",
					"providers":          "providers",
					"provider":           "Microsoft.Compute",
					"virtualMachines":    "virtualMachines",
					"virtualMachineName": "machine1",
				},
			},
		},
		{
			name:        "virtual machine id - insensitive",
			input:       "/subScriptions/11112222-3333-4444-555566667777/resourcegroups/BoB/pRoviders/microsoft.Compute/virtualmachines/machine1",
			insensitive: true,
			expected: &ParseResult{
				map[string]string{
					"subscriptions":      "subscriptions",
					"subscriptionId":     "11112222-3333-4444-555566667777",
					"resourceGroups":     "resourceGroups",
					"resourceGroupName":  "BoB",
					"providers":          "providers",
					"provider":           "Microsoft.Compute",
					"virtualMachines":    "virtualMachines",
					"virtualMachineName": "machine1",
				},
			},
		},
		{
			name:        "virtual machine extension id - sensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourceGroups/BoB/providers/Microsoft.Compute/virtualMachines/machine1/extensions/extension1",
			insensitive: false,
		},
		{
			name:        "virtual machine extension id - insensitive",
			input:       "/subScriptions/11112222-3333-4444-555566667777/resourcegroups/BoB/pRoviders/microsoft.Compute/virtualmachines/machine1/extenSions/extension1",
			insensitive: true,
		},
	}
	for _, test := range testData {
		t.Logf("Test %q..", test.name)
		parser := NewParser(segments)
		actual, err := parser.Parse(test.input, test.insensitive)
		validateResult(t, actual, test.expected, err)
	}
}

func TestParseVirtualMachineExtensionId(t *testing.T) {
	segments := []Segment{
		StaticSegment("subscriptions", "subscriptions", "example"),
		SubscriptionIdSegment("subscriptionId", "example"),
		StaticSegment("resourceGroups", "resourceGroups", "example"),
		ResourceGroupSegment("resourceGroupName", "example"),
		StaticSegment("providers", "providers", "example"),
		ResourceProviderSegment("provider", "Microsoft.Compute", "example"),
		StaticSegment("virtualMachines", "virtualMachines", "example"),
		UserSpecifiedSegment("virtualMachineName", "example"),
		StaticSegment("extensions", "extensions", "example"),
		UserSpecifiedSegment("extensionName", "example"),
	}
	testData := []struct {
		name        string
		input       string
		expected    *ParseResult
		insensitive bool
	}{
		{
			name:        "virtual machine id - sensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourceGroups/BoB/providers/Microsoft.Compute/virtualMachines/machine1",
			insensitive: false,
		},
		{
			name:        "virtual machine id - insensitive",
			input:       "/subScriptions/11112222-3333-4444-555566667777/resourcegroups/BoB/pRoviders/microsoft.Compute/virtualmachines/machine1",
			insensitive: true,
		},
		{
			name:        "virtual machine extensions list - sensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourceGroups/BoB/providers/Microsoft.Compute/virtualMachines/machine1/extensions",
			insensitive: false,
		},
		{
			name:        "virtual machine extensions list - insensitive",
			input:       "/subScriptions/11112222-3333-4444-555566667777/resourcegroups/BoB/pRoviders/microsoft.Compute/virtualmachines/machine1/extensions",
			insensitive: true,
		},
		{
			name:        "virtual machine extensions id - empty name - sensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourceGroups/BoB/providers/Microsoft.Compute/virtualMachines/machine1/extensions/",
			insensitive: false,
		},
		{
			name:        "virtual machine extensions id - empty name - insensitive",
			input:       "/subScriptions/11112222-3333-4444-555566667777/resourcegroups/BoB/pRoviders/microsoft.Compute/virtualmachines/machine1/extensions/",
			insensitive: true,
		},
		{
			name:        "virtual machine extension id - sensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourceGroups/BoB/providers/Microsoft.Compute/virtualMachines/machine1/extensions/extension1",
			insensitive: false,
			expected: &ParseResult{
				map[string]string{
					"subscriptions":      "subscriptions",
					"subscriptionId":     "11112222-3333-4444-555566667777",
					"resourceGroups":     "resourceGroups",
					"resourceGroupName":  "BoB",
					"providers":          "providers",
					"provider":           "Microsoft.Compute",
					"virtualMachines":    "virtualMachines",
					"virtualMachineName": "machine1",
					"extensions":         "extensions",
					"extensionName":      "extension1",
				},
			},
		},
		{
			name:        "resource groups id - insensitive",
			input:       "/subScriptions/11112222-3333-4444-555566667777/resourcegroups/BoB/pRoviders/microsoft.Compute/virtualmachines/machine1/exTensions/extension1",
			insensitive: true,
			expected: &ParseResult{
				map[string]string{
					"subscriptions":      "subscriptions",
					"subscriptionId":     "11112222-3333-4444-555566667777",
					"resourceGroups":     "resourceGroups",
					"resourceGroupName":  "BoB",
					"providers":          "providers",
					"provider":           "Microsoft.Compute",
					"virtualMachines":    "virtualMachines",
					"virtualMachineName": "machine1",
					"extensions":         "extensions",
					"extensionName":      "extension1",
				},
			},
		},
	}
	for _, test := range testData {
		t.Logf("Test %q..", test.name)
		parser := NewParser(segments)
		actual, err := parser.Parse(test.input, test.insensitive)
		validateResult(t, actual, test.expected, err)
	}
}

func TestParseAdvancedThreatProtectionId(t *testing.T) {
	segments := []Segment{
		ScopeSegment("scope", "example"),
		StaticSegment("providers", "providers", "example"),
		ResourceProviderSegment("provider", "Microsoft.Security", "example"),
		StaticSegment("advancedThreatProtectionSettings", "advancedThreatProtectionSettings", "example"),
		UserSpecifiedSegment("name", "example"),
	}
	testData := []struct {
		name        string
		input       string
		expected    *ParseResult
		insensitive bool
	}{
		{
			name:        "resource group id - sensitive",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourceGroups/BoB/providers/Microsoft.Security/advancedThreatProtectionSettings/someName",
			insensitive: false,
			expected: &ParseResult{
				map[string]string{
					"scope":                            "/subscriptions/11112222-3333-4444-555566667777/resourceGroups/BoB",
					"providers":                        "providers",
					"provider":                         "Microsoft.Security",
					"advancedThreatProtectionSettings": "advancedThreatProtectionSettings",
					"name":                             "someName",
				},
			},
		},
		{
			name:        "resource group id - sensitive invalid",
			input:       "/subscripTions/11112222-3333-4444-555566667777/resourcEgroups/BoB/proviDers/Microsoft.security/advancedthreatProtectionSettings/someName",
			insensitive: false,
		},
		{
			name:        "resource group id - insensitive",
			input:       "/subscripTions/11112222-3333-4444-555566667777/resourcEgroups/BoB/proviDers/Microsoft.security/advancedthreatProtectionSettings/someName",
			insensitive: true,
			expected: &ParseResult{
				map[string]string{
					"scope":                            "/subscripTions/11112222-3333-4444-555566667777/resourcEgroups/BoB",
					"providers":                        "providers",
					"provider":                         "Microsoft.Security",
					"advancedThreatProtectionSettings": "advancedThreatProtectionSettings",
					"name":                             "someName",
				},
			},
		},
		{
			name:        "resource group id - sensitive w/extra",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourceGroups/BoB/providers/Microsoft.Security/advancedThreatProtectionSettings/someName/extra",
			insensitive: false,
		},
		{
			name:        "resource group id - insensitive w/extra",
			input:       "/subscriptions/11112222-3333-4444-555566667777/resourceGroups/BoB/providers/Microsoft.Security/advancedThreatProtectionSettings/someName/extra",
			insensitive: true,
		},
		{
			name:        "virtual machine id - sensitive",
			input:       "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/virtualMachines/machine1/providers/Microsoft.Security/advancedThreatProtectionSettings/someName",
			insensitive: false,
			expected: &ParseResult{
				map[string]string{
					"scope":                            "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/virtualMachines/machine1",
					"providers":                        "providers",
					"provider":                         "Microsoft.Security",
					"advancedThreatProtectionSettings": "advancedThreatProtectionSettings",
					"name":                             "someName",
				},
			},
		},
		{
			name:        "virtual machine id - sensitive invalid",
			input:       "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/virtualMachines/machine1/Providers/Microsoft.SecuRity/advancedThreatprotectionSettings/someName",
			insensitive: false,
		},
		{
			name:        "virtual machine id - insensitive",
			input:       "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/virtualMachines/machine1/Providers/Microsoft.SecuRity/advancedThreatprotectionSettings/someName",
			insensitive: true,
			expected: &ParseResult{
				map[string]string{
					"scope":                            "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/virtualMachines/machine1",
					"providers":                        "providers",
					"provider":                         "Microsoft.Security",
					"advancedThreatProtectionSettings": "advancedThreatProtectionSettings",
					"name":                             "someName",
				},
			},
		},
	}
	for _, test := range testData {
		t.Logf("Test %q..", test.name)
		parser := NewParser(segments)
		actual, err := parser.Parse(test.input, test.insensitive)
		validateResult(t, actual, test.expected, err)
	}
}

func TestParseIdContainingAConstant(t *testing.T) {
	segments := []Segment{
		StaticSegment("planets", "planets", "example"),
		ConstantSegment("planetName", []string{"Mars", "Earth"}, "example"),
	}
	testData := []struct {
		name        string
		input       string
		expected    *ParseResult
		insensitive bool
	}{
		{
			name:        "planets - top level - sensitive",
			input:       "/planets",
			insensitive: false,
		},
		{
			name:        "planets - top level - insensitive",
			input:       "/plaNets",
			insensitive: true,
		},
		{
			name:        "planets - earth - sensitive",
			input:       "/planets/Earth",
			insensitive: false,
			expected: &ParseResult{
				map[string]string{
					"planets":    "planets",
					"planetName": "Earth",
				},
			},
		},
		{
			name:        "planets - earth - insensitive",
			input:       "/planets/earth",
			insensitive: true,
			expected: &ParseResult{
				map[string]string{
					"planets":    "planets",
					"planetName": "Earth",
				},
			},
		},
		{
			name:        "planets - mars - sensitive",
			input:       "/planets/Mars",
			insensitive: false,
			expected: &ParseResult{
				map[string]string{
					"planets":    "planets",
					"planetName": "Mars",
				},
			},
		},
		{
			name:        "planets - mars - insensitive",
			input:       "/planets/mars",
			insensitive: true,
			expected: &ParseResult{
				map[string]string{
					"planets":    "planets",
					"planetName": "Mars",
				},
			},
		},
		{
			name:        "planets - Pluto (invalid) - sensitive",
			input:       "/planets/Pluto",
			insensitive: false,
		},
		{
			name:        "planets - Pluto (invalid) - insensitive",
			input:       "/planets/pluto",
			insensitive: true,
		},
	}
	for _, test := range testData {
		t.Logf("Test %q..", test.name)
		parser := NewParser(segments)
		actual, err := parser.Parse(test.input, test.insensitive)
		validateResult(t, actual, test.expected, err)
	}
}

func TestParseIdContainingAScopePrefix(t *testing.T) {
	segments := []Segment{
		ScopeSegment("scope", "example"),
		StaticSegment("extensions", "extensions", "example"),
		UserSpecifiedSegment("extensionName", "example"),
	}
	testData := []struct {
		name        string
		input       string
		expected    *ParseResult
		insensitive bool
	}{
		{
			name:        "missing scope - sensitive",
			input:       "/extensions/bob",
			insensitive: false,
		},
		{
			name:        "missing scope - insensitive",
			input:       "/extenSions/bob",
			insensitive: true,
		},
		{
			name:        "scope - single level - sensitive",
			input:       "/planets/extensions/terraform",
			insensitive: false,
			expected: &ParseResult{
				map[string]string{
					"scope":         "/planets",
					"extensions":    "extensions",
					"extensionName": "terraform",
				},
			},
		},
		{
			name:        "scope - single level - insensitive",
			input:       "/planets/extenSions/terraform",
			insensitive: true,
			expected: &ParseResult{
				map[string]string{
					"scope":         "/planets",
					"extensions":    "extensions",
					"extensionName": "terraform",
				},
			},
		},
		{
			name:        "scope - multiple level - sensitive",
			input:       "/solarSystems/milkyWay/planets/mars/extensions/terraform",
			insensitive: false,
			expected: &ParseResult{
				map[string]string{
					"scope":         "/solarSystems/milkyWay/planets/mars",
					"extensions":    "extensions",
					"extensionName": "terraform",
				},
			},
		},
		{
			name:        "scope - multiple level - insensitive",
			input:       "/solarSystems/milkyWay/planets/mars/extenSions/terraform",
			insensitive: true,
			expected: &ParseResult{
				map[string]string{
					"scope":         "/solarSystems/milkyWay/planets/mars",
					"extensions":    "extensions",
					"extensionName": "terraform",
				},
			},
		},
	}
	for _, test := range testData {
		t.Logf("Test %q..", test.name)
		parser := NewParser(segments)
		actual, err := parser.Parse(test.input, test.insensitive)
		validateResult(t, actual, test.expected, err)
	}
}

func TestParseIdContainingAScopeSuffix(t *testing.T) {
	segments := []Segment{
		StaticSegment("subscriptions", "subscriptions", "example"),
		SubscriptionIdSegment("subscriptionId", "example"),
		ScopeSegment("scope", "example"),
	}
	testData := []struct {
		name        string
		input       string
		expected    *ParseResult
		insensitive bool
	}{
		{
			name:        "missing scope - sensitive",
			input:       "/subscriptions/1111",
			insensitive: false,
		},
		{
			name:        "missing scope - insensitive",
			input:       "/subscriPtions/1111",
			insensitive: true,
		},
		{
			name:        "scope - single level - sensitive",
			input:       "/subscriptions/1111/someThing",
			insensitive: false,
			expected: &ParseResult{
				map[string]string{
					"scope":          "/someThing",
					"subscriptions":  "subscriptions",
					"subscriptionId": "1111",
				},
			},
		},
		{
			name:        "scope - single level - insensitive",
			input:       "/subscrIptions/1111/someThing",
			insensitive: true,
			expected: &ParseResult{
				map[string]string{
					"scope":          "/someThing",
					"subscriptions":  "subscriptions",
					"subscriptionId": "1111",
				},
			},
		},
		{
			name:        "scope - multiple level - sensitive",
			input:       "/subscriptions/1111/solarSystems/milkyWay/planets/mars",
			insensitive: false,
			expected: &ParseResult{
				map[string]string{
					"scope":          "/solarSystems/milkyWay/planets/mars",
					"subscriptions":  "subscriptions",
					"subscriptionId": "1111",
				},
			},
		},
		{
			name:        "scope - multiple level - insensitive",
			input:       "/subscriPtions/1111/solarSystems/milkyWay/planets/mars",
			insensitive: true,
			expected: &ParseResult{
				map[string]string{
					"scope":          "/solarSystems/milkyWay/planets/mars",
					"subscriptions":  "subscriptions",
					"subscriptionId": "1111",
				},
			},
		},
	}
	for _, test := range testData {
		t.Logf("Test %q..", test.name)
		parser := NewParser(segments)
		actual, err := parser.Parse(test.input, test.insensitive)
		validateResult(t, actual, test.expected, err)
	}
}

func TestParseIdContainingAScopeEitherEnd(t *testing.T) {
	segments := []Segment{
		ScopeSegment("start", "example"),
		StaticSegment("connections", "connections", "example"),
		SubscriptionIdSegment("connectionName", "example"),
		ScopeSegment("end", "example"),
	}
	testData := []struct {
		name        string
		input       string
		expected    *ParseResult
		insensitive bool
	}{
		{
			name:        "missing start - sensitive",
			input:       "/connections/BER-FCO/someOtherThing",
			insensitive: false,
		},
		{
			name:        "missing start - insensitive",
			input:       "/connecTions/BER-FCO/someOtherThing",
			insensitive: true,
		},
		{
			name:        "missing end - sensitive",
			input:       "/someThing/connections/BER-FCO",
			insensitive: false,
		},
		{
			name:        "missing end - insensitive",
			input:       "/someThing/connEctions/BER-FCO",
			insensitive: true,
		},
		{
			name:        "missing both ends - sensitive",
			input:       "/connections/BER-FCO",
			insensitive: false,
		},
		{
			name:        "missing both ends - insensitive",
			input:       "/connectiOns/BER-FCO",
			insensitive: true,
		},
		{
			name:        "scope - single level - sensitive",
			input:       "/someThing/connections/BER-FCO/someOtherThing",
			insensitive: false,
			expected: &ParseResult{
				map[string]string{
					"connectionName": "BER-FCO",
					"connections":    "connections",
					"end":            "/someOtherThing",
					"start":          "/someThing",
				},
			},
		},
		{
			name:        "scope - single level - insensitive",
			input:       "/someThing/Connections/BER-FCO/someOtherThing",
			insensitive: true,
			expected: &ParseResult{
				map[string]string{
					"connectionName": "BER-FCO",
					"connections":    "connections",
					"end":            "/someOtherThing",
					"start":          "/someThing",
				},
			},
		},
		{
			name:        "scope - multiple level - sensitive",
			input:       "/someThing/thats/really/awesome/connections/BER-FCO/someOtherThing/woah",
			insensitive: false,
			expected: &ParseResult{
				map[string]string{
					"connectionName": "BER-FCO",
					"connections":    "connections",
					"end":            "/someOtherThing/woah",
					"start":          "/someThing/thats/really/awesome",
				},
			},
		},
		{
			name:        "scope - multiple level - insensitive",
			input:       "/someThing/thats/really/awesome/conNections/BER-FCO/someOtherThing/woah",
			insensitive: true,
			expected: &ParseResult{
				map[string]string{
					"connectionName": "BER-FCO",
					"connections":    "connections",
					"end":            "/someOtherThing/woah",
					"start":          "/someThing/thats/really/awesome",
				},
			},
		},
	}
	for _, test := range testData {
		t.Logf("Test %q..", test.name)
		parser := NewParser(segments)
		actual, err := parser.Parse(test.input, test.insensitive)
		validateResult(t, actual, test.expected, err)
	}
}

func TestParseIdContainingJustAScope(t *testing.T) {
	segments := []Segment{
		ScopeSegment("scope", "example"),
	}
	testData := []struct {
		name        string
		input       string
		expected    *ParseResult
		insensitive bool
	}{
		{
			name:        "empty",
			input:       "",
			insensitive: false,
		},
		{
			name:        "slash - sensitive",
			input:       "/",
			insensitive: false,
			expected: &ParseResult{
				map[string]string{
					"scope": "/",
				},
			},
		},
		{
			name:        "slash - insensitive",
			input:       "/",
			insensitive: true,
			expected: &ParseResult{
				map[string]string{
					"scope": "/",
				},
			},
		},
		{
			name:        "single level - sensitive",
			input:       "/hello",
			insensitive: false,
			expected: &ParseResult{
				map[string]string{
					"scope": "/hello",
				},
			},
		},
		{
			name:        "single level - insensitive",
			input:       "/hello",
			insensitive: true,
			expected: &ParseResult{
				map[string]string{
					"scope": "/hello",
				},
			},
		},
		{
			name:        "multiple levels - sensitive",
			input:       "/hello/there/world",
			insensitive: false,
			expected: &ParseResult{
				map[string]string{
					"scope": "/hello/there/world",
				},
			},
		},
		{
			name:        "multiple levels - insensitive",
			input:       "/hello/there/world",
			insensitive: true,
			expected: &ParseResult{
				map[string]string{
					"scope": "/hello/there/world",
				},
			},
		},
	}
	for _, test := range testData {
		t.Logf("Test %q..", test.name)
		parser := NewParser(segments)
		actual, err := parser.Parse(test.input, test.insensitive)
		validateResult(t, actual, test.expected, err)
	}
}

func validateResult(t *testing.T, actual *ParseResult, expected *ParseResult, err error) {
	if err != nil {
		if expected == nil {
			return
		}

		t.Fatalf("got an error but didn't expect one: %+v", err)
	}
	if expected == nil {
		t.Fatalf("expected an error but didn't get one")
	}

	if expected == nil && actual == nil {
		return
	}

	if actual == nil {
		t.Fatalf("expected a parse result but didn't get one")
	}
	if expected == nil {
		t.Fatalf("expected no parse result but got %+v", actual.Parsed)
	}

	if !reflect.DeepEqual(expected.Parsed, actual.Parsed) {
		t.Fatalf("Diff between Expected and Actual.\n\nExpected: %+v\n\nActual: %+v", expected.Parsed, actual.Parsed)
	}
}
