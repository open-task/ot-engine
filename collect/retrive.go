package collect

import (
	. "github.com/xyths/ot-engine/types"
	"database/sql"
	"github.com/xyths/ot-engine/database"
	"fmt"
)

func GetAllMissions(db *sql.DB, offset int, limit int) (missions []Mission, err error) {
	publishList, err := database.GetAllPublished(db, offset, limit)
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
	if err != nil {
		fmt.Printf("Error when GetSolutions: %s", err)
		return missions, err
	}
	fillMissions(&missions, &solutions)

	return missions, err
}

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
	if err != nil {
		fmt.Printf("Error when GetSolutions: %s", err)
		return missions, err
	}
	fillMissions(&missions, &solutions)

	return missions, err
}

func GetSolutions(db *sql.DB, missions []string) (solutions []Solution, err error) {
	solutions, ids, err := database.GetSolutions(db, missions)
	processList, err := GetProcess(db, ids)
	fillSolutions(&solutions, &processList)
	return solutions, err
}

func GetProcess(db *sql.DB, solutions []string) (process []Process, err error) {
	process, _, err = database.GetProcess(db, solutions)
	return process, err
}

func fillMissions(missions *[]Mission, solutions *[]Solution) {
	for i := range *solutions {
		for j := range *missions {
			if (*solutions)[i].Mission == (*missions)[j].Mission {
				(*missions)[j].Solutions = append((*missions)[j].Solutions, (*solutions)[i])
				break
			}
		}
	}
}

func fillSolutions(solutions *[]Solution, process *[]Process) {
	for i := range *process {
		for j := range *solutions {
			if (*process)[i].Solution == (*solutions)[j].Solution {
				(*solutions)[j].Process = (*process)[i]
				(*solutions)[j].Status = (*process)[i].Status
				break
			}
		}
	}
}
