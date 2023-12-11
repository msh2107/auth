package tests

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/msh2107/auth/internal/api/user"
	"github.com/msh2107/auth/internal/model"
	"github.com/msh2107/auth/internal/service"
	"github.com/msh2107/auth/internal/service/mocks"
	desc "github.com/msh2107/auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"testing"
)

func TestUpdate(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService
	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, true, 10)
		role     = desc.Role(gofakeit.Int32())

		serviceErr = fmt.Errorf("service error")

		req = &desc.UpdateRequest{
			Id: id,
			Info: &desc.UpdateUserInfo{
				Name:     wrapperspb.String(name),
				Email:    wrapperspb.String(email),
				Password: wrapperspb.String(password),
				Role:     role,
			},
		}

		info = &model.UserInfo{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     role,
		}

		res = &emptypb.Empty{}
	)

	t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
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
				mock := mocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, id, info).Return(nil)
				return mock
			},
		},
		{
			name: "service error",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := mocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, id, info).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userServiceMock := tt.userServiceMock(mc)
			api := user.NewUserImplementation(userServiceMock)

			empty, err := api.Update(tt.args.ctx, tt.args.req)
			if err != nil {
				return
			}

			require.Equal(t, tt.want, empty)
			require.Equal(t, tt.err, err)
		})
	}
}
