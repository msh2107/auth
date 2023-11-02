package user

import (
	"context"
	"github.com/msh2107/auth/internal/client/db"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/msh2107/auth/internal/repository"
	"github.com/msh2107/auth/internal/repository/user/converter"
	modelRepo "github.com/msh2107/auth/internal/repository/user/model"

	"github.com/msh2107/auth/internal/model"
)

// repo - .
type repo struct {
	db db.Client
}

// NewRepository - .
func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
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

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
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

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&user.ID, &user.Info.Name, &user.Info.Email, &user.Info.Password, &user.Info.Role, &user.CreatedAt, &user.UpdatedAt)
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
		Set("updated_at", time.Now()).
		Set("name", info.Name).
		Set("email", info.Email).
		Set("role", info.Role)

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
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

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
