package collect

import (
	"context"
	"database/sql"
	"github.com/open-task/ot-engine/database"
	. "github.com/open-task/ot-engine/types"
	"log"
)

const (
	Accept = "accept"
	Reject = "reject"
)

func GetAllMissions(ctx context.Context, db *sql.DB, offset int, limit int) (missions []Mission, err error) {
	publishList, err := database.GetAllPublished(ctx, db, offset, limit)
	if err != nil {
		log.Printf("Error when GetMission: %s", err)
		return missions, err
	}
	var missionIdList []string
	for _, p := range publishList {
		var m Mission
		m.PublishEvent = p
		missionIdList = append(missionIdList, p.Mission)
		missions = append(missions, m)
	}
	solutions, err := GetSolutions(ctx, db, missionIdList)
	if err != nil {
		log.Printf("Error when GetSolutions: %s", err)
		return missions, err
	}
	fillMissions(&missions, &solutions)

	return missions, err
}

func GetMissions(ctx context.Context, db *sql.DB, address string, limit int) (missions []Mission, err error) {
	publishList, err := database.GetPublished(ctx, db, address, limit)
	if err != nil {
		log.Printf("Error when GetMission: %s", err)
		return missions, err
	}
	var missionIdList []string
	for _, p := range publishList {
		var m Mission
		m.PublishEvent = p
		missionIdList = append(missionIdList, p.Mission)
		missions = append(missions, m)
	}
	solutions, err := GetSolutions(ctx, db, missionIdList)
	if err != nil {
		log.Printf("Error when GetSolutions: %s", err)
		return missions, err
	}
	fillMissions(&missions, &solutions)

	return missions, err
}
func GetUnsolved(ctx context.Context, db *sql.DB, offset int, limit int) (missions []Mission, err error) {
	publishList, err := database.GetUnsolved(ctx, db, offset, limit)
	if err != nil {
		log.Printf("Error when GetMission: %s", err)
		return missions, err
	}
	var missionIdList []string
	for _, p := range publishList {
		var m Mission
		m.PublishEvent = p
		missionIdList = append(missionIdList, p.Mission)
		missions = append(missions, m)
	}
	solutions, err := GetSolutions(ctx, db, missionIdList)
	if err != nil {
		log.Printf("Error when GetSolutions: %s", err)
		return missions, err
	}
	fillMissions(&missions, &solutions)

	return missions, err
}

func GetOneMission(ctx context.Context, db *sql.DB, id string) (m Mission, err error) {
	publish, err := database.GetOneMission(ctx, db, id)
	if err != nil {
		log.Printf("Error when GetOneMission: %s", err)
		return m, err
	}
	var missions []Mission
	var missionIdList []string

	m.PublishEvent = publish
	missionIdList = append(missionIdList, publish.Mission)
	missions = append(missions, m)

	solutions, err := GetSolutions(ctx, db, missionIdList)
	if err != nil {
		log.Printf("Error when GetSolutions: %s", err)
		// half result, no solutions
		return m, nil
	}
	fillMissions(&missions, &solutions)
	m = missions[0]
	return
}

func GetSolutions(ctx context.Context, db *sql.DB, missions []string) (solutions []Solution, err error) {
	solutions, ids, err := database.GetSolutions(ctx, db, missions)
	processList, err := GetProcess(ctx, db, ids)
	fillSolutions(&solutions, &processList)
	return solutions, err
}

func GetProcess(ctx context.Context, db *sql.DB, solutions []string) (process []Process, err error) {
	process, _, err = database.GetProcess(ctx, db, solutions)
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
			if (*process)[i].Solution != (*solutions)[j].Solution {
				continue
			}

			(*solutions)[j].Process = (*process)[i]
			if (*process)[i].Action == Accept {
				(*solutions)[j].Status = Accepted
			} else if (*process)[i].Action == Reject {
				(*solutions)[j].Status = Rejected
			}
			break
		}
	}
}
