package database

import (
	"database/sql"
	"github.com/xyths/ot-engine/types"
)

func Publish(db *sql.DB, e types.PublishEvent) {
	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO publish (mission_id, reward, publisher, block, tx) VALUES(?, ?, ?, ?, ?)") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	_, err = stmtIns.Exec(e.Mission, e.Reward.String(), e.Publisher, e.Block, e.Tx) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func Solve(db *sql.DB, event types.PublishEvent) {
}

func Accept(db *sql.DB, event types.PublishEvent) {
}

func Reject(db *sql.DB, event types.PublishEvent) {
}

func Confirm(db *sql.DB, event types.PublishEvent) {
}
