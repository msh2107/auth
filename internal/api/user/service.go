package user

import (
	"github.com/msh2107/auth/internal/service"
	desc "github.com/msh2107/auth/pkg/user_v1"
)

// Implementation - .
type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

// NewUserImplementation - .
func NewUserImplementation(userService service.UserService) *Implementation {
	return &Implementation{userService: userService}
}
