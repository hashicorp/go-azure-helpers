// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identity

import (
	"reflect"
	"testing"
)

func ptr(s string) *string {
	return &s
}

func TestUserAssignedIdentityDetails_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		identity  UserAssignedIdentityDetails
		want    []byte
		wantErr bool
	}{
		{
			name: "empty identity",
			identity: UserAssignedIdentityDetails{},
			want: []byte("{}"),
		},
		{
			name: "nil identity",
			identity: NullUserAssignedIdentityDetails,
			want: []byte("null"),
		},
		{
			name: "other value identity",
			identity: UserAssignedIdentityDetails{
				ClientId:    ptr("other-client-id"),
			},
			want: []byte("{}"),
		},
		{
			name: "other empty identity",
			identity: UserAssignedIdentityDetails{
				ClientId:    ptr(""),
				PrincipalId: ptr(""),
			},
			want: []byte("{}"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t2 *testing.T) {
			u := tt.identity
			got, err := u.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t2.Errorf("UserAssignedIdentityDetails.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t2.Errorf("UserAssignedIdentityDetails.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
