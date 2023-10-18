package user

import (
	"github.com/msh2107/auth/internal/repository"
	"github.com/msh2107/auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
}

// NewService - .
func NewService(userRepository repository.UserRepository) service.UserService {
	return &serv{userRepository: userRepository}
}
