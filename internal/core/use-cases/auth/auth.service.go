package authservice

import (
	"github.com/EliabeBastosDias/cinema-api/internal/core/domain"
	"github.com/EliabeBastosDias/cinema-api/internal/core/ports"
	jwthandler "github.com/EliabeBastosDias/cinema-api/internal/core/utils/jwt"
	"github.com/EliabeBastosDias/cinema-api/internal/repositories"
	"github.com/EliabeBastosDias/cinema-api/pkg/logger"
)

type AuthService struct {
	userRepository ports.UserRepository
	JWTHandler     jwthandler.JWTHandler
	logger         logger.Provider
}

func NewAuthService(repo repositories.RepoProvider, jwtHandler jwthandler.JWTHandler, logger logger.Provider) *AuthService {
	return &AuthService{
		userRepository: repo.User(),
		JWTHandler:     jwtHandler,
		logger:         logger,
	}
}

func (s *AuthService) Login(params LoginParams) (*LoginResult, error) {
	return NewLoginCommand(
		s.userRepository,
		s.JWTHandler,
		s.logger,
	).Execute(params)
}

func (s *AuthService) RegisterUser(params RegisterUserParams) (*domain.User, error) {
	return NewRegisterUserCommand(
		s.userRepository,
		s.logger,
	).Execute(params)
}
