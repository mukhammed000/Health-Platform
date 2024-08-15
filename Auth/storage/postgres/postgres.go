package postgres

import (
	"auth/config"
	stg "auth/storage"
	"database/sql"
	"fmt"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db    *sql.DB
	rdb   *redis.Client
	users stg.UsersService
}

func NewPostgresStorage() (stg.InitRoot, error) {
	cfg := config.Load()
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser, cfg.PostgresPassword,
		cfg.PostgresHost, cfg.PostgresPort,
		cfg.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
	})
	

	return &Storage{Db: db, rdb: rdb}, nil
}

func (s *Storage) Users() stg.UsersService {
	if s.users == nil {
		s.users = &UserStorage{db: s.Db, rdb: s.rdb}
	}
	return s.users
}
