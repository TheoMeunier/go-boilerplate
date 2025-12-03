package db

import (
	"context"
	"fmt"
	"log"
)

type BaseRepo struct {
	DB *Database
}

func NewBaseRepo(database *Database) *BaseRepo {
	return &BaseRepo{DB: database}
}

func (r *BaseRepo) Ping() error {
	ctx := context.Background()
	if err := r.DB.Conn.Ping(ctx); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}
	log.Println("Database connection OK")
	return nil
}
