package res

import "time"

// FollowsRes :
type FollowsRes struct {
	UserName  string
	Name      string
	UserText  string
	UserIcon  string
	CreatedAt time.Time
}
