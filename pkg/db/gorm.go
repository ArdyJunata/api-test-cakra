package db

import (
	"fmt"
	"time"

	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm interface {
	Connect() error
	SetConnectionPool(maxOpen, maxIdle int, idleTime, lifeTime time.Duration) error
}

type GormPostgresDB struct {
	host   string
	port   string
	user   string
	pass   string
	dbname string
	ssl    string
	DB     *gorm.DB
}

func NewGormPostgres(host, port, user, pass, dbname, ssl string) Gorm {
	return &GormPostgresDB{
		host:   host,
		port:   port,
		user:   user,
		pass:   pass,
		dbname: dbname,
		ssl:    ssl,
	}
}

func (g *GormPostgresDB) Connect() error {
	fmt.Println(g.host)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		g.host, g.port, g.user, g.pass, g.dbname, g.ssl,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "nrpostgres",
		DSN:        dsn,
	}))
	if err != nil {
		return err
	}
	g.DB = db

	return nil
}

func (g *GormPostgresDB) SetConnectionPool(maxOpen, maxIdle int, idleTime, lifeTime time.Duration) error {
	myDB, err := g.DB.DB()
	if err != nil {
		return err
	}

	myDB.SetMaxOpenConns(maxOpen)
	myDB.SetMaxIdleConns(maxIdle)
	myDB.SetConnMaxIdleTime(idleTime * time.Second)
	myDB.SetConnMaxLifetime(lifeTime * time.Second)
	return nil
}
