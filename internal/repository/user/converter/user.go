package converter

import (
	"github.com/msh2107/auth/internal/model"
	modelRepo "github.com/msh2107/auth/internal/repository/user/model"
)

// ToUserFromRepo - .
func ToUserFromRepo(note *modelRepo.User) *model.User {
	return &model.User{
		ID:        note.ID,
		Info:      ToUserInfoFromRepo(note.Info),
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
}

// ToUserInfoFromRepo - .
func ToUserInfoFromRepo(info modelRepo.UserInfo) model.UserInfo {
	return model.UserInfo{
		Name:     info.Name,
		Email:    info.Email,
		Password: info.Password,
		Role:     info.Role,
	}
}
