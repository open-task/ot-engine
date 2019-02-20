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
	TxTime    string
}

type SolveEvent struct {
	Block    uint64
	Tx       string
	Solution string
	Mission  string
	Data     string
	Solver   string
	TxTime   string
}

type ProcessEvent struct {
	Block    uint64
	Tx       string
	Solution string
	TxTime   string // type is string, just for output
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
	TxTime      string
}

type ProcessStatus struct {
	Status string // Unprocessed, Accepted, Rejected
	Process       // AcceptEvent or RejectEvent
	// TODO: Argue and Confirm
}

type Solution struct {
	SolveEvent
	Status  string  // Unprocessed, Accepted, Rejected
	Process Process // AcceptEvent or RejectEvent
}

type Mission struct {
	PublishEvent
	Solutions []Solution
}
