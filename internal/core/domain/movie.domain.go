package domain

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	Token     uuid.UUID `json:"token" db:"token"`
	Name      string    `json:"name" db:"name"`
	Director  string    `json:"director" db:"director"`
	Duration  int       `json:"duration" db:"duration"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
