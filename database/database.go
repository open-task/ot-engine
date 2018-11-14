package database

import (
	"database/sql"
	"github.com/xyths/ot-engine/types"
	"log"
)

func Publish(db *sql.DB, e types.PublishEvent) (err error) {
	// 接受日志重复，并如实记录下来（下同）。
	stmtIns, err := db.Prepare("INSERT INTO publish (mission_id, reward, publisher, block, tx) VALUES(?, ?, ?, ?, ?)") // ? = placeholder
	if err != nil {
		log.Fatal(err)
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(e.Mission, e.Reward.String(), e.Publisher, e.Block, e.Tx)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func Solve(db *sql.DB, event types.PublishEvent) (err error) {
	return err
}

func Accept(db *sql.DB, event types.PublishEvent) (err error) {
	return err
}

func Reject(db *sql.DB, event types.PublishEvent) (err error) {
	return err
}

func Confirm(db *sql.DB, event types.PublishEvent) (err error) {
	return err
}
