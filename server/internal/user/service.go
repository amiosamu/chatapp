package user

import (
	"context"
	"github.com/amiosamu/chatapp/utils"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func (s service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPass, err := utils.HashedPassword(req.Password)

	if err != nil {
		return nil, err
	}
	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPass,
	}
	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	res := &CreateUserRes{
		ID:       strconv.Itoa(int(r.Id)),
		Username: r.Username,
		Email:    r.Email,
	}
	return res, nil

}

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s service) Signin(c context.Context, req *SigninUserReq) (*SigninUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()
	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &SigninUserRes{}, err
	}
	err = utils.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &SigninUserRes{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(u.Id)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.Id)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})
	accToken, err := token.SignedString([]byte(os.Getenv("SECRETKEY")))
	if err != nil {
		return &SigninUserRes{}, err
	}
	return &SigninUserRes{accessToken: accToken, Username: u.Username, ID: strconv.Itoa(int(u.Id))}, nil
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}
