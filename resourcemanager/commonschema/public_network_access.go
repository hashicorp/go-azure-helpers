// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package commonschema

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/network"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// PublicNetworkAccessOptional returns the schema for a `public_network_access` field that is Optional.
func PublicNetworkAccessOptional(supportsSecuredByPerimeter bool) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Default:      string(network.PublicNetworkAccessEnabled),
		ValidateFunc: validationFunctionForPublicNetworkAccess(supportsSecuredByPerimeter),
	}
}

// PublicNetworkAccessOptionalForceNew returns the schema for a `public_network_access` field that
// is both Optional and ForceNew.
func PublicNetworkAccessOptionalForceNew(supportsSecuredByPerimeter bool) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Default:      string(network.PublicNetworkAccessEnabled),
		ForceNew:     true,
		ValidateFunc: validationFunctionForPublicNetworkAccess(supportsSecuredByPerimeter),
	}
}

// PublicNetworkAccessRequired returns the schema for a `public_network_access` field that is Required.
func PublicNetworkAccessRequired(supportsSecuredByPerimeter bool) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validationFunctionForPublicNetworkAccess(supportsSecuredByPerimeter),
	}
}

// PublicNetworkAccessRequiredForceNew returns the schema for a `public_network_access` field that
// is both Required and ForceNew.
func PublicNetworkAccessRequiredForceNew(supportsSecuredByPerimeter bool) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validationFunctionForPublicNetworkAccess(supportsSecuredByPerimeter),
	}
}

// validationFunctionForPublicNetworkAccess returns the validation function for the `public_network_access` field
func validationFunctionForPublicNetworkAccess(supportsSecuredByPerimeter bool) schema.SchemaValidateFunc {
	if supportsSecuredByPerimeter {
		return validation.StringInSlice([]string{
			string(network.PublicNetworkAccessDisabled),
			string(network.PublicNetworkAccessEnabled),
			string(network.PublicNetworkAccessSecuredByPerimeter),
		}, false)
	}

	return validation.StringInSlice([]string{
		string(network.PublicNetworkAccessDisabled),
		string(network.PublicNetworkAccessEnabled),
	}, false)
}
