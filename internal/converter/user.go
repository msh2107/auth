package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/msh2107/auth/internal/model"
	desc "github.com/msh2107/auth/pkg/user_v1"
)

// ToUserFromService - .
func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id:        user.ID,
		Info:      ToUserInfoFromService(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

// ToUserInfoFromService - .
func ToUserInfoFromService(info model.UserInfo) *desc.UserInfo {
	return &desc.UserInfo{
		Name:     info.Name,
		Email:    info.Email,
		Password: info.Password,
		Role:     info.Role,
	}
}

// ToUserInfoFromDesc - .
func ToUserInfoFromDesc(info *desc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Name:     info.Name,
		Email:    info.Email,
		Password: info.Password,
		Role:     info.Role,
	}
}

// ToUserFromUpdateDesc - .
func ToUserFromUpdateDesc(info *desc.UpdateUserInfo) *model.UserInfo {
	return &model.UserInfo{
		Name:     info.GetName().Value,
		Email:    info.GetEmail().GetValue(),
		Password: info.GetPassword().GetValue(),
		Role:     info.GetRole(),
	}
}
