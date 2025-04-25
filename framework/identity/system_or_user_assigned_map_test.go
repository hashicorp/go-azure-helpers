package identity_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/go-azure-helpers/framework/identity"
	"github.com/hashicorp/go-azure-helpers/framework/typehelpers"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	rmidentity "github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
)

func TestExpandSystemOrUserAssignedMap(t *testing.T) {
	ctx := context.Background()
	diags := diag.Diagnostics{}

	cases := []struct {
		Name     string
		Input    typehelpers.ListNestedObjectValueOf[identity.IdentityModel]
		Expected *rmidentity.SystemOrUserAssignedMap
	}{
		{
			Name:  "null",
			Input: typehelpers.NewListNestedObjectValueOfNull[identity.IdentityModel](ctx),
			Expected: &rmidentity.SystemOrUserAssignedMap{
				Type:        rmidentity.TypeNone,
				PrincipalId: "",
				TenantId:    "",
				IdentityIds: nil,
			},
		},
		{
			Name: "explicit none",
			Input: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []identity.IdentityModel{
				{
					Type:        types.StringValue(string(rmidentity.TypeNone)),
					PrincipalID: types.StringNull(),
					TenantID:    types.StringNull(),
					IdentityIDs: typehelpers.NewSetValueOfNull[types.String](ctx),
				},
			}),
			Expected: &rmidentity.SystemOrUserAssignedMap{
				Type:        rmidentity.TypeNone,
				PrincipalId: "",
				TenantId:    "",
				IdentityIds: map[string]rmidentity.UserAssignedIdentityDetails{},
			},
		},
		{
			Name: "SystemAssigned",
			Input: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []identity.IdentityModel{
				{
					Type:        types.StringValue(string(rmidentity.TypeSystemAssigned)),
					PrincipalID: types.StringValue("000000-0000-0000-0000-000000000000"),
					TenantID:    types.StringValue("000000-0000-0000-0000-000000000001"),
					IdentityIDs: typehelpers.NewSetValueOfNull[types.String](ctx),
				},
			}),
			Expected: &rmidentity.SystemOrUserAssignedMap{
				Type:        rmidentity.TypeSystemAssigned,
				PrincipalId: "000000-0000-0000-0000-000000000000",
				TenantId:    "000000-0000-0000-0000-000000000001",
				IdentityIds: map[string]rmidentity.UserAssignedIdentityDetails{},
			},
		},
		{
			Name: "SystemAssigned, UserAssigned",
			Input: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []identity.IdentityModel{
				{
					Type:        types.StringValue(string(rmidentity.TypeSystemAssignedUserAssigned)),
					PrincipalID: types.StringValue("000000-0000-0000-0000-000000000002"),
					TenantID:    types.StringValue("000000-0000-0000-0000-000000000003"),
					IdentityIDs: typehelpers.NewSetValueOfMust[types.String](ctx, []attr.Value{
						types.StringValue("100000-0000-0000-0000-000000000000"),
					}),
				},
			}),
			Expected: &rmidentity.SystemOrUserAssignedMap{
				Type:        rmidentity.TypeSystemAssignedUserAssigned,
				PrincipalId: "000000-0000-0000-0000-000000000002",
				TenantId:    "000000-0000-0000-0000-000000000003",
				IdentityIds: map[string]rmidentity.UserAssignedIdentityDetails{
					"100000-0000-0000-0000-000000000000": {},
				},
			},
		},
		{
			Name: "SystemAssigned, UserAssigned (multiple)",
			Input: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []identity.IdentityModel{
				{
					Type:        types.StringValue(string(rmidentity.TypeSystemAssignedUserAssigned)),
					PrincipalID: types.StringValue("000000-0000-0000-0000-000000000002"),
					TenantID:    types.StringValue("000000-0000-0000-0000-000000000003"),
					IdentityIDs: typehelpers.NewSetValueOfMust[types.String](ctx, []attr.Value{
						types.StringValue("100000-0000-0000-0000-000000000000"),
						types.StringValue("800000-0000-0000-0000-000000000000"),
						types.StringValue("900000-0000-0000-0000-000000000000"),
					}),
				},
			}),
			Expected: &rmidentity.SystemOrUserAssignedMap{
				Type:        rmidentity.TypeSystemAssignedUserAssigned,
				PrincipalId: "000000-0000-0000-0000-000000000002",
				TenantId:    "000000-0000-0000-0000-000000000003",
				IdentityIds: map[string]rmidentity.UserAssignedIdentityDetails{
					"100000-0000-0000-0000-000000000000": {},
					"800000-0000-0000-0000-000000000000": {},
					"900000-0000-0000-0000-000000000000": {},
				},
			},
		},
	}

	for _, tc := range cases {
		result := &rmidentity.SystemOrUserAssignedMap{}
		identity.ExpandToSystemOrUserAssignedMap(ctx, tc.Input, result, &diags)

		if !reflect.DeepEqual(result, tc.Expected) {
			t.Fatalf("\nTesting: %s\nExpected: %+v\nGot: %+v\nDiags: %+v", tc.Name, tc.Expected, result, diags.Errors())
		}
	}
}

