package tests

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/msh2107/auth/internal/model"
	"github.com/msh2107/auth/internal/repository"
	repoMocks "github.com/msh2107/auth/internal/repository/mocks"
	"github.com/msh2107/auth/internal/service/user"
	desc "github.com/msh2107/auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGet(t *testing.T) {
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		password  = gofakeit.Password(true, true, true, true, true, 10)
		role      = desc.Role(gofakeit.Int32())
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		repoErr = fmt.Errorf("repo error")

		resp = &model.User{
			ID: id,
			Info: model.UserInfo{
				Name:     name,
				Email:    email,
				Password: password,
				Role:     role,
			},
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{
				Time:  updatedAt,
				Valid: true,
			},
		}
	)
	t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		want               *model.User
		err                error
		userRepositoryMock userRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: resp,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, id).Return(resp, nil)
				return mock
			},
		},
		{
			name: "repo error case",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: nil,
			err:  repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepoMock := tt.userRepositoryMock(mc)
			service := user.NewMockUserService(userRepoMock)

			id, err := service.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, id)
		})
	}
}
