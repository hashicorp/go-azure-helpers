package commonschema

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type SystemAssignedUserAssignedIdentity struct {
	Type                    IdentityType `tfschema:"type"`
	PrincipalId             string       `tfschema:"principal_id"`
	TenantId                string       `tfschema:"tenant_id"`
	UserAssignedIdentityIds []string     `tfschema:"identity_ids"`
}

func SystemAssignedUserAssignedIdentitySchema() *schema.Schema {
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
						string(userAssigned),
						string(systemAssigned),
						string(systemAssignedUserAssigned),
					}, false),
				},
				"identity_ids": {
					Type:     schema.TypeSet,
					Optional: true,
					Elem: &schema.Schema{
						Type:         schema.TypeString,
						ValidateFunc: commonids.ValidateUserAssignedIdentityID,
					},
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

func SystemAssignedUserAssignedIdentitySchemaDataSource() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"identity_ids": {
					Type:     schema.TypeList,
					Computed: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
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

func ExpandSystemAssignedUserAssignedIdentity(input []interface{}) (*SystemAssignedUserAssignedIdentity, error) {
	if len(input) == 0 || input[0] == nil {
		return &SystemAssignedUserAssignedIdentity{
			Type: none,
		}, nil
	}

	v := input[0].(map[string]interface{})

	config := &SystemAssignedUserAssignedIdentity{
		Type: IdentityType(v["type"].(string)),
	}

	identityIdsRaw := v["identity_ids"].(*schema.Set).List()

	if len(identityIdsRaw) != 0 {
		if config.Type != userAssigned && config.Type != systemAssignedUserAssigned {
			return nil, fmt.Errorf("`identity_ids` can only be specified when `type` includes `UserAssigned`")
		}

		identityIds := make([]string, 0)
		for _, v := range identityIdsRaw {
			identityIds = append(identityIds, v.(string))
		}

		config.UserAssignedIdentityIds = identityIds
	}

	return config, nil
}

func FlattenSystemAssignedUserAssignedIdentity(input *SystemAssignedUserAssignedIdentity) []interface{} {
	if input == nil || strings.EqualFold(string(input.Type), string(none)) {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"type":         input.Type,
			"identity_ids": input.UserAssignedIdentityIds,
			"principal_id": input.PrincipalId,
			"tenant_id":    input.TenantId,
		},
	}
}
