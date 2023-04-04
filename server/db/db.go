package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewDB(cfg Config) (*Database, error) {
	dsn, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	//err = dsn.Ping()
	//if err != nil {
	//	return nil, err
	//}
	return &Database{db: dsn}, nil
}

func (d *Database) Close() {
	err := d.db.Close()
	if err != nil {
		return
	}
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
