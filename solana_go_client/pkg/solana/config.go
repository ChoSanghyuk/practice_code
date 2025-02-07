package solana

import (
	"time"

	"github.com/gagliardetto/solana-go/rpc"
)

type MintAccountNumber int
type TargetAccountNumber int
type TotalAccountNumber int

type WalletManagerConfig struct {
	N MintAccountNumber
	M TargetAccountNumber
	T TotalAccountNumber
}

type SolManagerConfig struct {
	RPCURL     string
	WSURL      string
	Commitment rpc.CommitmentType
	IsSync     bool
	Timeout    time.Duration
}
