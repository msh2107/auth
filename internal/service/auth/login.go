package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/msh2107/auth/internal/model"
	"github.com/msh2107/auth/internal/utils"
	"os"
)

func (s *serv) Login(ctx context.Context, info *model.UserInfo) (string, error) {
	refreshTokenSecretKey := os.Getenv("REFRESH_TOKEN_SECRET_KEY")
	if len(refreshTokenSecretKey) == 0 {
		return "", errors.New("refresh token secret key  not found ")
	}
	infoFromDB, err := s.userRepository.GetByName(ctx, info.Name)
	if err != nil {
		return "", err
	}

	if utils.VerifyPassword(infoFromDB.Info.Password, info.Password) {
		refreshToken, err := utils.GenerateToken(*info, []byte(refreshTokenSecretKey), refreshTokenExpiration)
		if err != nil {
			return "", fmt.Errorf("failed to generate token: %s", err.Error())
		}

		return refreshToken, nil
	}
	return "", errors.New("incorrect password")
}
