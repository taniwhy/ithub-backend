package clock

import "time"

// Now : 現在時刻の生成
var Now = func() time.Time {
	return time.Now().UTC()
}
