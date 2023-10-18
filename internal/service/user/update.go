package user

import (
	"context"

	"github.com/msh2107/auth/internal/model"
)

func (s *serv) Update(ctx context.Context, id int64, info *model.UserInfo) error {
	err := s.userRepository.Update(ctx, id, info)
	if err != nil {
		return err
	}

	return nil
}
