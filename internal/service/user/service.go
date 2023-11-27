package user

import (
	"github.com/msh2107/auth/internal/client/db"
	"github.com/msh2107/auth/internal/repository"
	"github.com/msh2107/auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

// NewService - .
func NewService(userRepository repository.UserRepository, txManager db.TxManager) service.UserService {
	return &serv{userRepository: userRepository, txManager: txManager}
}

func NewMockService(deps ...interface{}) service.UserService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.UserRepository:
			srv.userRepository = s
		}
	}

	return &srv
}
