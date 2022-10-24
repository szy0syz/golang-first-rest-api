package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ConfingDB struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func Connect(cnf ConfingDB) (*sqlx.DB, error) {
	// 好吧，得习惯这样的字符串拼接
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cnf.Host,
		cnf.Port,
		cnf.User,
		cnf.Password,
		cnf.Name,
	)
	db, err := sqlx.Connect("postgres", dsn)
	return db, err
}
