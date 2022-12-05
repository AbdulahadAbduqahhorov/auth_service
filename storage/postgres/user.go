package postgres

import (
	"database/sql"
	"errors"

	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/auth_service/genproto/auth_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/auth_service/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) repo.UserRepoI {
	return userRepo{
		db: db,
	}
}
func (u userRepo) CreateUser(req *auth_service.CreateUserRequest) (string, error) {
	id := uuid.New().String()
	_, err := u.db.Exec(`INSERT INTO 
		"user" (
			id,
			username,
			password,
			user_type
			) 
		VALUES (
			$1, 
			$2,
			$3,
			$4
			)`,
		id,
		req.Username,
		req.Password,
		req.UserType.String(),  //string()
	)
	if err != nil {
		return "", err
	}
	return id, nil

}

func (u userRepo) GetUserList(req *auth_service.GetUserListRequest) (*auth_service.GetUserListResponse, error) {
	res := &auth_service.GetUserListResponse{
		Users: make([]*auth_service.User, 0),
	}
	
	rows, err := u.db.Queryx(`SELECT 
		id,
		username,
		password,
		user_type,
		created_at,
		updated_at 
		FROM "user"
		WHERE (username ILIKE '%' || $1 || '%') 
		LIMIT $2
		OFFSET $3
	`,
		req.Search,
		int(req.Limit),
		int(req.Offset),
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			user auth_service.User

			updatedAt sql.NullString
			userType string
		)
		 
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Password,
			&userType,
			&user.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		if updatedAt.Valid {
			user.UpdatedAt = updatedAt.String
		}
		user.UserType = auth_service.UserT(auth_service.UserT_value[userType])
		
		res.Users = append(res.Users, &user)

	}

	return res, nil

}

func (u userRepo) GetUserById(id string) (*auth_service.User, error) {
	res := &auth_service.User{}
	var (
		updatedAt sql.NullString
		userType string
	)
	err := u.db.QueryRow(`
	SELECT 
		id,
		username,
		password,
		user_type,
		created_at,
		updated_at 
	FROM "user"
	WHERE id=$1`, id).Scan(
		&res.Id,
		&res.Username,
		&res.Password,
		&userType,
		&res.CreatedAt,
		&updatedAt,
	)
	if updatedAt.Valid {
		res.UpdatedAt = updatedAt.String
	}
	res.UserType = auth_service.UserT(auth_service.UserT_value[userType])

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u userRepo) UpdateUser(req *auth_service.UpdateUserRequest) error {
	res, err := u.db.NamedExec(`
	UPDATE  "user" SET 
		password=:p, 
		updated_at=now() 
		WHERE id=:i `, map[string]interface{}{
		"p": req.Password,
		"i": req.Id,
	})
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n > 0 {
		return nil
	}
	return errors.New("author not found")
}

func (u userRepo) DeleteUser(id string) error {

	res, err := u.db.Exec(`DELETE FROM "user" WHERE id=$1`, id)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n >0 {
		return nil
	}
	return errors.New("author not found")
}

func (u userRepo) GetUserByUsername(username string) (*auth_service.User, error) {
	res := &auth_service.User{}
	var (
		updatedAt sql.NullString
		userType string
	)
	err := u.db.QueryRow(`
	SELECT 
		id,
		username,
		password,
		user_type,
		created_at,
		updated_at 
	FROM "user"
	WHERE username=$1`, username).Scan(
		&res.Id,
		&res.Username,
		&res.Password,
		&userType,
		&res.CreatedAt,
		&updatedAt,
	)
	if updatedAt.Valid {
		res.UpdatedAt = updatedAt.String
	}
	res.UserType = auth_service.UserT(auth_service.UserT_value[userType])

	if err != nil {
		return nil, err
	}

	return res, nil
}