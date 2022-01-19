package commonschema

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/location"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ExtendedLocation() *schema.Schema {
	return &schema.Schema{
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		StateFunc:        location.StateFunc,
		DiffSuppressFunc: location.DiffSuppressFunc,
		ValidateFunc:     validation.StringIsNotEmpty,
	}
}

func Location() *schema.Schema {
	return &schema.Schema{
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		ValidateFunc:     location.EnhancedValidate,
		StateFunc:        location.StateFunc,
		DiffSuppressFunc: location.DiffSuppressFunc,
	}
}

func LocationOptional() *schema.Schema {
	return &schema.Schema{
		Type:             schema.TypeString,
		Optional:         true,
		ForceNew:         true,
		StateFunc:        location.StateFunc,
		DiffSuppressFunc: location.DiffSuppressFunc,
	}
}

func LocationComputed() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func LocationWithoutForceNew() *schema.Schema {
	return &schema.Schema{
		Type:             schema.TypeString,
		Required:         true,
		ValidateFunc:     location.EnhancedValidate,
		StateFunc:        location.StateFunc,
		DiffSuppressFunc: location.DiffSuppressFunc,
	}
}
