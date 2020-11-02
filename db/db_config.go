package db

type DBConfig struct {
	Username string
	Password string
	Host string
	Port string
	DbName string
	Sslmode string
	MaxConn int
}