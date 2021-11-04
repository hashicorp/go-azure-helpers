package commonschema

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type SystemAssignedIdentity struct {
	Type        IdentityType `tfschema:"type"`
	PrincipalId string       `tfschema:"principal_id"`
	TenantId    string       `tfschema:"tenant_id"`
}

func SystemAssignedIdentitySchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeString,
					Required: true,
					ValidateFunc: validation.StringInSlice([]string{
						string(systemAssigned),
					}, false),
				},
				"principal_id": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"tenant_id": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}

func SystemAssignedIdentitySchemaDataSource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"principal_id": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"tenant_id": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}

func ExpandSystemAssignedIdentity(input []interface{}) (*SystemAssignedIdentity, error) {
	if len(input) == 0 || input[0] == nil {
		return &SystemAssignedIdentity{
			Type: none,
		}, nil
	}

	return &SystemAssignedIdentity{
		Type: systemAssigned,
	}, nil
}

func FlattenSystemAssignedIdentity(input *SystemAssignedIdentity) []interface{} {
	if input == nil || strings.EqualFold(string(input.Type), string(none)) {
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
