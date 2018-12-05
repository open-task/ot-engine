package jsonrpc

import (
	"github.com/xyths/ot-engine/types"
	"github.com/xyths/ot-engine/collect"
	"fmt"
	"database/sql"
)

type EngineRPC struct {
	Version string
	DB *sql.DB
}

func (e *EngineRPC) GetPublished(address string, limit int) (missions []types.Mission) {
	missions1, err := collect.GetMissions(e.DB, address, limit)
	if err != nil {
		fmt.Printf("Error When GetMission: %s", err.Error())
	} else {
		missions = missions1
	}
	return missions
}
