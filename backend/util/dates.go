package util

import "time"

func DateAfter(first, second time.Time) bool {
	return first.After(second)
}

func DateBefore(first, second time.Time) bool {
	return first.Before(second)
}
