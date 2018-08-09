package types

import "time"

func Now() time.Time {
	return time.Now().Round(0).UTC()
}
