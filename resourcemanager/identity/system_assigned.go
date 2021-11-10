package identity

import (
	"encoding/json"
	"strings"
)

var _ json.Marshaler = &SystemAssigned{}

type SystemAssigned struct {
	Type        Type   `tfschema:"type" json:"type"`
	PrincipalId string `tfschema:"principal_id" json:"principalId"`
	TenantId    string `tfschema:"tenant_id" json:"tenantId"`
}

func (s *SystemAssigned) MarshalJSON() ([]byte, error) {
	// we use a custom marshal function here since we can only send the Type field
	out := map[string]interface{}{
		"type": string(TypeNone),
	}
	if s != nil && s.Type == TypeSystemAssigned {
		out["type"] = string(TypeSystemAssigned)
	}
	return json.Marshal(out)
}

func ExpandSystemAssigned(input []interface{}) (*SystemAssigned, error) {
	if len(input) == 0 || input[0] == nil {
		return &SystemAssigned{
			Type: TypeNone,
		}, nil
	}

	return &SystemAssigned{
		Type: TypeSystemAssigned,
	}, nil
}

func FlattenSystemAssigned(input *SystemAssigned) []interface{} {
	if input == nil || strings.EqualFold(string(input.Type), string(TypeNone)) {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"type":         input.Type,
			"principal_id": input.PrincipalId,
			"tenant_id":    input.TenantId,
		},
	}
}