package identity

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ json.Marshaler = &SystemAndUserAssignedMap{}

type SystemAndUserAssignedMap struct {
	*identity.SystemAndUserAssignedMap
}

func (s *SystemAndUserAssignedMap) MarshalJSON() ([]byte, error) {
	identityType := TypeNone
	userAssignedIdentityIds := map[string]identity.UserAssignedIdentityDetails{}

	switch s.Type {
	case identity.TypeSystemAssigned:
		identityType = TypeSystemAssigned
	case identity.TypeUserAssigned:
		identityType = TypeUserAssigned
	case identity.TypeSystemAssignedUserAssigned:
		identityType = TypeSystemAssignedUserAssigned
	}

	if identityType != TypeNone {
		userAssignedIdentityIds = s.IdentityIds
	}

	out := map[string]interface{}{
		"type":                   string(identityType),
		"userAssignedIdentities": nil,
	}
	if len(userAssignedIdentityIds) > 0 {
		out["userAssignedIdentities"] = userAssignedIdentityIds
	}

	return json.Marshal(out)
}

func ExpandSystemAndUserAssignedMapFromModel(input types.List) (result *identity.SystemAndUserAssignedMap, diags diag.Diagnostics) {
	bg := context.Background()
	if input.IsNull() || input.IsUnknown() {
		return &identity.SystemAndUserAssignedMap{
			Type:        identity.TypeNone,
			IdentityIds: nil,
		}, nil
	}

	identities := make([]ModelSystemAssignedUserAssigned, 0)

	diags = input.ElementsAs(bg, &identities, false)
	if diags.HasError() {
		return nil, diags
	}

	id := identities[0]
	ids := make([]string, 0)
	diags = id.IdentityIds.ElementsAs(bg, &ids, false)

	identityIds := make(map[string]identity.UserAssignedIdentityDetails, len(ids))
	for _, v := range ids {
		identityIds[v] = identity.UserAssignedIdentityDetails{
			// intentionally empty since the expand shouldn't send these values
		}
	}
	if len(identityIds) > 0 && (id.Type.ValueString() != string(TypeSystemAssignedUserAssigned) && id.Type.ValueString() != string(TypeUserAssigned)) {
		diags.AddError("identity error", fmt.Sprintf("`identity_ids` can only be specified when `type` is set to %q or %q", TypeSystemAssignedUserAssigned, TypeUserAssigned))
		return nil, diags
	}

	return &identity.SystemAndUserAssignedMap{
		Type:        identity.Type(id.Type.ValueString()),
		IdentityIds: identityIds,
	}, nil
}

func FlattenSystemAndUserAssignedMapToModel(input *identity.SystemAndUserAssignedMap) (result types.List, diags diag.Diagnostics) {
	bg := context.Background()
	diags = make([]diag.Diagnostic, 0)

	if input == nil {
		result = types.ListNull(types.ObjectType{}.WithAttributeTypes(IdentityAttr))
		return
	}

	i := *input
	identityObject := ModelSystemAssignedUserAssigned{
		Type:        types.StringValue(string(i.Type)),
		PrincipalId: types.StringValue(i.PrincipalId),
		TenantId:    types.StringValue(i.TenantId),
	}

	ids := make([]string, 0)
	for k := range i.IdentityIds {
		ids = append(ids, k)
	}

	identityIdsList, d := types.ListValueFrom(bg, types.StringType, ids)
	if d.HasError() {
		diags = append(diags, d...)
		return
	}

	identityObject.IdentityIds = identityIdsList

	identityObjectValue, d := types.ObjectValueFrom(bg, ModelSystemAssignedUserAssignedAttr, identityObject)

	result, diags = types.ListValueFrom(bg, types.ObjectType{}.WithAttributeTypes(ModelSystemAssignedUserAssignedAttr), []types.Object{identityObjectValue})

	return
}
