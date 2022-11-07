package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Postgres interface {
	Connect() error
	SetConnectionPool(maxOpen, maxIdle int, idleTime, lifeTime time.Duration) error
}

type PostgresDB struct {
	host   string
	port   string
	user   string
	pass   string
	dbname string
	ssl    string
	DB     *sql.DB
}

func NewPostgresDB(host, port, user, pass, dbname, ssl string) Postgres {
	return &PostgresDB{
		host:   host,
		port:   port,
		user:   user,
		pass:   pass,
		dbname: dbname,
		ssl:    ssl,
	}
}

func (p *PostgresDB) Connect() error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		p.host, p.port, p.user, p.pass, p.dbname, p.ssl,
	)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	p.DB = db

	return nil

}

func (p *PostgresDB) SetConnectionPool(maxOpen, maxIdle int, idleTime, lifeTime time.Duration) error {
	if p.DB == nil {
		return fmt.Errorf("db is nil")
	}

	p.DB.SetMaxOpenConns(maxOpen)
	p.DB.SetMaxIdleConns(maxIdle)
	p.DB.SetConnMaxIdleTime(idleTime)
	p.DB.SetConnMaxLifetime(lifeTime)

	return nil
}
