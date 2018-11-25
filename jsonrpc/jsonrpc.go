package jsonrpc

import "github.com/xyths/ot-engine/types"

type EngineRPC struct {
	Version string
}

func (t *EngineRPC) GetPublished(address string, limit int) (events []types.PublishEvent) {
	var p types.PublishEvent
	p.Mission = "m1"
	events = append(events, p)
	return events
}
