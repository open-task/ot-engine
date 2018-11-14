package types

import "math/big"

type PublishEvent struct {
	Block uint64
	Tx string
	Mission string
	Reward *big.Int
	Publisher string
}

type SolveEvent struct {
	Block int
	Tx string
	Solution string
	Mission string
	Data string
}

type AcceptEvent struct {
	Block int
	Tx string
	Solution string
}


type RejectEvent struct {
	Block int
	Tx string
	Solution string
}

type ConfirmEvent struct {
	Block int
	Tx string
	Solution string
	Arbitration string
}