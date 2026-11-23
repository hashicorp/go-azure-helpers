// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package network

// PublicNetworkAccess specifies whether a given Azure Resource is publicly accessible (fully/partially) or
// private.
type PublicNetworkAccess string

const (
	// PublicNetworkAccessDisabled specifies that Public Network Access is Disabled.
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"

	// PublicNetworkAccessEnabled specifies that Public Network Access is Enabled.
	PublicNetworkAccessEnabled PublicNetworkAccess = "Enabled"

	// PublicNetworkAccessSecuredByPerimeter specifies that Public Network Access is controlled by
	// the Network Security Perimeter.
	PublicNetworkAccessSecuredByPerimeter PublicNetworkAccess = "SecuredByPerimeter"
)
