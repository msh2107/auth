package user

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/msh2107/auth/pkg/user_v1"
)

// Delete - .
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.userService.Delete(ctx, req.GetId())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("deleted user with id: %d \n", req.GetId())
	return &emptypb.Empty{}, nil
}
