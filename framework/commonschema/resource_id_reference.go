package commonschema

import (
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func ResourceIDReferenceRequired() resourceschema.StringAttribute {
	return resourceschema.StringAttribute{
		Required:   true,
		Validators: []validator.String{},
	}
}
