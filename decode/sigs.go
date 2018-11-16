package decode

import "github.com/ethereum/go-ethereum/crypto"

var (
	publishSig = []byte("Publish(string,uint256)")
	solveSig   = []byte("Solve(string,string,string)")
	acceptSig  = []byte("Accept(string)")
	rejectSig  = []byte("Reject(string)")
	confirmSig = []byte("Confirm(string,string)")

	PublishSigHash = crypto.Keccak256Hash(publishSig)
	SolveSigHash = crypto.Keccak256Hash(solveSig)
	AcceptSigHash = crypto.Keccak256Hash(acceptSig)
	RejectSigHash = crypto.Keccak256Hash(rejectSig)
	ConfirmSigHash = crypto.Keccak256Hash(confirmSig)
)