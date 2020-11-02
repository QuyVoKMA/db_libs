package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	_"github.com/lib/pq"
)

var conn *sqlx.DB



func Init(config *DBConfig){
	connString :=fmt.Sprintf("%v %v %v %v %v %v",
		config.Host,
		config.Port,
		config.Username,
		config.DbName,
		config.Sslmode,
		config.Password,
	)
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Fatalln("Could not connect to DB")
	}
	db.SetMaxIdleConns(config.MaxConn)
	db.SetMaxOpenConns(config.MaxConn)
	conn =db
}

func GetConn() *sqlx.DB{
	return conn
}
func SetConn(db *sqlx.DB){
	conn =db
}
