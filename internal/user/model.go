package user

import "time"

type AppUser struct {
	Id              int64
	IsAdmin         bool
	IsBanned        bool
	State           string
	LastMessageDate time.Time
	RegisterDate    time.Time
}
