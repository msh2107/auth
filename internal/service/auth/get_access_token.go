package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/msh2107/auth/internal/utils"
	"os"
)

func (s *serv) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	refreshTokenSecretKey := os.Getenv("REFRESH_TOKEN_SECRET_KEY")
	if len(refreshTokenSecretKey) == 0 {
		return "", errors.New("refresh token secret key  not found ")
	}

	accessTokenSecretKey := os.Getenv("ACCESS_TOKEN_SECRET_KEY")
	if len(accessTokenSecretKey) == 0 {
		return "", errors.New("access token secret key  not found ")
	}

	claims, err := utils.VerifyToken(refreshToken, []byte(refreshTokenSecretKey))
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	user, err := s.userRepository.GetByName(ctx, claims.Username)
	if err != nil {
		return "", err
	}

	accessToken, err := utils.GenerateToken(user.Info, []byte(accessTokenSecretKey), accessTokenExpiration)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %s", err.Error())
	}

	return accessToken, nil
}
