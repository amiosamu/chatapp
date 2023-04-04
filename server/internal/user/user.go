package user

import "context"

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type SigninUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninUserRes struct {
	accessToken string
	ID          string `json:"id"`
	Username    string `json:"username"`
}

type Repository interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type Service interface {
	CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error)
	Signin(c context.Context, req *SigninUserReq) (*SigninUserRes, error)
}
