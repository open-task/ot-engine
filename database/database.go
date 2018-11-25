package database

import (
	"database/sql"
	"github.com/xyths/ot-engine/types"
	"log"
)

func Publish(db *sql.DB, e types.PublishEvent) (err error) {
	// 接受日志重复，并如实记录下来（下同）。
	stmtIns, err := db.Prepare("INSERT INTO publish (mission_id, reward, publisher, block, tx) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(e.Mission, e.Reward.String(), e.Publisher, e.Block, e.Tx)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func Solve(db *sql.DB, e types.SolveEvent) (err error) {
	stmtIns, err := db.Prepare("INSERT INTO solve (solution_id, mission_id, context, solver, block, tx) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(e.Solution, e.Mission, e.Data, e.Solver, e.Block, e.Tx)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func Accept(db *sql.DB, e types.AcceptEvent) (err error) {
	stmtIns, err := db.Prepare("INSERT INTO accept (solution_id, block, tx) VALUES(?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(e.Solution, e.Block, e.Tx)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func Reject(db *sql.DB, e types.RejectEvent) (err error) {
	stmtIns, err := db.Prepare("INSERT INTO reject (solution_id, block, tx) VALUES(?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(e.Solution, e.Block, e.Tx)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func Confirm(db *sql.DB, e types.ConfirmEvent) (err error) {
	stmtIns, err := db.Prepare("INSERT INTO confirm (solution_id, arbitration_id, block, tx) VALUES(?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(e.Solution, e.Arbitration, e.Block, e.Tx)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func GetPublished(db *sql.DB, address string, limit string) (events []types.PublishEvent, err error) {
	stmt, err := db.Prepare("SELECT mission_id, reward, txtime FROM publish WHERE publisher = ? LIMIT ?")
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(address, limit)
	if err != nil {
		log.Println(err)
		return
	}
	for rows.Next() {
		var p types.PublishEvent
		var txTime int
		err = rows.Scan(&p.Mission, &p.Reward, &txTime)
		if err != nil {
			log.Println(err)
			continue
		}
		events = append(events, p)
		_ = txTime
	}
	return events, err
}

func GetSolved(db *sql.DB, address string, limit string) (events []types.PublishEvent, err error) {

	return events, err
}

func GetAccepted(db *sql.DB, address string, limit string) (events []types.PublishEvent, err error) {

	return events, err
}

func GetRejected(db *sql.DB, address string, limit string) (events []types.PublishEvent, err error) {

	return events, err
}
