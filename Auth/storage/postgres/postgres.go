package postgres

import (
	"database/sql"
	"fmt"
	"auth/config"
	stg "auth/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db    *sql.DB
	users stg.UsersService
}

func NewPostgresStorage() (stg.InitRoot, error) {
	config := config.Load()
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.PostgresUser, config.PostgresPassword,
		config.PostgresHost, config.PostgresPort,
		config.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Storage{Db: db}, nil
}

func (s *Storage) Users() stg.UsersService {
	if s.users == nil {
		s.users = &UserStorage{db: s.Db}
	}
	return s.users
}