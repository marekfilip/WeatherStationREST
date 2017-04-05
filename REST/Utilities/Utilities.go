package Utilities

import "time"

const DURATION_LIMIT time.Duration = time.Duration(48) * time.Hour

func IsLimitExceeded(from, to time.Time) bool {
	return to.Sub(from) > DURATION_LIMIT
}
