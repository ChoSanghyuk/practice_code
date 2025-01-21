package solana

import "github.com/gagliardetto/solana-go/rpc"

type MintAccountNumber int
type TargetAccountNumber int

type WalletManagerConfig struct {
	N MintAccountNumber
	M TargetAccountNumber
}

type SolManagerConfig struct {
	RPCURL     string
	WSURL      string
	Commitment rpc.CommitmentType
	IsSync     bool
}
