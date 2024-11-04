package domain

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	Admin  Role = "admin"
	Client Role = "client"
	Guest  Role = "guest"
)

type User struct {
	Token     uuid.UUID `json:"token" database:"token"`
	Username  string    `json:"username" database:"username"`
	Email     string    `json:"email" database:"email"`
	Password  string    `json:"password" database:"password"`
	Roles     Role      `json:"role" database:"role"`
	CreatedAt time.Time `json:"createdAt" database:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" database:"updated_at"`
	LastLogin time.Time `json:"lastLogin" database:"last_login"`
	Active    bool      `json:"active" database:"active"`
}
