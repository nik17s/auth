package postgres

import (
	"context"
	entity "github.com/nik17s/auth"
)

func (r *Repository) CreateUser(login, password, email, phone string) (int, error) {
	createUserQuery := "INSERT INTO users (login, password, email, phone) VALUES ($1, $2, $3, $4) RETURNING id"

	row := r.pool.QueryRow(context.Background(), createUserQuery, login, password, email, phone)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) GetUser(login, password string) (*entity.User, error) {
	getUserQuery := "SELECT * FROM users WHERE login=$1 AND password=$2"

	row := r.pool.QueryRow(context.Background(), getUserQuery, login, password)

	var user entity.User
	if err := row.Scan(&user.Id, &user.Login, &user.Password, &user.Email, &user.Phone); err != nil {
		return nil, err
	}

	return &user, nil
}
