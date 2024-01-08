package model

import "time"

// User represents a user on the social network
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	NickName  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Senha     string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
