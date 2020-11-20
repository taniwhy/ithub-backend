package clock

import "time"

// Now : 現在時刻
var Now = func() time.Time {
	return time.Now().UTC()
}
