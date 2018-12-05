package collect

import (
	. "github.com/xyths/ot-engine/types"
	"database/sql"
	"github.com/xyths/ot-engine/database"
	"fmt"
)

func GetMissions(db *sql.DB, address string, limit int) (missions []Mission, err error) {
	publishList, err := database.GetPublished(db, address, limit)
	if err != nil {
		fmt.Printf("Error when GetMission: %s", err)
		return missions, err
	}
	var missionIdList []string
	for _, p := range publishList {
		var m Mission
		m.PublishEvent = p
		missionIdList = append(missionIdList, p.Mission)
		missions = append(missions, m)
	}
	solutions, err := GetSolutions(db, missionIdList)
	if err!= nil {
		fmt.Printf("Error when GetSolutions: %s", err)
		return missions, err
	}
	fillMissions(missions, solutions)

	return missions, err
}

func GetSolutions(db *sql.DB, missions []string) (solutions []Solution, err error) {
	solutions, ids, err := database.GetSolutions(db, missions)
	processList, err := GetProcess(db,ids)
	fillSolutions(solutions,processList)
	return solutions, err
}

func GetProcess(db *sql.DB, solutions []string) (process []Process, err error) {
	solutions, ids, err := database.GetProcess(db, solutions)
	return solutions, err
}

func fillMissions(missions []Mission, solutions []Solution) {

}

func fillSolutions(solutions []Solution, process []Process) {

}