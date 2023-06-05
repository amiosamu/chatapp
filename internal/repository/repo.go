package repository

import (
	"context"
	"database/sql"
	"github.com/amiosamu/chatapp/entity"
)

type DBTX interface {
	PrepareContext(ctx context.Context, query string, args ...interface{})
	ExecContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type User interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserByName(ctx context.Context, name string) (*entity.User, error)
}

type Repositories struct {
	DBTX
	User
}
