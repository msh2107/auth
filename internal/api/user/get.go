package user

import (
	"context"
	"fmt"
	"log"

	"github.com/msh2107/auth/internal/converter"
	desc "github.com/msh2107/auth/pkg/user_v1"
)

// Get - .
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	userObj, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	respStr := fmt.Sprintf("Id: %v, Name: %v, Email: %v, Role: %v, CreatedAt: %v, UpdatedAt: %v\n",
		userObj.ID,
		userObj.Info.Name,
		userObj.Info.Email,
		userObj.Info.Role,
		userObj.CreatedAt,
		userObj.UpdatedAt,
	)

	log.Println(respStr)

	return &desc.GetResponse{
		User: converter.ToUserFromService(userObj),
	}, nil
}
