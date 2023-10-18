package user

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/msh2107/auth/internal/converter"
	desc "github.com/msh2107/auth/pkg/user_v1"
)

// Update - .
func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := i.userService.Update(ctx, req.GetId(), converter.ToUserFromUpdateDesc(req.Info))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Printf("updated user with id: %d \n", req.GetId())

	return &emptypb.Empty{}, nil
}
