package banco

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connction() (*sql.DB, error) {
	stringConection := fmt.Sprintf("%s:%s@/%s?charset=ut8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	db, err := sql.Open("mysql", stringConection)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
	}
	return db, nil
}
