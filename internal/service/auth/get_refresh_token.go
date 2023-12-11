package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/msh2107/auth/internal/utils"
	"os"
)

func (s *serv) GetRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	refreshTokenSecretKey := os.Getenv("REFRESH_TOKEN_SECRET_KEY")
	if len(refreshTokenSecretKey) == 0 {
		return "", errors.New("refresh token secret key  not found ")
	}

	claims, err := utils.VerifyToken(refreshToken, []byte(refreshTokenSecretKey))
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	user, err := s.userRepository.GetByName(ctx, claims.Username)
	if err != nil {
		return "", err
	}

	newRefreshToken, err := utils.GenerateToken(user.Info, []byte(refreshTokenSecretKey), refreshTokenExpiration)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %s", err.Error())
	}

	return newRefreshToken, nil
}
