package pg

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/msh2107/auth/internal/models"
)

// UserRepository - .
type UserRepository struct {
	pool *pgxpool.Pool
}

// NewUserRepository - .
func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

// Create - .
func (r *UserRepository) Create(ctx context.Context, user models.User) (int64, error) {
	builderInsert := sq.Insert("\"user\"").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "email", "password", "role").
		Values(user.Name, user.Email, user.Password, user.Role).Suffix("RETURNING id")

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
func (r *UserRepository) Get(ctx context.Context, id int64) (models.User, error) {
	builderSelectOne := sq.Select("id", "name", "email", "password", "role", "created_at", "updated_at").
		From("\"user\"").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	var updatedAt sql.NullTime
	err = r.pool.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &updatedAt)
	if err != nil {
		return models.User{}, err
	}

	if updatedAt.Valid {
		user.UpdatedAt = updatedAt.Time
	}

	return user, nil
}

// Update - .
func (r *UserRepository) Update(ctx context.Context, user models.User) error {
	builderUpdate := sq.Update("\"user\"").
		Set("name", user.Name).
		Set("email", user.Email).
		Set("role", user.Role).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": user.ID}).
		PlaceholderFormat(sq.Dollar)

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
func (r *UserRepository) Delete(ctx context.Context, id int64) error {
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
