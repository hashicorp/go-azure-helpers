package identity

import (
	"encoding/json"
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
