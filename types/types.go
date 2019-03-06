package types

import (
	"math/big"
)

type PublishEvent struct {
	Block     uint64   `json:"block"`
	Tx        string   `json:"tx"`
	Mission   string   `json:"mission"`
	Reward    *big.Int `json:"reward"`
	Data      string   `json:"data"`
	Publisher string   `json:"publisher"`
	TxTime    string   `json:"time"`
}

type SolveEvent struct {
	Block    uint64 `json:"block"`
	Tx       string `json:"tx"`
	Solution string `json:"solution"`
	Mission  string `json:"mission"`
	Data     string `json:"data"`
	Solver   string `json:"solver"`
	TxTime   string `json:"time"`
}

type ProcessEvent struct {
	Block    uint64 `json:"block"`
	Tx       string `json:"tx"`
	Solution string `json:"solution"`
	TxTime   string `json:"time"` // type is string, just for output
	Status   string `json:"status"`  // accept or reject
}
type Process ProcessEvent
type AcceptEvent ProcessEvent
type RejectEvent ProcessEvent

type ConfirmEvent struct {
	Block       uint64 `json:"block"`
	Tx          string `json:"tx"`
	Solution    string `json:"solution"`
	Arbitration string `json:"arbitration"`
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
