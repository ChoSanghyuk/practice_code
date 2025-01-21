package parameters

type MintReq struct {
	Amount int `json:"amount"`
}

type TransferReq struct {
	Amount int `json:"amount"`
}

type TargetTokenBalanceReq struct {
	MintAddress  string `json:"mint_address"`
	OwnerAddress string `json:"owner_address"`
}

type TargetSolBalanceReq struct {
	OwnerAddress string `json:"owner_address"`
}

type TargetTransferTokenReq struct {
	MintWallet   string `json:"mint_wallet"`
	OwnerWallet  string `json:"owner_wallet"`
	TargetWallet string `json:"target_wallet"`
	Amount       int    `json:"amount"`
}

type DeployRes struct {
	Signature  string `json:"signature"`
	MintWallet string `json:"mint_wallet"`
}

type TransactionRes struct {
	Signature string `json:"signature"`
}

type TokenBalanceRes struct {
	MintAddress  string `json:"mint_address"`
	OwnerAddress string `json:"owner_address"`
	Balance      string `json:"balance"`
}

type TargetBalanceRes struct {
	Balance string `json:"balance"`
}
