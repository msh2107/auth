package v1

import (
	"context"
	"fmt"
	"log"
	"strings"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/msh2107/auth/internal/models"
	"github.com/msh2107/auth/internal/repository"

	desc "github.com/msh2107/auth/pkg/user_v1"
)

// UserServer -
type UserServer struct {
	desc.UnimplementedUserV1Server
	repository repository.Repository
}

// NewUserServer - .
func NewUserServer(repository repository.Repository) *UserServer {
	return &UserServer{
		repository: repository,
	}
}

// Create - .
func (s *UserServer) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	reqStr := fmt.Sprintf("Received Create:\n\tName: %v,\n\tEmail: %v,\n\tPassword: %v,\n\tPassword confirm: %v,\n\tRole: %v\n",
		req.GetName(),
		req.GetEmail(),
		req.GetPassword(),
		req.GetPasswordConfirm(),
		req.GetRole(),
	)

	log.Println(reqStr)

	user := models.User{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     int32(req.GetRole()),
	}
	id, err := s.repository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

// Get - .
func (s *UserServer) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Received Get:\n\tId: %v\n", req.GetId())

	user, err := s.repository.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	resp := desc.GetResponse{
		Id:        req.GetId(),
		Name:      user.Name,
		Email:     user.Email,
		Role:      desc.Role(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
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
func (s *UserServer) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {

	buf := strings.Builder{}
	buf.WriteString("Received Update:\n")
	idStr := fmt.Sprintf("\tId: %v\n", req.GetId())
	var user models.User
	buf.WriteString(idStr)
	user.ID = req.GetId()

	if req.Name != nil {
		buf.WriteString(fmt.Sprintf("\tName: %v\n", req.GetName().GetValue()))
		user.Name = req.GetName().GetValue()
	}

	if req.Email != nil {
		buf.WriteString(fmt.Sprintf("\tEmail: %v\n", req.GetEmail().GetValue()))
		user.Email = req.GetEmail().GetValue()
	}

	if req.Role != desc.Role_UNDEFINED {
		buf.WriteString(fmt.Sprintf("\tRole: %v\n", req.GetRole()))
		user.Role = int32(req.GetRole())
	}

	err := s.repository.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	log.Println(buf.String())
	return &emptypb.Empty{}, nil
}

// Delete - .
func (s *UserServer) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Received Delete:\n\tId: %v", req.GetId())
	err := s.repository.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
