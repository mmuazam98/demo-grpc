package service

import (
	"context"
	"fmt"

	"github.com/mmuazam98/demo-grpc/user"
)

type CustomUserService struct {
	user.UnimplementedUserServiceServer
	Users  map[int32]*user.User
	NextID int32
}

// CreateUser implements user.UserServiceServer.
func (cus *CustomUserService) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	cus.NextID++

	u := &user.User{
		Id:      cus.NextID,
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	}

	cus.Users[cus.NextID] = u

	return &user.CreateUserResponse{
		User: u,
	}, nil
}

// DeleteUser implements user.UserServiceServer.
func (cus *CustomUserService) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	if _, ok := cus.Users[req.Id]; !ok {
		return nil, fmt.Errorf("user with ID %d not found", req.Id)
	}

	delete(cus.Users, req.Id)

	return &user.DeleteUserResponse{
		Success: true,
	}, nil
}

// GetUser implements user.UserServiceServer.
func (cus *CustomUserService) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	existingUser, ok := cus.Users[req.Id]
	if !ok {
		return &user.GetUserResponse{}, fmt.Errorf("user with ID %d not found", req.Id)
	}

	return &user.GetUserResponse{
		User: existingUser,
	}, nil
}

// UpdateUser implements user.UserServiceServer.
func (cus *CustomUserService) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	existingUser, ok := cus.Users[req.Id]
	if !ok {
		return nil, fmt.Errorf("user with ID %d not found", req.Id)
	}

	if req.Name != "" {
		existingUser.Name = req.Name
	}
	if req.Email != "" {
		existingUser.Email = req.Email
	}
	if req.Phone != "" {
		existingUser.Phone = req.Phone
	}
	if req.Address != "" {
		existingUser.Address = req.Address
	}

	cus.Users[req.Id] = existingUser

	return &user.UpdateUserResponse{
		User: existingUser,
	}, nil
}
