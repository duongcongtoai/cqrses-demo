package db

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var (
	config = map[string]string{
		"host":     "",
		"password": "",
		"database": "",
		"port":     "",
		"username": "",
	}
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)

	}
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName("cqrs")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}

func NewMysqlConnection() (*sqlx.DB, error) {
	connString := config["username"] +
		":" + config["password"] +
		"@" + "(" + config["host"] +
		":" + config["port"] +
		")/" + config["database"]
	return sqlx.Connect("mysql", connString)
}
