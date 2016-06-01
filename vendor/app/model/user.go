package model

import (
	"time"

)

type User struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created_at"`
}
