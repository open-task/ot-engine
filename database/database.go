package database

import (
	"database/sql"
	. "github.com/xyths/ot-engine/types"
	"log"
	"errors"
	"strings"
	"fmt"
	"math/big"
)

func Publish(db *sql.DB, e PublishEvent) (err error) {
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

func Solve(db *sql.DB, e SolveEvent) (err error) {
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

func Accept(db *sql.DB, e AcceptEvent) (err error) {
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

func Reject(db *sql.DB, e RejectEvent) (err error) {
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

func Confirm(db *sql.DB, e ConfirmEvent) (err error) {
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

func GetPublished(db *sql.DB, address string, limit int) (events []PublishEvent, err error) {
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
		var p PublishEvent
		var rewardStr sql.NullString
		var txTimeStr sql.NullString
		err = rows.Scan(&p.Mission, &rewardStr, &txTimeStr)
		if err != nil {
			log.Println(err)
			continue
		}
		p.Reward, _ = new(big.Int).SetString(rewardStr.String, 10)
		events = append(events, p)
	}
	return events, err
}

func GetSolutions(db *sql.DB, missions []string) (solutions []Solution, ids []string, err error) {
	if len(missions) <= 0 {
		err = errors.New("no mission id")
		return
	}
	query := "SELECT mission_id, solution_id, context, solver FROM solve WHERE mission_id in ('"
	query += strings.Join(missions, "','")
	query += "');"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Printf("Database Error when retrive solve: %s", err.Error())
		return
	}
	for rows.Next() {
		var s Solution
		err1 := rows.Scan(&s.Mission, &s.Solution, &s.Data, &s.Solver)
		if err1 != nil {
			log.Println(err1)
			continue
		}
		solutions = append(solutions, s)
		ids = append(ids, s.Solution)
	}

	return
}

func getProcessed(db *sql.DB, solutions []string, status string) (process []Process, ids []string, err error) {
	if len(solutions) <= 0 {
		err = errors.New("no solution id")
		return
	}
	status = strings.ToLower(status)
	if status != "reject" && status != "accept" {
		err = errors.New("status SHOULD be 'accept' or 'reject'")
		return
	}
	query := "SELECT solution_id, txtime FROM "
	query += status
	query += " WHERE solution_id in ('"
	query += strings.Join(solutions, "','")
	query += "');"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Printf("Database Error when retrive %s: %s", status, err.Error())
		return
	}
	for rows.Next() {
		var p Process
		err1 := rows.Scan(&p.Solution, &p.Time)
		if err1 != nil {
			log.Println(err1)
			continue
		}
		process = append(process, p)
		ids = append(ids, p.Solution) // success ids
	}

	return
}

func GetProcess(db *sql.DB, solutions []string) (process []Process, ids []string, err error) {
	p1, l1, e1 := getProcessed(db, solutions, "accept")
	if e1 != nil {
		fmt.Println(e1)
		return
	}
	p2, l2, e2 := getProcessed(db, solutions, "reject")
	if e2 != nil {
		fmt.Println(e2)
		return
	}
	process = append(p1, p2...)
	ids = append(l1, l2...)
	return
}
