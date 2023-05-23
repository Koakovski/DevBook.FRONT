package model

import (
	"time"
)

type User struct {
	Id           uint64        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	NickName     string        `json:"nickName,omitempty"`
	Email        string        `json:"email,omitempty"`
	Password     string        `json:"password,omitempty"`
	CreatedAt    time.Time     `json:"createdAt,omitempty"`
	Followers    []User        `json:"followers"`
	Following    []User        `json:"following"`
	Publications []Publication `json:"publications"`
}
