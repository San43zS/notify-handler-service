package psql

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	URL    string
	Driver string
}

type dbParams struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func getDBParams() *dbParams {
	return &dbParams{
		host:     viper.GetString("DB.HOST"),
		port:     viper.GetString("DB.PORT"),
		user:     viper.GetString("DB.USER"),
		password: viper.GetString("DB.PASSWORD"),
		dbname:   viper.GetString("DB.DBNAME"),
	}
}

func (db dbParams) ParseURL() string {
	template := viper.GetString("DB.URLTEMPLATE")

	return fmt.Sprintf(template, db.host, db.port, db.dbname, db.user, db.password)
}

func NewConfig() *Config {
	test := getDBParams()

	return &Config{
		URL:    test.ParseURL(),
		Driver: viper.GetString("DB.DRIVER"),
	}
}

func GetUniqueViolationErr() string {
	return viper.GetString("ERR.USER_ALREADY_EXISTS")
}
