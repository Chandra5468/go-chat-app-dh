package users

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

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInserted int
	query := "insert into users(username, password, email) values ($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInserted)

	if err != nil {
		return &User{}, err
	}

	user.ID = int64(lastInserted)

	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}

	query := "select id, email, username, password from users where email = $1"

	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.Username, &u.Password)

	if err != nil {
		return &User{}, err
	}
	return &u, nil
}
