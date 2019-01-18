package jsonrpc

import (
	. "github.com/xyths/ot-engine/types"
	"github.com/xyths/ot-engine/collect"
	"fmt"
	"database/sql"
)

type EngineRPC struct {
	Version string
	DB *sql.DB
}

func (e *EngineRPC) GetAllPublished(address string, offset int, limit int) (missions []Mission) {
	missions1, err := collect.GetMissions(e.DB, address, limit)
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

func (e *EngineRPC) GetUnsolved(address string, limit int) (missions []Mission) {
	// Only Unsolved Mission
	return missions
}
