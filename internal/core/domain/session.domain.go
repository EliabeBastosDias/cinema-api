package domain

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Token           uuid.UUID `json:"token" db:"token"`
	MovieToken      string    `json:"movie_token" db:"movie_token"`
	ThreaterToken   string    `json:"thread_token" db:"thread_token"`
	SessionDatetime time.Time `json:"session_datetime" db:"session_datetime"`
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time `json:"updatedAt" db:"updated_at"`
	Active          bool      `json:"active" db:"active"`
}
