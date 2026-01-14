// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

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

func TestExpandSystemAndUserAssignedList(t *testing.T) {
	ctx := context.Background()
	diags := diag.Diagnostics{}

	cases := []struct {
		Name     string
		Input    typehelpers.ListNestedObjectValueOf[identity.IdentityModel]
		Expected *rmidentity.SystemAndUserAssignedList
	}{
		{
			Name:  "null",
			Input: typehelpers.NewListNestedObjectValueOfNull[identity.IdentityModel](ctx),
			Expected: &rmidentity.SystemAndUserAssignedList{
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
			Expected: &rmidentity.SystemAndUserAssignedList{
				Type:        rmidentity.TypeNone,
				PrincipalId: "",
				TenantId:    "",
				IdentityIds: nil,
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
			Expected: &rmidentity.SystemAndUserAssignedList{
				Type:        rmidentity.TypeSystemAssigned,
				PrincipalId: "000000-0000-0000-0000-000000000000",
				TenantId:    "000000-0000-0000-0000-000000000001",
				IdentityIds: nil,
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
			Expected: &rmidentity.SystemAndUserAssignedList{
				Type:        rmidentity.TypeSystemAssignedUserAssigned,
				PrincipalId: "000000-0000-0000-0000-000000000002",
				TenantId:    "000000-0000-0000-0000-000000000003",
				IdentityIds: []string{
					"100000-0000-0000-0000-000000000000",
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
						types.StringValue("200000-0000-0000-0000-000000000000"),
						types.StringValue("300000-0000-0000-0000-000000000000"),
					}),
				},
			}),
			Expected: &rmidentity.SystemAndUserAssignedList{
				Type:        rmidentity.TypeSystemAssignedUserAssigned,
				PrincipalId: "000000-0000-0000-0000-000000000002",
				TenantId:    "000000-0000-0000-0000-000000000003",
				IdentityIds: []string{
					"100000-0000-0000-0000-000000000000",
					"200000-0000-0000-0000-000000000000",
					"300000-0000-0000-0000-000000000000",
				},
			},
		},
	}

	for _, tc := range cases {
		result := &rmidentity.SystemAndUserAssignedList{}
		identity.ExpandToSystemAndUserAssignedList(ctx, tc.Input, result, &diags)

		if !reflect.DeepEqual(result, tc.Expected) {
			t.Errorf("\nTesting: %s\nExpected: %+v\nGot: %+v\nDiags: %+v", tc.Name, tc.Expected, result, diags.Errors())
		}
	}
}

func TestFlattenSystemAndUserAssignedList(t *testing.T) {
	ctx := context.Background()
	diags := diag.Diagnostics{}

	cases := []struct {
		Name     string
		Input    *rmidentity.SystemAndUserAssignedList
		Expected typehelpers.ListNestedObjectValueOf[identity.IdentityModel]
	}{
		{
			Name:     "null",
			Input:    nil,
			Expected: typehelpers.NewListNestedObjectValueOfNull[identity.IdentityModel](ctx),
		},
		{
			Name: "explicit none",
			Input: &rmidentity.SystemAndUserAssignedList{
				Type:        rmidentity.TypeNone,
				PrincipalId: "",
				TenantId:    "",
				IdentityIds: nil,
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
			Input: &rmidentity.SystemAndUserAssignedList{
				Type:        rmidentity.TypeSystemAssigned,
				PrincipalId: "000000-0000-0000-0000-000000000000",
				TenantId:    "000000-0000-0000-0000-000000000001",
				IdentityIds: nil,
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
			Input: &rmidentity.SystemAndUserAssignedList{
				Type:        rmidentity.TypeSystemAssignedUserAssigned,
				PrincipalId: "000000-0000-0000-0000-000000000002",
				TenantId:    "000000-0000-0000-0000-000000000003",
				IdentityIds: []string{
					"100000-0000-0000-0000-000000000000",
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
			Input: &rmidentity.SystemAndUserAssignedList{
				Type:        rmidentity.TypeSystemAssignedUserAssigned,
				PrincipalId: "000000-0000-0000-0000-000000000002",
				TenantId:    "000000-0000-0000-0000-000000000003",
				IdentityIds: []string{
					"100000-0000-0000-0000-000000000000",
					"200000-0000-0000-0000-000000000000",
					"300000-0000-0000-0000-000000000000",
				},
			},
			Expected: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []identity.IdentityModel{
				{
					Type:        types.StringValue(string(rmidentity.TypeSystemAssignedUserAssigned)),
					PrincipalID: types.StringValue("000000-0000-0000-0000-000000000002"),
					TenantID:    types.StringValue("000000-0000-0000-0000-000000000003"),
					IdentityIDs: typehelpers.NewSetValueOfMust[types.String](ctx, []attr.Value{
						types.StringValue("100000-0000-0000-0000-000000000000"),
						types.StringValue("200000-0000-0000-0000-000000000000"),
						types.StringValue("300000-0000-0000-0000-000000000000"),
					}),
				},
			}),
		},
	}

	for _, tc := range cases {
		result := typehelpers.ListNestedObjectValueOf[identity.IdentityModel]{}
		identity.FlattenFromSystemAndUserAssignedList(ctx, tc.Input, &result, &diags)

		if !tc.Expected.Equal(result) {
			t.Errorf("\nTesting: %s\nExpected: %+v\nGot: %+v\nDiags: %+v", tc.Name, tc.Expected, result, diags.Errors())
		}
	}
}
