package jsonrpc

import (
	"database/sql"
	"fmt"
	"github.com/open-task/ot-engine/collect"
	. "github.com/open-task/ot-engine/types"
)

type EngineRPC struct {
	Version string
	DB      *sql.DB
}

func (e *EngineRPC) GetAllPublished(offset int, limit int) (missions []Mission) {
	missions1, err := collect.GetAllMissions(e.DB, offset, limit)
	if err != nil {
		fmt.Printf("Error When GetMission: %s", err.Error())
	} else {
		missions = missions1
		fmt.Println(missions1)
	}
	return missions
}

func (e *EngineRPC) GetPublished(address string, limit int) (missions []Mission) {
	missions1, err := collect.GetMissions(e.DB, address, limit)
	if err != nil {
		fmt.Printf("Error When GetMission: %s", err.Error())
	} else {
		missions = missions1
		fmt.Println(missions1)
	}
	return missions
}

func (e *EngineRPC) GetUnsolved(offset int, limit int) (missions []Mission) {
	missions1, err := collect.GetUnsolved(e.DB, offset, limit)
	if err != nil {
		fmt.Printf("Error When GetMission: %s", err.Error())
	} else {
		missions = missions1
		fmt.Println(missions1)
	}
	return missions
}

func (e *EngineRPC) GetMissionInfo(id string) (mission Mission) {
	mission, err := collect.GetOneMission(e.DB, id)
	if err != nil {
		fmt.Printf("Error When GetMission: %s", err.Error())
		return Mission{}
	}
	return mission
}