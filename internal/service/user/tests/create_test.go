package tests

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/msh2107/auth/internal/client/db"
	txMocks "github.com/msh2107/auth/internal/client/db/mocks"
	"github.com/msh2107/auth/internal/client/db/pg"
	"github.com/msh2107/auth/internal/client/db/transaction"
	"github.com/msh2107/auth/internal/model"
	"github.com/msh2107/auth/internal/repository"
	repoMocks "github.com/msh2107/auth/internal/repository/mocks"
	"github.com/msh2107/auth/internal/service/user"
	desc "github.com/msh2107/auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type txTransactorMockFunc func(mc *minimock.Controller) db.Transactor

	type args struct {
		ctx context.Context
		req *model.UserInfo
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, true, 10)
		role     = desc.Role(gofakeit.Int32())

		txM TxMock

		repoErr = fmt.Errorf("repo error")

		req = &model.UserInfo{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     role,
		}
	)
	t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		userRepositoryMock userRepositoryMockFunc
		txTransactorMock   txTransactorMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: id,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(pg.MakeContextTx(ctx, &txM), req).Return(id, nil)
				mock.GetMock.Return(nil, nil)
				return mock
			},
			txTransactorMock: func(mc *minimock.Controller) db.Transactor {
				mock := txMocks.NewTransactorMock(mc)
				mock.BeginTxMock.Return(&txM, nil)
				return mock
			},
		},

		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				return mock
			},
			txTransactorMock: func(mc *minimock.Controller) db.Transactor {
				mock := txMocks.NewTransactorMock(mc)
				mock.BeginTxMock.Return(&txM, repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := tt.userRepositoryMock(mc)
			txTrans := transaction.NewTransactionManager(tt.txTransactorMock(mc))
			service := user.NewService(userRepo, txTrans)
			id, err := service.Create(tt.args.ctx, tt.args.req)
			if err != nil {
				return
			}

			require.ErrorIs(t, err, tt.err)
			require.Equal(t, id, tt.want)
		})
	}
}
