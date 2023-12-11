package tests

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/msh2107/auth/internal/api/user"
	"github.com/msh2107/auth/internal/model"
	"github.com/msh2107/auth/internal/service"
	serviceMocks "github.com/msh2107/auth/internal/service/mocks"
	desc "github.com/msh2107/auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
)

func TestGet(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService
	type args struct {
		ctx context.Context
		req *desc.GetRequest
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

		serviceErr = fmt.Errorf("service error")

		req = &desc.GetRequest{
			Id: id,
		}

		serviceRes = &model.User{
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

		res = &desc.GetResponse{
			User: &desc.User{
				Id: id,
				Info: &desc.UserInfo{
					Name:     name,
					Email:    email,
					Password: password,
					Role:     role,
				},
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: timestamppb.New(updatedAt),
			}}
	)

	t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(serviceRes, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userServiceMock := tt.userServiceMock(mc)
			api := user.NewUserImplementation(userServiceMock)

			u, err := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, u)
			require.Equal(t, tt.err, err)
		})
	}
}
