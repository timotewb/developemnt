package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type serverType struct {
	Name string `json:"name"`
	Mac  string `json:"mac"`
	IP   string `json:"ip"`
}

func main() {
	const file string = "sqlite3.db"
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	rows, err := db.Query("select name, ip, mac from server_details")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var data []serverType
	for rows.Next() {
		var details serverType
		err = rows.Scan(&details.Name, &details.IP, &details.Mac)
		if err != nil {
			fmt.Println(err)
		}
		data = append(data, details)
	}
	fmt.Print(data)
 }

 // https://zetcode.com/golang/sqlite3/