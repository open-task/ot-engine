package jsonrpc

import (
	"database/sql"
	"github.com/open-task/ot-engine/collect"
	. "github.com/open-task/ot-engine/types"
	"log"
)

type EngineRPC struct {
	Version string
	DB      *sql.DB
}

func (e *EngineRPC) GetAllPublished(offset int, limit int) (missions []Mission) {
	log.Printf("GetAllPublished called: offset = %d, limit = %d\n", offset, limit)
	err := e.DB.Ping()
	if err != nil {
		log.Printf("Error when ping database: %s", err.Error())
	}
	missions1, err := collect.GetAllMissions(e.DB, offset, limit)
	if err != nil {
		log.Printf("Error When GetMission: %s", err.Error())
	} else {
		missions = missions1
		log.Println(missions1)
	}
	log.Printf("GetAllPublished return %d missions\n", len(missions))
	return missions
}

func (e *EngineRPC) GetPublished(address string, limit int) (missions []Mission) {
	missions1, err := collect.GetMissions(e.DB, address, limit)
	if err != nil {
		log.Printf("Error When GetMission: %s", err.Error())
	} else {
		missions = missions1
		log.Println(missions1)
	}
	return missions
}

func (e *EngineRPC) GetUnsolved(offset int, limit int) (missions []Mission) {
	missions1, err := collect.GetUnsolved(e.DB, offset, limit)
	if err != nil {
		log.Printf("Error When GetMission: %s", err.Error())
	} else {
		missions = missions1
		log.Println(missions1)
	}
	return missions
}

func (e *EngineRPC) GetMissionInfo(id string) (mission Mission) {
	mission, err := collect.GetOneMission(e.DB, id)
	if err != nil {
		log.Printf("Error When GetMission: %s", err.Error())
		return Mission{}
	}
	return mission
}
