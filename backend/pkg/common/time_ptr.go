package common

import "time"

func TimeDurationPtr(v time.Duration) *time.Duration {
	return new(v)
}
