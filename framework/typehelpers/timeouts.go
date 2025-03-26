package typehelpers

import "time"

type Timeouts struct {
	defaultCreateTimeout time.Duration
	defaultReadTimeout   time.Duration
	defaultUpdateTimeout time.Duration
	defaultDeleteTimeout time.Duration
}
