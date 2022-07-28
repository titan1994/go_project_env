package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

func SomeDBTest(url string) {
	dbPool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	var greeting string
	err = dbPool.QueryRow(context.Background(), "select 'TEST OK'").Scan(&greeting)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}

func GetPool(url string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.Connect(context.Background(), url)
	return dbPool, err
}
