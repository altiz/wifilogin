package irbis

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/godror/godror"
)

type id struct {
	id int
}

type login struct {
	login  string
	passwd string
}

func ConnectDB(dbname string) (*sql.DB, error) {

	db, err := sql.Open("godror", `user="wifiservice" password="wifi" connectString="e-scan:1521/irbis"  poolSessionMaxLifetime=24h poolSessionTimeout=30s`)
	if err != nil {
		return nil, err
	}

	return db, errors.New(fmt.Sprint("OK"))
}
func QueryDB()
rows, err := db.Query("select wifi_02_login_seq.nextval id from dual")
	if err != nil {
		panic(err)
	}
	id := id{}
	rows.Next()
	err1 := rows.Scan(&id.id)
	if err1 != nil {
		fmt.Println(err)
	}