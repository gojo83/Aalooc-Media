package util

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

type Conn struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func GetConnection(option Conn) *sqlx.DB {
	if Db != nil {
		fmt.Println("Same instance")
		return Db
	} else {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			option.Host, option.Port, option.User, option.Password, option.Dbname)
		Db, err := sqlx.Open("postgres", psqlInfo)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return Db
	}
}
