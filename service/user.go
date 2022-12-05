package service

import (
	"context"

	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/auth_service/genproto/auth_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/auth_service/storage"
	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/auth_service/util"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authService struct {
	auth_service.UnimplementedAuthServiceServer
	stg storage.StorageI
}

func NewAuthService(db *sqlx.DB) *authService {
	return &authService{
		stg: storage.NewStoragePg(db),
	}
}

func (s *authService) CreateUser(ctx context.Context, req *auth_service.CreateUserRequest) (*auth_service.CreateUserResponse, error) {
	hashedPassword,err := util.HashPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "util.HashPassword() : %v", err)
	}
	req.Password=hashedPassword
	id, err := s.stg.User().CreateUser(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "method CreateUser: %v", err)

	}
	return &auth_service.CreateUserResponse{Id: id}, nil
}

func (s *authService) GetUserList(ctx context.Context, req *auth_service.GetUserListRequest) (*auth_service.GetUserListResponse, error) {
	res, err := s.stg.User().GetUserList(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "method GetUser: %v", err)

	}
	return res, nil
}

func (s *authService) GetUserById(ctx context.Context, req *auth_service.GetUserByIdRequest) (*auth_service.User, error) {
	res, err := s.stg.User().GetUserById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "method GetUserById: %v", err)

	}
	return res, nil
}

func (s *authService) UpdateUser(ctx context.Context, req *auth_service.UpdateUserRequest) (*auth_service.UpdateUserResponse, error) {
	hashedPassword,err := util.HashPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "util.HashPassword() : %v", err)
	}
	req.Password=hashedPassword
	err = s.stg.User().UpdateUser(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "method UpdateUser: %v", err)

	}

	return &auth_service.UpdateUserResponse{Status: "Updated"}, nil
}
func (s *authService) DeleteUser(ctx context.Context, req *auth_service.DeleteUserRequest) (*auth_service.DeleteUserResponse, error) {
	err := s.stg.User().DeleteUser(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "method DeleteUser: %v", err)

	}
	return &auth_service.DeleteUserResponse{Status: "Deleted"}, nil
}
