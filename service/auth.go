package service

import (
	"context"
	"time"

	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/auth_service/genproto/auth_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/auth_service/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authService) Login(ctx context.Context, req *auth_service.LoginRequest) (*auth_service.TokenResponse, error) {

	user, err := s.stg.User().GetUserByUsername(req.Username)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "username or password is wrong")
	}
	match, err := util.ComparePassword(user.Password, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "method ComparePassword: %v", err)

	}
	if !match {
		return nil, status.Errorf(codes.Unauthenticated, "username or password is wrong")
	}	
	m:=map[string]interface{}{
		"user_id":user.Id,
		"username":user.Username,
	}
	tokenStr,err:=util.GenerateJWT(m,time.Minute*10,"MySecretKey")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "method GenerateJWT: %v", err)

	}
	return &auth_service.TokenResponse{Token: tokenStr}, nil
}


