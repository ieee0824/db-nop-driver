package main

import (
	"database/sql"
	"fmt"

	_ "github.com/ieee0824/db-nop-driver"
)

type obj struct{}

func main() {
	db, err := sql.Open("dbnop", "nop")
	if err != nil {
		panic(err)
	}

	row := db.QueryRow("SELECT 1")

	ret := []obj{}
	if err := row.Scan(&ret); err != nil {
		panic(err)
	}
	fmt.Println(ret)
}
