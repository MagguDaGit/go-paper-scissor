package util

import (
	"fmt"
	"time"
)

func DateTimeNowString() string {
	return time.Now().Format("02.01.2006 15:04:05")
}

func DurationToString(duration time.Duration) string {
	switch {
	case duration < time.Microsecond:
		return fmt.Sprintf("DURATION: %dns ", duration.Nanoseconds())
	case duration < time.Millisecond:
		return fmt.Sprintf("DURATION: %dµs ", duration.Microseconds())
	case duration < time.Second:
		return fmt.Sprintf("DURATION: %dms ", duration.Milliseconds())
	default:
		return fmt.Sprintf("DURATION: %.2fs ", duration.Seconds())
	}
}
