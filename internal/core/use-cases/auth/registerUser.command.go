package authservice

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

// RegisterUserCommand estrutura para o comando de registro de usuário
type RegisterUserCommand struct {
	userRepository ports.UserRepository
	logger         logger.Provider
}

type RegisterUserParams struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func NewRegisterUserCommand(userRepo ports.UserRepository, logger logger.Provider) *RegisterUserCommand {
	return &RegisterUserCommand{
		userRepository: userRepo,
		logger:         logger,
	}
}

func (cmd *RegisterUserCommand) Execute(params RegisterUserParams) (*domain.User, error) {
	existingUser, err := cmd.userRepository.GetByEmail(params.Email)
	if err == nil && existingUser != nil {
		cmd.logger.Error("Email already registered", errors.New("email already registered"))
		return nil, fmt.Errorf("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		cmd.logger.Error("Password encryption failed", err)
		return nil, fmt.Errorf("could not encrypt password: %w", err)
	}

	newUser := &domain.User{
		Email:    params.Email,
		Password: string(hashedPassword),
	}

	createdUser, err := cmd.userRepository.Create(newUser)
	if err != nil {
		cmd.logger.Error("User creation failed", err)
		return nil, fmt.Errorf("could not create user: %w", err)
	}

	cmd.logger.Info("User registered successfully", params.Email)
	return createdUser, nil
}
