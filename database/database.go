package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	. "github.com/open-task/ot-engine/types"
	"log"
	"math/big"
	"strings"
)

var Decimals = big.NewFloat(1e+18)

func Publish(ctx context.Context, db *sql.DB, e PublishEvent) (err error) {
	// 接受日志重复，并如实记录下来（下同）。
	stmtIns, err := db.PrepareContext(ctx, `INSERT INTO mission (
mission_id, reward, context, publisher, block, tx, txtime
) VALUES( 
?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(e.Mission, e.Reward.String(), e.Data, e.Publisher, e.Block, e.Tx, e.TxTime)
	if err != nil {
		return err
	}
	return err
}

func Solve(ctx context.Context, db *sql.DB, e SolveEvent) (err error) {
	// 如果mission_id/solution_id重复，会导致错误关联（后期通过合约解决）
	stmtIns, err := db.PrepareContext(ctx, `INSERT INTO solution (
solution_id, mission_id, context, solver, block, tx, txtime
) VALUES(
?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(e.Solution, e.Mission, e.Data, e.Solver, e.Block, e.Tx, e.TxTime)
	if err != nil {
		log.Println(err)
		return err
	}

	err = addSolutionNum(ctx, db, e.Mission)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func Accept(ctx context.Context, db *sql.DB, e AcceptEvent) (err error) {
	stmtIns, err := db.PrepareContext(ctx, `INSERT INTO accept (
solution_id, block, tx, txtime
) VALUES(
?, ?, ?, ?)`)
	if err != nil {
		//log.Println(err)
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(e.Solution, e.Block, e.Tx, e.TxTime)
	if err != nil {
		//log.Println(err)
		return err
	}

	err = updateSolved(ctx, db, e.Solution)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

func Reject(ctx context.Context, db *sql.DB, e RejectEvent) (err error) {
	stmtIns, err := db.PrepareContext(ctx, "INSERT INTO reject (solution_id, block, tx, txtime) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(e.Solution, e.Block, e.Tx, e.TxTime)
	if err != nil {
		log.Println(err)
		return err
	}

	// TODO: change mission status
	return err
}

func Confirm(ctx context.Context, db *sql.DB, e ConfirmEvent) (err error) {
	stmtIns, err := db.PrepareContext(ctx, `INSERT INTO confirm (
solution_id, arbitration_id, block, tx, txtime
) VALUES(
?, ?, ?, ?)`)
	if err != nil {
		//log.Println(err)
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(e.Solution, e.Arbitration, e.Block, e.Tx, e.TxTime)
	if err != nil {
		//log.Println(err)
		return err
	}

	return err
}

func GetAllPublished(ctx context.Context, db *sql.DB, offset int, limit int) (events []PublishEvent, err error) {
	log.Printf("database.GetAllPublished called: offset = %d, limit = %d\n", offset, limit)
	stmt, err := db.PrepareContext(ctx, `SELECT
block, tx, mission_id, reward, context, publisher, solution_num, solved, txtime
FROM mission
ORDER BY block DESC
LIMIT ?, ?`)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, offset, limit)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	log.Println("get rows")
	for rows.Next() {
		var p PublishEvent
		var solved bool
		var rewardStr sql.NullString
		err = rows.Scan(&p.Block, &p.Tx, &p.Mission, &rewardStr, &p.Data, &p.Publisher, &p.SolutionNumber, &solved, &p.TxTime)
		if err != nil {
			log.Println(err)
			continue
		}
		p.UpdateStatus(solved)
		var success bool
		p.Reward, success = big.NewInt(0).SetString(rewardStr.String, 10)
		if !success {
			p.Reward = big.NewInt(0)
		}
		p.RewardInDET, success = big.NewFloat(0).SetString(rewardStr.String)
		if !success {
			p.RewardInDET = big.NewFloat(0)
		}
		p.RewardInDET.Quo(p.RewardInDET, Decimals)
		events = append(events, p)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return events, err
}

func GetPublished(ctx context.Context, db *sql.DB, address string, limit int) (events []PublishEvent, err error) {
	stmt, err := db.PrepareContext(ctx, `SELECT
block, tx, mission_id, reward, publisher, solution_num, solved, txtime
FROM mission
WHERE publisher = ?
ORDER BY block DESC
LIMIT ?`)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, address, limit)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var p PublishEvent
		var solved bool
		var rewardStr sql.NullString
		err = rows.Scan(&p.Block, &p.Tx, &p.Mission, &rewardStr, &p.Publisher, &p.SolutionNumber, &solved, &p.TxTime)
		if err != nil {
			log.Println(err)
			continue
		}
		p.UpdateStatus(solved)
		var success bool
		p.Reward, success = big.NewInt(0).SetString(rewardStr.String, 10)
		if !success {
			p.Reward = big.NewInt(0)
		}
		p.RewardInDET, success = big.NewFloat(0).SetString(rewardStr.String)
		if !success {
			p.RewardInDET = big.NewFloat(0)
		}
		p.RewardInDET.Quo(p.RewardInDET, Decimals)
		events = append(events, p)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return events, err
}

func GetUnsolved(ctx context.Context, db *sql.DB, offset int, limit int) (events []PublishEvent, err error) {
	stmt, err := db.PrepareContext(ctx, `SELECT
block, tx, mission_id, reward, context, publisher, solution_num, solved, txtime
FROM mission
WHERE solved = FALSE
ORDER BY block DESC
LIMIT ?, ?`)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, offset, limit)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var p PublishEvent
		var solved bool
		var rewardStr sql.NullString
		err = rows.Scan(&p.Block, &p.Tx, &p.Mission, &rewardStr, &p.Data, &p.Publisher, &p.SolutionNumber, &solved, &p.TxTime)
		if err != nil {
			log.Println(err)
			continue
		}
		p.UpdateStatus(solved)
		var success bool
		p.Reward, success = big.NewInt(0).SetString(rewardStr.String, 10)
		if !success {
			p.Reward = big.NewInt(0)
		}
		p.RewardInDET, success = big.NewFloat(0).SetString(rewardStr.String)
		if !success {
			p.RewardInDET = big.NewFloat(0)
		}
		p.RewardInDET.Quo(p.RewardInDET, Decimals)
		events = append(events, p)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return events, err
}

func GetOneMission(ctx context.Context, db *sql.DB, id string) (p PublishEvent, err error) {
	stmt, err := db.PrepareContext(ctx, `SELECT
block, tx, mission_id, reward, context, publisher, solution_num, solved, txtime
FROM mission
WHERE mission_id = ?`)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	var rewardStr sql.NullString
	var solved bool

	err = stmt.QueryRowContext(ctx, id, 1).Scan(&p.Block, &p.Tx, &p.Mission, &rewardStr, &p.Data, &p.Publisher, &p.SolutionNumber, &solved, &p.TxTime)
	if err != nil {
		log.Println(err)
		return
	}
	p.UpdateStatus(solved)
	var success bool
	p.Reward, success = big.NewInt(0).SetString(rewardStr.String, 10)
	if !success {
		p.Reward = big.NewInt(0)
	}
	p.RewardInDET, success = big.NewFloat(0).SetString(rewardStr.String)
	if !success {
		p.RewardInDET = big.NewFloat(0)
	}
	p.RewardInDET.Quo(p.RewardInDET, Decimals)

	return
}

func GetSolutions(ctx context.Context, db *sql.DB, missions []string) (solutions []Solution, ids []string, err error) {
	if len(missions) <= 0 {
		err = errors.New("no mission id")
		return
	}
	query := "SELECT block, tx, mission_id, solution_id, context, solver, txtime FROM solution WHERE mission_id in ('"
	query += strings.Join(missions, "','")
	query += "') ORDER BY block DESC;"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		fmt.Printf("Database Error when retrive solution: %s", err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		var s Solution
		err1 := rows.Scan(&s.Block, &s.Tx, &s.Mission, &s.Solution, &s.Data, &s.Solver, &s.TxTime)
		if err1 != nil {
			log.Println(err1)
			continue
		}
		s.Status = Unprocessed
		solutions = append(solutions, s)
		ids = append(ids, s.Solution)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func getProcessed(ctx context.Context, db *sql.DB, solutions []string, action string) (process []Process, ids []string, err error) {
	if len(solutions) <= 0 {
		err = errors.New("no solution id")
		return
	}
	action = strings.ToLower(action)
	if action != "reject" && action != "accept" {
		err = errors.New("status SHOULD be 'accept' or 'reject'")
		return
	}
	query := "SELECT block, tx, solution_id, txtime FROM "
	query += action
	query += " WHERE solution_id in ('"
	query += strings.Join(solutions, "','")
	query += "') ORDER BY block DESC;"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		fmt.Printf("Database Error when retrive %s: %s", action, err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		var p Process
		err1 := rows.Scan(&p.Block, &p.Tx, &p.Solution, &p.TxTime)
		if err1 != nil {
			log.Println(err1)
			continue
		}
		p.Action = action
		process = append(process, p)
		ids = append(ids, p.Solution) // success ids
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func GetProcess(ctx context.Context, db *sql.DB, solutions []string) (process []Process, ids []string, err error) {
	p1, l1, e1 := getProcessed(ctx, db, solutions, "accept")
	if e1 != nil {
		fmt.Println(e1)
		return
	}
	p2, l2, e2 := getProcessed(ctx, db, solutions, "reject")
	if e2 != nil {
		fmt.Println(e2)
		return
	}
	process = append(p1, p2...)
	ids = append(l1, l2...)
	return
}

func SetFrom(ctx context.Context, db *sql.DB, from *big.Int) (err error) {
	stmtIns, err := db.PrepareContext(ctx, "INSERT INTO config (k, v) VALUES('from', ?) ON DUPLICATE KEY UPDATE v = ?")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(from.String(), from.String())
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

func GetFrom(ctx context.Context, db *sql.DB) (from *big.Int, err error) {
	from = big.NewInt(0)
	stmtIns, err := db.PrepareContext(ctx, "SELECT v FROM config WHERE k = 'from' LIMIT 1")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmtIns.Close()

	rows, err := stmtIns.Query()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rewardStr sql.NullString
		err = rows.Scan(&rewardStr)
		if err != nil {
			log.Println(err)
			continue
		}
		var success bool
		from, success = big.NewInt(0).SetString(rewardStr.String, 10)
		if !success {
			from = big.NewInt(0)
		}
		break
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return from, err
}

func addSolutionNum(ctx context.Context, db *sql.DB, missionId string) (err error) {
	stmtIns, err := db.PrepareContext(ctx, `UPDATE mission
SET    solution_num = solution_num + 1
WHERE  mission_id = ?`)
	if err != nil {
		//log.Println(err)
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(missionId)
	if err != nil {
		//log.Println(err)
		return err
	}
	return
}

func updateSolved(ctx context.Context, db *sql.DB, solutionId string) (err error) {
	// 未做成事务，暂时不考虑异常失败
	stmtIns2, err := db.PrepareContext(ctx, `UPDATE mission
SET    solved = true
WHERE  mission_id IN (SELECT mission_id
                      FROM   solution
                      WHERE  solution_id = ?) `)
	if err != nil {
		//log.Println(err)
		return err
	}
	defer stmtIns2.Close()

	_, err = stmtIns2.Exec(solutionId)
	if err != nil {
		//log.Println(err)
		return err
	}

	return
}
