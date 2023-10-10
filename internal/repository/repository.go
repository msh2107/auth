package repository

import (
	"context"

	"github.com/msh2107/auth/internal/models"
)

// Repository - .
type Repository interface {
	Create(ctx context.Context, user models.User) (int64, error)
	Get(ctx context.Context, id int64) (models.User, error)
	Update(ctx context.Context, user models.User) error
	Delete(ctx context.Context, id int64) error
}
