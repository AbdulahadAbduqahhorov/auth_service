package repo

import (
	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/auth_service/genproto/auth_service"
)



type UserRepoI interface {
	CreateUser(req *auth_service.CreateUserRequest) (string, error)
	GetUserList(req *auth_service.GetUserListRequest) (*auth_service.GetUserListResponse, error)
	GetUserById(id string) (*auth_service.User, error)
	UpdateUser(req *auth_service.UpdateUserRequest) error
	DeleteUser(id string) error
	GetUserByUsername(username string) (*auth_service.User, error)
}
