package db

import "testing"

const (
	host   = "localhost"
	port   = "8888"
	user   = "nbid"
	pass   = "nbid"
	dbname = "nbid"
	ssl    = "disable"
)

func TestConnectDB(t *testing.T) {
	db := NewPostgresDB(host, port, user, pass, dbname, ssl)
	err := db.Connect()

	if err != nil {
		t.Fatalf("cannt connect to db with error %s", err.Error())
	}
}
