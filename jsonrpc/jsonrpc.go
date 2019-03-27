package jsonrpc

import (
	"context"
	"database/sql"
	"github.com/open-task/ot-engine/collect"
	. "github.com/open-task/ot-engine/types"
	"log"
	"time"
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	missions1, err := collect.GetAllMissions(ctx, e.DB, offset, limit)
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	missions1, err := collect.GetMissions(ctx, e.DB, address, limit)
	if err != nil {
		log.Printf("Error When GetMission: %s", err.Error())
	} else {
		missions = missions1
		log.Println(missions1)
	}
	return missions
}

func (e *EngineRPC) GetUnsolved(offset int, limit int) (missions []Mission) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	missions1, err := collect.GetUnsolved(ctx, e.DB, offset, limit)
	if err != nil {
		log.Printf("Error When GetMission: %s", err.Error())
	} else {
		missions = missions1
		log.Println(missions1)
	}
	return missions
}

func (e *EngineRPC) GetMissionInfo(id string) (mission Mission) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	mission, err := collect.GetOneMission(ctx, e.DB, id)
	if err != nil {
		log.Printf("Error When GetMission: %s", err.Error())
		return Mission{}
	}
	return mission
}
