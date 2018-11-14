package main

import (
	"database/sql"
	"fmt"
	"math/big"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	db, err := sql.Open("mysql", "engine:decopentask@/ot_local")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for reading data
	query := "SELECT id, tx, mission_id, reward FROM publish limit 5"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var (
			id      string
			tx      string
			mission string
			rs      sql.NullString
		)
		err := rows.Scan(&id, &tx, &mission, &rs)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %s, tx: %s, mission: %s, reward: %s\n", id, tx, mission, rs.String)
	}

	stmtOut, err := db.Prepare("SELECT reward FROM publish WHERE reward >= ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	r1, ok := new(big.Int).SetString("5", 10)

	if ok {
		_, err = stmtOut.Exec(r1.String())
		if err != nil {
			panic(err.Error())
		}
		var rs sql.NullString
		var reward *big.Int
		err = stmtOut.QueryRow(1).Scan(&rs) // WHERE number = 1
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		reward, _ = new(big.Int).SetString(rs.String, 10)
		fmt.Printf("The reward of 1st row is: %s\n", reward.String())
	}

}
