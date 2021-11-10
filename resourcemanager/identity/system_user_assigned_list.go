package identity

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

var _ json.Marshaler = &SystemUserAssignedList{}

type SystemUserAssignedList struct {
	Type        Type     `tfschema:"type" json:"type"`
	PrincipalId string   `tfschema:"principal_id" json:"principalId"`
	TenantId    string   `tfschema:"tenant_id" json:"tenantId"`
	IdentityIds []string `tfschema:"identity_ids" json:"userAssignedIdentities"`
}

func (s *SystemUserAssignedList) MarshalJSON() ([]byte, error) {
	// we use a custom marshal function here since we can only send the Type / UserAssignedIdentities field
	identityType := TypeNone
	userAssignedIdentityIds := []string{}

	if s != nil {
		if s.Type == TypeSystemAssigned {
			identityType = TypeSystemAssigned
		}
		if s.Type == TypeSystemAssignedUserAssigned {
			identityType = TypeSystemAssignedUserAssigned
		}
		if s.Type == TypeUserAssigned {
			identityType = TypeUserAssigned
		}

		if identityType != TypeNone {
			userAssignedIdentityIds = s.IdentityIds
		}
	}

	out := map[string]interface{}{
		"type":                   string(identityType),
		"userAssignedIdentities": userAssignedIdentityIds,
	}
	return json.Marshal(out)
}

// ExpandSystemAssignedUserAssignedList expands the schema input into a SystemUserAssignedList struct
func ExpandSystemAssignedUserAssignedList(input []interface{}) (*SystemUserAssignedList, error) {
	identityType := TypeNone
	identityIds := make([]string, 0)

	if len(input) > 0 {
		raw := input[0].(map[string]interface{})
		typeRaw := raw["type"].(string)
		if typeRaw == string(TypeSystemAssigned) {
			identityType = TypeSystemAssigned
		}
		if typeRaw == string(TypeSystemAssignedUserAssigned) {
			identityType = TypeSystemAssignedUserAssigned
		}
		if typeRaw == string(TypeUserAssigned) {
			identityType = TypeUserAssigned
		}

		identityIdsRaw := raw["identity_ids"].([]interface{})
		for _, v := range identityIdsRaw {
			identityIds = append(identityIds, v.(string))
		}
	}

	if identityType == TypeNone {
		return &SystemUserAssignedList{
			Type:        TypeNone,
			IdentityIds: []string{},
		}, nil
	}

	return &SystemUserAssignedList{
		Type:        identityType,
		IdentityIds: identityIds,
	}, nil
}

// FlattenSystemAssignedUserAssignedList turns a SystemUserAssignedList into a []interface{}
func FlattenSystemAssignedUserAssignedList(input *SystemUserAssignedList) (*[]interface{}, error) {
	if input == nil || (input.Type != TypeSystemAssigned && input.Type != TypeSystemAssignedUserAssigned && input.Type != TypeUserAssigned) {
		return &[]interface{}{}, nil
	}

	identityIds := make([]string, 0)
	for _, raw := range input.IdentityIds {
		id, err := commonids.ParseUserAssignedIdentityIDInsensitively(raw)
		if err != nil {
			return nil, fmt.Errorf("parsing %q as a User Assigned Identity ID: %+v", raw, err)
		}
		identityIds = append(identityIds, id.ID())
	}

	return &[]interface{}{
		map[string]interface{}{
			"type":         string(input.Type),
			"identity_ids": identityIds,
			"principal_id": input.PrincipalId,
			"tenant_id":    input.TenantId,
		},
	}, nil
}
