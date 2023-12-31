package repository

import (
	"context"

	"github.com/msh2107/auth/internal/model"
)

// UserRepository - .
type UserRepository interface {
	Create(ctx context.Context, info *model.UserInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, id int64, info *model.UserInfo) error
	Delete(ctx context.Context, id int64) error
	GetByName(ctx context.Context, name string) (*model.User, error)
}
