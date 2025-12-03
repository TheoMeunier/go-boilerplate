package db

import (
	"context"
	"fmt"
	"log"

	"github.com/go-pg/pg/v10"
)

type Database struct {
	Conn *pg.DB
}

func NewDatabase(user, password, host, port, name string) (*Database, error) {
	opt := &pg.Options{
		User:     user,
		Password: password,
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Database: name,
	}

	db := pg.Connect(opt)

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Connected to database")
	return &Database{Conn: db}, nil
}

func (d *Database) Close() error {
	return d.Conn.Close()
}
