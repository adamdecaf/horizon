package storage

import (
	"time"
)

type TwitterUser struct {
	CreatedAt time.Time
	Id string
	Name string
	ScreenName string
}

type BasicTweet struct {
	CreatedAt time.Time
	Id string
	Text string
	User TwitterUser
}
