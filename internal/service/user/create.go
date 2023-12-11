package user

import (
	"context"
	"github.com/msh2107/auth/internal/utils"

	"github.com/msh2107/auth/internal/model"
)

func (s *serv) Create(ctx context.Context, info *model.UserInfo) (int64, error) {
	var id int64
	passwordHash, err := utils.HashPassword(info.Password)
	if err != nil {
		return 0, err
	}
	info.Password = passwordHash

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.userRepository.Create(ctx, info)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.userRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
