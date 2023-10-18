package user

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/msh2107/auth/internal/repository"
	"github.com/msh2107/auth/internal/repository/user/converter"
	modelRepo "github.com/msh2107/auth/internal/repository/user/model"

	"github.com/msh2107/auth/internal/model"
)

// repo - .
type repo struct {
	pool *pgxpool.Pool
}

// NewRepository - .
func NewRepository(pool *pgxpool.Pool) repository.UserRepository {
	return &repo{pool: pool}
}

// Create - .
func (r *repo) Create(ctx context.Context, info *model.UserInfo) (int64, error) {
	builderInsert := sq.Insert("\"user\"").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "email", "password", "role").
		Values(info.Name, info.Email, info.Password, info.Role).Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	var id int64
	err = r.pool.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Get - .
func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builderSelectOne := sq.Select("id", "name", "email", "password", "role", "created_at", "updated_at").
		From("\"user\"").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		return nil, err
	}

	var user modelRepo.User
	err = r.pool.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Info.Name, &user.Info.Email, &user.Info.Password, &user.Info.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

// Update - .
func (r *repo) Update(ctx context.Context, id int64, info *model.UserInfo) error {
	builderUpdate := sq.Update("\"user\"").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		Set("updated_at", time.Now())
	if info.Name != "" {
		builderUpdate = builderUpdate.Set("name", info.Name)
	}
	if info.Email != "" {
		builderUpdate = builderUpdate.Set("email", info.Email)
	}
	if info.Role != 0 {
		builderUpdate = builderUpdate.Set("role", info.Role)
	}
	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

// Delete - .
func (r *repo) Delete(ctx context.Context, id int64) error {
	builderDelete := sq.Delete("\"user\"").Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar)

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
