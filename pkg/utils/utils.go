package utils

import (
	"fmt"
	"time"
)

func ToPointer[T any](v T) *T {
	return &v
}

func ToValue[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}

func GetTimeUnix(unix int64) time.Time {
	return time.Unix(unix, 0)
}

func ToString[T any](v T) string {
	return fmt.Sprintf("%v", v)
}
