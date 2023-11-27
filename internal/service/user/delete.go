package user

import (
	"context"
	"errors"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		exist, errTx := s.userRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		if exist == nil {
			return errors.New("user not found")
		}

		errTx = s.userRepository.Delete(ctx, id)
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
