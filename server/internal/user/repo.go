package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r repository) CreateUser(ctx context.Context, u *User) (*User, error) {
	var insertedID int
	query := `INSERT INTO users(username, password,email) VALUES ($1, $2, $3) returning id`
	err := r.db.QueryRowContext(ctx, query, u.Username, u.Password, u.Email).Scan(&insertedID)
	if err != nil {
		return &User{}, err
	}
	u.Id = int64(insertedID)
	return u, nil
}

func (r repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	user := User{}

	query := `SELECT id, email,username, password FROM users WHERE email = $1`
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.Id, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return &User{}, nil
	}
	return &user, nil

}
