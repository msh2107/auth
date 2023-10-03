package v1

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/msh2107/auth/pkg/user_v1"
)

// AuthServer -
type AuthServer struct {
	desc.UnimplementedUserV1Server
}

// NewAuthServer - .
func NewAuthServer() *AuthServer {
	return &AuthServer{}
}

// Create - .
func (s *AuthServer) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	reqStr := fmt.Sprintf("Received Create:\n\tName: %v,\n\tEmail: %v,\n\tPassword: %v,\n\tPassword confirm: %v,\n\tRole: %v\n",
		req.GetName(),
		req.GetEmail(),
		req.GetPassword(),
		req.GetPasswordConfirm(),
		req.GetRole(),
	)

	log.Println(reqStr)

	randInt64, err := rand.Int(rand.Reader, new(big.Int).SetInt64(1<<62))
	if err != nil {
		return nil, err
	}

	id := randInt64.Int64()

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

// Get - .
func (s *AuthServer) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
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

// Update - .
func (s *AuthServer) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {

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

// Delete - .
func (s *AuthServer) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Received Delete:\n\tId: %v", req.GetId())
	return &emptypb.Empty{}, nil
}
