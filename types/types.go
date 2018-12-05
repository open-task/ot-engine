package types

import (
	"math/big"
)

type PublishEvent struct {
	Block     uint64
	Tx        string
	Mission   string
	Reward    *big.Int
	Publisher string
}

type SolveEvent struct {
	Block    uint64
	Tx       string
	Solution string
	Mission  string
	Data     string
	Solver   string
}

type ProcessEvent struct {
	Block    uint64
	Tx       string
	Solution string
	Time     string // type is string, just for output
	Status   string // accept or reject
}
type Process ProcessEvent
type AcceptEvent ProcessEvent
type RejectEvent ProcessEvent

type ConfirmEvent struct {
	Block       uint64
	Tx          string
	Solution    string
	Arbitration string
}

type ProcessStatus struct {
	Status string // Unprocessed, Accepted, Rejected
	Process       // AcceptEvent or RejectEvent
	// TODO: Argue and Confirm
}

type Solution struct {
	SolveEvent
	Status string // Unprocessed, Accepted, Rejected
	Process Process       // AcceptEvent or RejectEvent
}

type Mission struct {
	PublishEvent
	Solutions []Solution
}

/*func (m Mission) JsonString(prefix string) (json string) {
	json += fmt.Sprintln("{")
	json += m.PublishEvent.JsonString(prefix + "\t")
	for _, s := range m.Solutions {
		json += s.JsonString(prefix + "\t")
	}
	json += fmt.Sprintln("}")
	return json
}


func (p PublishEvent) JsonString(prefix string) (json string) {
	json += fmt.Sprintln("{")

	json += prefix + string(p.Block)
	json += prefix + p.Tx
	json += prefix + p.Mission

	json += fmt.Sprintln("}")
	return json
}*/
