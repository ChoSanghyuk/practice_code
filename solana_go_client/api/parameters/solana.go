package parameters

type PerformanceCreateAccountReq struct {
	SolAmount int `json:"sol_amount" validate:"required,gte=1,lte=1000"`
}

type PerformanceCreateAccountRes struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

// type PerformanceCreateSPLTokenReq struct {
// 	PayerPrivateKey string `json:"payer_private_key"`
// 	MintAuthority   string `json:"mint_authority"`
// 	FreezeAuthority string `json:"freeze_authority"`
// }

type PerformanceCreateSPLTokenRes struct {
	Signature   string `json:"signature"`
	MintAccount string `json:"mint_account"`
}

type PerformanceMintReq struct {
	// PayerPrivateKey string `json:"payer_private_key"`
	// MintAccount     string `json:"mint_account"`
	// MintAuthority   string `json:"mint_authority"`
	Amount int `json:"amount"`
}

type CreateTokenWithMintReq struct {
	// PayerPrivateKey string `json:"payer_private_key"`
	// MintAccount     string `json:"mint_account"`
	// MintAuthority   string `json:"mint_authority"`
	Amount int `json:"amount"`
}

type PerformanceMintRes struct {
	Signature string `json:"signature"`
}

type TransferTokenReq struct {
	Amount int `json:"amount"`
}

type TokenBalanceRes struct {
	MintAccount  string `json:"mint_account"`
	OwnerAccount string `json:"owner_account"`
	Balance      string `json:"balance"`
}

type TargetTokenBalanceReq struct {
	MintAccount  string `json:"mint_account"`
	OwnerAccount string `json:"owner_account"`
}

type TargetSolBalanceReq struct {
	OwnerAccount string `json:"owner_account"`
}

type TargetBalanceRes struct {
	Balance string `json:"balance"`
}
