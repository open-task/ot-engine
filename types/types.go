package types

import (
	"math/big"
)

type PublishEvent struct {
	Block          uint64     `json:"block"`
	Tx             string     `json:"tx"`
	Mission        string     `json:"mission_id"`
	Reward         *big.Int   `json:"reward_wei"`
	RewardInDET    *big.Float `json:"reward_det"`
	Data           string     `json:"data"`
	Publisher      string     `json:"publisher"`
	SolutionNumber uint       `json:"solution_number"`
	Status         string     `json:"status"` // Published, Unsolve, Solved
	TxTime         string     `json:"time"`
}

func (p *PublishEvent) UpdateStatus(solved bool) bool {
	if solved {
		p.Status = Solved
	} else if p.SolutionNumber > 0 {
		p.Status = Unsolve
	} else {
		p.Status = Published
	}
	return true
}

const (
	Published = "Published"
	Unsolve   = "Unsolved"
	Solved    = "Solved"
)

type SolveEvent struct {
	Block    uint64 `json:"block"`
	Tx       string `json:"tx"`
	Solution string `json:"solution_id"`
	Mission  string `json:"mission_id"`
	Data     string `json:"data"`
	Solver   string `json:"solver"`
	TxTime   string `json:"time"`
}

type ProcessEvent struct {
	Block    uint64 `json:"block"`
	Tx       string `json:"tx"`
	Solution string `json:"solution_id"`
	Status   string `json:"status"` // accept or reject
	TxTime   string `json:"time"`   // type is string, just for output
}
type Process ProcessEvent
type AcceptEvent ProcessEvent
type RejectEvent ProcessEvent

type ConfirmEvent struct {
	Block       uint64 `json:"block"`
	Tx          string `json:"tx"`
	Solution    string `json:"solution_id"`
	Arbitration string `json:"arbitration_id"`
	TxTime      string `json:"time"`
}

type ProcessStatus struct {
	Status string `json:"status"` // Unprocessed, Accepted, Rejected
	Process                       // AcceptEvent or RejectEvent
	// TODO: Argue and Confirm
}

type Solution struct {
	SolveEvent
	Status  string  `json:"status"`  // Unprocessed, Accepted, Rejected
	Process Process `json:"process"` // AcceptEvent or RejectEvent
}

type Mission struct {
	PublishEvent
	Solutions []Solution `json:"solutions"`
}
