package domain

import (
	"time"

	"github.com/google/uuid"
)

type Threater struct {
	Token       uuid.UUID `json:"token" db:"token"`
	Number      int       `json:"number" db:"number"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}
