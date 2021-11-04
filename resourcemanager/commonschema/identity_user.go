package commonschema

import (
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type UserAssignedIdentity struct {
	Type                    IdentityType `tfschema:"type"`
	UserAssignedIdentityIds []string     `tfschema:"identity_ids"`
}

func UserAssignedIdentitySchema() *schema.Schema {
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
					}, false),
				},
				"identity_ids": {
					Type:     schema.TypeSet,
					Required: true,
					Elem: &schema.Schema{
						Type:         schema.TypeString,
						ValidateFunc: commonids.ValidateUserAssignedIdentityID,
					},
				},
			},
		},
	}
}

func UserAssignedIdentitySchemaDataSource() *schema.Schema {
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
			},
		},
	}
}

func ExpandUserAssignedIdentity(input []interface{}) (*UserAssignedIdentity, error) {
	if len(input) == 0 || input[0] == nil {
		return &UserAssignedIdentity{
			Type: none,
		}, nil
	}

	v := input[0].(map[string]interface{})

	identityIds := make([]string, 0)
	for _, v := range v["identity_ids"].(*schema.Set).List() {
		identityIds = append(identityIds, v.(string))
	}

	return &UserAssignedIdentity{
		Type:                    userAssigned,
		UserAssignedIdentityIds: identityIds,
	}, nil
}

func FlattenUserAssignedIdentity(input *UserAssignedIdentity) []interface{} {
	if input == nil || strings.EqualFold(string(input.Type), string(none)) {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"type":         input.Type,
			"identity_ids": input.UserAssignedIdentityIds,
		},
	}
}
