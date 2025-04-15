// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package typehelpers

import "time"

type Timeouts struct {
	defaultCreateTimeout time.Duration
	defaultReadTimeout   time.Duration
	defaultUpdateTimeout time.Duration
	defaultDeleteTimeout time.Duration
}
