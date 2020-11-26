package handlers

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

type TTime struct {
	Sysdate sql.NullTime `db:"S"`
}

func test() {
	db, err := sql.Open("godror", `user="ttk_billing" password="wdbip" connectString="e-scan:1521/irbis  poolSessionMaxLifetime=24h poolSessionTimeout=30s"`)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("OK")
	rows, err := db.Query("select sysdate S from dual")
	if err != nil {

		panic(err)
	}
	fmt.Println("OK")
	var thedate string
	for rows.Next() {
		rows.Scan(&thedate)
	}
	fmt.Println("OK")
	defer db.Close()
	fmt.Println(thedate)
	/*
		ei := TTime{}
		eil := []TTime{}
		for rows.Next() {

			if err != nil {
				log.Fatal(err)
			}
			eil = append(eil, ei)
		}*/
	//defer rows.Close()
}
