package model

import (
	"database/sql"
	"time"

	desc "github.com/msh2107/auth/pkg/user_v1"
)

// User - .
type User struct {
	ID        int64
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

// UserInfo - .
type UserInfo struct {
	Name     string
	Email    string
	Password string
	Role     desc.Role
}
