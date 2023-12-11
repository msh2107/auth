package model

import (
	"github.com/dgrijalva/jwt-go"
	desc "github.com/msh2107/auth/pkg/user_v1"
)

const authPrefix = "Bearer "

type UserClaims struct {
	jwt.StandardClaims
	Username string    `json:"username"`
	Role     desc.Role `json:"role"`
}
