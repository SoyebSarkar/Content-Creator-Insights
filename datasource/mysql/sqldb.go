package mysql

import (
	"database/sql"
	"fmt"

	"github.com/SoyebSarkar/content-creator-insight/datasource/config"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	databaseDetails := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		config.MysqlUser,
		config.MySqlPassword,
		config.MysqlHost,
		config.MySqlTable,
	)
	var err error
	Client, err = sql.Open("mysql", databaseDetails)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("MySql database configured")

}
