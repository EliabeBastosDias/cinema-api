package authservice

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	jwthandler "github.com/EliabeBastosDias/cinema-api/internal/core/utils/jwt"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type LoginCommand struct {
	userRepository ports.UserRepository
	jwtHandler     jwthandler.JWTHandler
	logger         logger.Provider
}

type LoginParams struct {
	Email    string
	Password string
}

func NewLoginCommand(userRepo ports.UserRepository, jwtHandler jwthandler.JWTHandler, logger logger.Provider) *LoginCommand {
	return &LoginCommand{
		userRepository: userRepo,
		jwtHandler:     jwtHandler,
		logger:         logger,
	}
}

type LoginResult struct {
	User  *domain.User
	Token string
}

func (cmd *LoginCommand) Execute(params LoginParams) (*LoginResult, error) {
	cmd.logger.Info("LoginCommand initiated", params.Email)

	user, err := cmd.userRepository.GetByEmail(params.Email)
	if err != nil {
		cmd.logger.Error("LoginCommand failed: user not found", err)
		return nil, fmt.Errorf("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		cmd.logger.Error("LoginCommand failed: invalid password", err)
		return nil, fmt.Errorf("invalid credentials")
	}

	user.LastLogin = time.Now()
	err = cmd.userRepository.Update(user)
	if err != nil {
		cmd.logger.Error("LoginCommand failed: could not update last login", err)
		return nil, fmt.Errorf("could not update last login: %w", err)
	}

	token, err := cmd.jwtHandler.GenerateToken(user)
	if err != nil {
		cmd.logger.Error("LoginCommand failed: could not generate JWT", err)
		return nil, fmt.Errorf("could not generate JWT: %w", err)
	}

	cmd.logger.Info("LoginCommand succeeded", params.Email)

	return &LoginResult{
		User:  user,
		Token: token,
	}, nil
}
