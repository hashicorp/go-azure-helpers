// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package commonschema

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// NOTE: we intentionally don't have an Optional & Computed here for behavioural consistency.

// NOTE: there's two different types of SystemOrSingleUserAssignedIdentity supported by Azure:
// The first (List) represents the IdentityIDs as a List of Strings
// The other (Map) represents the IdentityIDs as a Map of String : Object (containing Client/PrincipalID)
// from a users perspective however, these should both be represented using the same schema
// so we have a single schema and separate Expand/Flatten functions

// SystemOrSingleUserAssignedIdentityRequired returns the System or User Assigned Identity schema where this is Required
func SystemOrSingleUserAssignedIdentityRequired() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeString,
					Required: true,
					ValidateFunc: validation.StringInSlice([]string{
						string(identity.TypeUserAssigned),
						string(identity.TypeSystemAssigned),
					}, false),
				},
				"identity_ids": {
					Type:     schema.TypeSet,
					Optional: true,
					MaxItems: 1,
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

// SystemOrSingleUserAssignedIdentityRequiredForceNew returns the System or User Assigned Identity schema where this is Required and ForceNew
func SystemOrSingleUserAssignedIdentityRequiredForceNew() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		ForceNew: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeString,
					Required: true,
					ForceNew: true,
					ValidateFunc: validation.StringInSlice([]string{
						string(identity.TypeUserAssigned),
						string(identity.TypeSystemAssigned),
					}, false),
				},
				"identity_ids": {
					Type:     schema.TypeSet,
					Optional: true,
					ForceNew: true,
					MaxItems: 1,
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

// SystemOrSingleUserAssignedIdentityOptional returns the System or User Assigned Identity schema where this is Optional
func SystemOrSingleUserAssignedIdentityOptional() *schema.Schema {
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
						string(identity.TypeUserAssigned),
						string(identity.TypeSystemAssigned),
					}, false),
				},
				"identity_ids": {
					Type:     schema.TypeSet,
					Optional: true,
					MaxItems: 1,
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

// SystemOrSingleUserAssignedIdentityOptionalForceNew returns the System or User Assigned Identity schema where this is Optional and ForceNew
func SystemOrSingleUserAssignedIdentityOptionalForceNew() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		ForceNew: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeString,
					Required: true,
					ForceNew: true,
					ValidateFunc: validation.StringInSlice([]string{
						string(identity.TypeUserAssigned),
						string(identity.TypeSystemAssigned),
					}, false),
				},
				"identity_ids": {
					Type:     schema.TypeSet,
					Optional: true,
					ForceNew: true,
					MaxItems: 1,
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

// SystemOrSingleUserAssignedIdentityComputed returns the System or User Assigned Identity schema where this is Computed
func SystemOrSingleUserAssignedIdentityComputed() *schema.Schema {
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