func TestFlattenSystemOrUserAssignedMap(t *testing.T) {
	ctx := context.Background()
	diags := diag.Diagnostics{}

	cases := []struct {
		Name     string
		Input    *rmidentity.SystemOrUserAssignedMap
		Expected typehelpers.ListNestedObjectValueOf[identity.IdentityModel]
	}{
		{
			Name:     "null",
			Input:    nil,
			Expected: typehelpers.NewListNestedObjectValueOfNull[identity.IdentityModel](ctx),
		},
		{
			Name: "explicit none",
			Input: &rmidentity.SystemOrUserAssignedMap{
				Type:        rmidentity.TypeNone,
				PrincipalId: "",
				TenantId:    "",
				IdentityIds: map[string]rmidentity.UserAssignedIdentityDetails{},
			},
			Expected: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []identity.IdentityModel{
				{
					Type:        types.StringValue(string(rmidentity.TypeNone)),
					PrincipalID: types.StringValue(""),
					TenantID:    types.StringValue(""),
					IdentityIDs: typehelpers.NewSetValueOfNull[types.String](ctx),
				},
			}),
		},
		{
			Name: "SystemAssigned",
			Input: &rmidentity.SystemOrUserAssignedMap{
				Type:        rmidentity.TypeSystemAssigned,
				PrincipalId: "000000-0000-0000-0000-000000000000",
				TenantId:    "000000-0000-0000-0000-000000000001",
				IdentityIds: map[string]rmidentity.UserAssignedIdentityDetails{},
			},
			Expected: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []identity.IdentityModel{
				{
					Type:        types.StringValue(string(rmidentity.TypeSystemAssigned)),
					PrincipalID: types.StringValue("000000-0000-0000-0000-000000000000"),
					TenantID:    types.StringValue("000000-0000-0000-0000-000000000001"),
					IdentityIDs: typehelpers.NewSetValueOfNull[types.String](ctx),
				},
			}),
		},
		{
			Name: "SystemAssigned, UserAssigned",
			Input: &rmidentity.SystemOrUserAssignedMap{
				Type:        rmidentity.TypeSystemAssignedUserAssigned,
				PrincipalId: "000000-0000-0000-0000-000000000002",
				TenantId:    "000000-0000-0000-0000-000000000003",
				IdentityIds: map[string]rmidentity.UserAssignedIdentityDetails{
					"100000-0000-0000-0000-000000000000": {},
				},
			},
			Expected: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []identity.IdentityModel{
				{
					Type:        types.StringValue(string(rmidentity.TypeSystemAssignedUserAssigned)),
					PrincipalID: types.StringValue("000000-0000-0000-0000-000000000002"),
					TenantID:    types.StringValue("000000-0000-0000-0000-000000000003"),
					IdentityIDs: typehelpers.NewSetValueOfMust[types.String](ctx, []attr.Value{
						types.StringValue("100000-0000-0000-0000-000000000000"),
					}),
				},
			}),
		},
		{
			Name: "SystemAssigned, UserAssigned (multiple)",
			Input: &rmidentity.SystemOrUserAssignedMap{
				Type:        rmidentity.TypeSystemAssignedUserAssigned,
				PrincipalId: "000000-0000-0000-0000-000000000002",
				TenantId:    "000000-0000-0000-0000-000000000003",
				IdentityIds: map[string]rmidentity.UserAssignedIdentityDetails{
					"100000-0000-0000-0000-000000000000": {},
					"100000-0000-0000-0000-000000000002": {},
					"100000-0000-0000-0000-000000000003": {},
				},
			},
			Expected: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []identity.IdentityModel{
				{
					Type:        types.StringValue(string(rmidentity.TypeSystemAssignedUserAssigned)),
					PrincipalID: types.StringValue("000000-0000-0000-0000-000000000002"),
					TenantID:    types.StringValue("000000-0000-0000-0000-000000000003"),
					IdentityIDs: typehelpers.NewSetValueOfMust[types.String](ctx, []attr.Value{
						types.StringValue("100000-0000-0000-0000-000000000000"),
						types.StringValue("100000-0000-0000-0000-000000000002"),
						types.StringValue("100000-0000-0000-0000-000000000003"),
					}),
				},
			}),
		},
	}

	for _, tc := range cases {
		result := typehelpers.ListNestedObjectValueOf[identity.IdentityModel]{}
		identity.FlattenFromSystemOrUserAssignedMap(ctx, tc.Input, &result, &diags)

		if !tc.Expected.Equal(result) {
			t.Errorf("\nTesting: %s\nExpected: %+v\nGot: %+v\nDiags: %+v", tc.Name, tc.Expected, result, diags.Errors())
		}
	}
}
