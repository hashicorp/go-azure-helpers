package identity_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/go-azure-helpers/framework/identity"
	"github.com/hashicorp/go-azure-helpers/framework/typehelpers"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	rmidentity "github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
)

func TestExpandSystemAssigned(t *testing.T) {
	ctx := context.Background()
	diags := diag.Diagnostics{}

	cases := []struct {
		Name     string
		Input    typehelpers.ListNestedObjectValueOf[identity.SystemIdentityModel]
		Expected *rmidentity.SystemAssigned
	}{
		{
			Name:  "null",
			Input: typehelpers.NewListNestedObjectValueOfNull[identity.SystemIdentityModel](ctx),
			Expected: &rmidentity.SystemAssigned{
				Type:        rmidentity.TypeNone,
				PrincipalId: "",
				TenantId:    "",
			},
		},
		{
			Name: "explicit none",
			Input: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []identity.SystemIdentityModel{
				{
					Type:        types.StringValue(string(identity.TypeNone)),
					PrincipalID: types.StringNull(),
					TenantID:    types.StringNull(),
				},
			}),
			Expected: &rmidentity.SystemAssigned{
				Type:        rmidentity.TypeNone,
				PrincipalId: "",
				TenantId:    "",
			},
		},
		{
			Name: "SystemAssigned",
			Input: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []identity.SystemIdentityModel{
				{
					Type:        types.StringValue(string(identity.TypeSystemAssigned)),
					PrincipalID: types.StringValue("000000-0000-0000-0000-000000000000"),
					TenantID:    types.StringValue("000000-0000-0000-0000-000000000001"),
				},
			}),
			Expected: &rmidentity.SystemAssigned{
				Type:        rmidentity.TypeSystemAssigned,
				PrincipalId: "000000-0000-0000-0000-000000000000",
				TenantId:    "000000-0000-0000-0000-000000000001",
			},
		},
		{
			Name: "SystemAssigned, UserAssigned",
			Input: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []identity.SystemIdentityModel{
				{
					Type:        types.StringValue(string(identity.TypeSystemAssignedUserAssigned)),
					PrincipalID: types.StringValue("000000-0000-0000-0000-000000000002"),
					TenantID:    types.StringValue("000000-0000-0000-0000-000000000003"),
				},
			}),
			Expected: &rmidentity.SystemAssigned{
				Type:        rmidentity.TypeSystemAssignedUserAssigned,
				PrincipalId: "000000-0000-0000-0000-000000000002",
				TenantId:    "000000-0000-0000-0000-000000000003",
			},
		},
		{
			Name: "SystemAssigned, UserAssigned (multiple)",
			Input: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []identity.SystemIdentityModel{
				{
					Type:        types.StringValue(string(identity.TypeSystemAssignedUserAssigned)),
					PrincipalID: types.StringValue("000000-0000-0000-0000-000000000002"),
					TenantID:    types.StringValue("000000-0000-0000-0000-000000000003"),
				},
			}),
			Expected: &rmidentity.SystemAssigned{
				Type:        rmidentity.TypeSystemAssignedUserAssigned,
				PrincipalId: "000000-0000-0000-0000-000000000002",
				TenantId:    "000000-0000-0000-0000-000000000003",
			},
		},
	}

	for _, tc := range cases {
		result := &rmidentity.SystemAssigned{}
		identity.ExpandToSystemAssigned(ctx, tc.Input, result, &diags)

		if !reflect.DeepEqual(result, tc.Expected) {
			t.Errorf("\nTesting: %s\nExpected: %+v\nGot: %+v\nDiags: %+v", tc.Name, tc.Expected, result, diags.Errors())
		}
	}
}

func TestFlattenSystemAssigned(t *testing.T) {
	ctx := context.Background()
	diags := diag.Diagnostics{}

	cases := []struct {
		Name     string
		Input    *rmidentity.SystemAssigned
		Expected typehelpers.ListNestedObjectValueOf[identity.SystemIdentityModel]
	}{
		{
			Name:     "null",
			Input:    nil,
			Expected: typehelpers.NewListNestedObjectValueOfNull[identity.SystemIdentityModel](ctx),
		},
		{
			Name: "explicit none",
			Input: &rmidentity.SystemAssigned{
				Type:        rmidentity.TypeNone,
				PrincipalId: "",
				TenantId:    "",
			},
			Expected: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []identity.SystemIdentityModel{
				{
					Type:        types.StringValue(string(identity.TypeNone)),
					PrincipalID: types.StringValue(""),
					TenantID:    types.StringValue(""),
				},
			}),
		},
		{
			Name: "SystemAssigned",
			Input: &rmidentity.SystemAssigned{
				Type:        rmidentity.TypeSystemAssigned,
				PrincipalId: "000000-0000-0000-0000-000000000000",
				TenantId:    "000000-0000-0000-0000-000000000001",
			},
			Expected: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []identity.SystemIdentityModel{
				{
					Type:        types.StringValue(string(identity.TypeSystemAssigned)),
					PrincipalID: types.StringValue("000000-0000-0000-0000-000000000000"),
					TenantID:    types.StringValue("000000-0000-0000-0000-000000000001"),
				},
			}),
		},
		{
			Name: "SystemAssigned, UserAssigned",
			Input: &rmidentity.SystemAssigned{
				Type:        rmidentity.TypeSystemAssignedUserAssigned,
				PrincipalId: "000000-0000-0000-0000-000000000002",
				TenantId:    "000000-0000-0000-0000-000000000003",
			},
			Expected: typehelpers.NewListNestedObjectValueOfValueSliceMust(ctx, []identity.SystemIdentityModel{
				{
					Type:        types.StringValue(string(identity.TypeSystemAssignedUserAssigned)),
					PrincipalID: types.StringValue("000000-0000-0000-0000-000000000002"),
					TenantID:    types.StringValue("000000-0000-0000-0000-000000000003"),
				},
			}),
		},
	}

	for _, tc := range cases {
		result := typehelpers.ListNestedObjectValueOf[identity.SystemIdentityModel]{}
		identity.FlattenFromSystemAssigned(ctx, tc.Input, &result, &diags)

		if !tc.Expected.Equal(result) {
			t.Errorf("\nTesting: %s\nExpected: %+v\nGot: %+v\nDiags: %+v", tc.Name, tc.Expected, result, diags.Errors())
		}
	}
}
