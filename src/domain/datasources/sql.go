package datasources

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type SqlDb struct {
	Connect *sql.DB
}

func NewSqlDb() *SqlDb {
	db, err := sql.Open("mysql", os.Getenv("MYSQL"))
	if err != nil {
		log.Fatal("error connection : ", err)
	}

	return &SqlDb{Connect: db}

}
