package service

import (
	"context"

	"github.com/msh2107/auth/internal/model"
)

// UserService - .
type UserService interface {
	Create(ctx context.Context, info *model.UserInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, id int64, info *model.UserInfo) error
	Delete(ctx context.Context, id int64) error
}

type AuthService interface {
	Login(ctx context.Context, info *model.UserInfo) (string, error)
	GetRefreshToken(ctx context.Context, RefreshToken string) (string, error)
	GetAccessToken(ctx context.Context, RefreshToken string) (string, error)
}

type AccessService interface {
	Check(ctx context.Context, address string) error
}
