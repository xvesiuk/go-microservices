package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xvesiuk/go-microservices/svc-posts/internal/config"
)

func NewPool(conf *config.DatabaseConfig) (*pgxpool.Pool, error) {
	connPool, err := pgxpool.NewWithConfig(context.Background(), configure(conf))
	if err != nil {
		log.Println("Error while creating connection to the database!!")
		return nil, err
	}

	connection, err := connPool.Acquire(context.Background())
	if err != nil {
		log.Println("Error while acquiring connection from the database pool!")
		return nil, err
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		log.Println("Could not ping database")
		return nil, err
	}

	fmt.Println("Connected to the database!")
	return connPool, nil
}

func configure(conf *config.DatabaseConfig) *pgxpool.Config {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)

	log.Println(url)

	dbConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = conf.MaxConns
	dbConfig.MinConns = conf.MinConns
	dbConfig.MaxConnLifetime = conf.MaxConnLifetime
	dbConfig.MaxConnIdleTime = conf.MaxConnIdleTime
	dbConfig.HealthCheckPeriod = conf.HealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = conf.ConnectTimeout

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		log.Println("Before acquiring the connection pool to the database!!")
		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("After releasing the connection pool to the database!!")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		log.Println("Closed the connection pool to the database!!")
	}

	return dbConfig
}
