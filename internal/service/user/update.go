package user

import (
	"context"
	"errors"
	"github.com/msh2107/auth/internal/model"
)

func (s *serv) Update(ctx context.Context, id int64, info *model.UserInfo) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		exist, errTx := s.userRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		if exist == nil {
			return errors.New("user not found")
		}

		updateExistInfo(exist, info)

		errTx = s.userRepository.Update(ctx, id, &exist.Info)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func updateExistInfo(old *model.User, new *model.UserInfo) {
	if new.Name != "" {
		old.Info.Name = new.Name
	}

	if new.Email != "" {
		old.Info.Email = new.Email
	}

	if new.Role != 0 {
		old.Info.Role = new.Role
	}
}
