package tracrpc

import "time"

// String returns the string pointer.
func String(val string) *string { return &val }

// Int returns the int pointer.
func Int(val int) *int { return &val }

// Bool returns the bool pointer.
func Bool(val bool) *bool { return &val }

// Time returns the time.Time pointer.
func Time(val time.Time) *time.Time { return &val }
