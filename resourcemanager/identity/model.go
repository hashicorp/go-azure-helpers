// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identity

import "encoding/json"

var _ json.Marshaler = UserAssignedIdentityDetails{}

// use a random pointer value
var nullStringPtr = func(s string) *string { return &s }("")

var NullUserAssignedIdentityDetails = UserAssignedIdentityDetails{
	ClientId:    nullStringPtr,
	PrincipalId: nullStringPtr,
}

type UserAssignedIdentityDetails struct {
	ClientId    *string `json:"clientId,omitempty"`
	PrincipalId *string `json:"principalId,omitempty"`
}

func (u UserAssignedIdentityDetails) MarshalJSON() ([]byte, error) {
	// none of these properties can be set, so we'll just flatten an empty struct
	if u == NullUserAssignedIdentityDetails {
		return []byte("null"), nil
	}
	return json.Marshal(map[string]interface{}{})
}
