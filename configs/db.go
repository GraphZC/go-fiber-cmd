package configs

import (
	_ "database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func ConnectDB() error {
	var err error
	Db, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", Cfg.DBUser, Cfg.DBPass, Cfg.DBHost, Cfg.DBPort, Cfg.DBName))
	if err != nil {
		return err
	}
	
	err = Db.Ping()
	if err != nil {
		return err
	}

	return nil
}