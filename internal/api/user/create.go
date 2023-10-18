package user

import (
	"context"
	"log"

	"github.com/msh2107/auth/internal/converter"
	desc "github.com/msh2107/auth/pkg/user_v1"
)

// Create - .
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToUserInfoFromDesc(req.GetInfo()))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Printf("created user with id: %d \n", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
