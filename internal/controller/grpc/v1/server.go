package v1

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	desc "github.com/msh2107/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/rand"
	"strings"
)

type AuthServer struct {
	desc.UnimplementedUserV1Server
}

func NewAuthServer() *AuthServer {
	return &AuthServer{}
}

func (s *AuthServer) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	reqStr := fmt.Sprintf("Received Create:\n\tName: %v,\n\tEmail: %v,\n\tPassword: %v,\n\tPassword confirm: %v,\n\tRole: %v\n",
		req.GetName(),
		req.GetEmail(),
		req.GetPassword(),
		req.GetPasswordConfirm(),
		req.GetRole(),
	)

	log.Println(reqStr)

	id := rand.Int63()

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
func (s *AuthServer) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Received Get:\n\tId: %v\n", req.GetId())

	role := gofakeit.RandString([]string{"ADMIN", "USER"})
	resp := desc.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      desc.Role(desc.Role_value[role]),
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}

	respStr := fmt.Sprintf("Response Get:\n\tId: %v,\n\tName: %v,\n\tEmail: %v,\n\tRole: %v,\n\tCreatedAt: %v,\n\tUpdatedAt: %v\n",
		resp.Id,
		resp.Name,
		resp.Email,
		resp.Role,
		resp.CreatedAt,
		resp.UpdatedAt,
	)

	log.Println(respStr)

	return &resp, nil
}
func (s *AuthServer) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {

	buf := strings.Builder{}
	buf.WriteString("Received Update:\n")
	idStr := fmt.Sprintf("\tId: %v\n", req.GetId())
	buf.WriteString(idStr)

	if req.Name != nil {
		buf.WriteString(fmt.Sprintf("\tName: %v\n", req.GetName().GetValue()))
	}

	if req.Email != nil {
		buf.WriteString(fmt.Sprintf("\tEmail: %v\n", req.GetEmail().GetValue()))
	}

	if req.Role != desc.Role_UNDEFINED {
		buf.WriteString(fmt.Sprintf("\tRole: %v\n", req.GetRole()))
	}

	log.Println(buf.String())
	return &emptypb.Empty{}, nil
}
func (s *AuthServer) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Received Delete:\n\tId: %v", req.GetId())
	return &emptypb.Empty{}, nil
}
