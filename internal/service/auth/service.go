package auth

import (
	"github.com/msh2107/auth/internal/client/db"
	"github.com/msh2107/auth/internal/repository"
	"github.com/msh2107/auth/internal/service"
	"time"
)

const (
	refreshTokenExpiration = 60 * time.Minute
	accessTokenExpiration  = 5 * time.Minute
)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

// NewAuthService - .
func NewAuthService(userRepository repository.UserRepository, txManager db.TxManager) service.AuthService {
	return &serv{userRepository: userRepository, txManager: txManager}
}
