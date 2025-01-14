package solana

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"
	"workspace/pkg/log"

	"github.com/gagliardetto/solana-go"
	associatedtokenaccount "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

type SolanaManager struct {
	conf   SolManagerConfig
	rpc    *rpc.Client
	ws     *ws.Client
	logger log.Logger
}

func NewSolanaManager(conf SolManagerConfig) (*SolanaManager, error) {

	rpcClient := rpc.New(conf.RPCURL)

	wsClient, err := ws.Connect(context.Background(), conf.WSURL)
	if err != nil {
		return nil, err
	}

	return &SolanaManager{
		conf:   conf,
		rpc:    rpcClient,
		ws:     wsClient,
		logger: log.GetLogger("solana_manager"),
	}, nil
}

func (m *SolanaManager) CreateAccount(ctx context.Context) *solana.Wallet {
	return solana.NewWallet()
}

func (m *SolanaManager) AccountInfo(ctx context.Context, account solana.PublicKey) (out *rpc.GetAccountInfoResult, err error) {
	return m.rpc.GetAccountInfo(ctx, account)
}

func (m *SolanaManager) RequestAirdrop(ctx context.Context, wallet *solana.Wallet, amount uint64) (string, error) {
	// _, err := m.rpc.RequestAirdrop(ctx, wallet.PublicKey(), solana.LAMPORTS_PER_SOL*amount/100, "")
	sig, err := m.rpc.RequestAirdrop(ctx, wallet.PublicKey(), solana.LAMPORTS_PER_SOL*amount, rpc.CommitmentFinalized)
	if err != nil {
		return "", err
	}
	return sig.String(), nil
}

func (m *SolanaManager) CreateAccountWithFaucet(ctx context.Context, solAmount uint64) (*solana.Wallet, error) {
	wallet := solana.NewWallet()
	_, err := m.rpc.RequestAirdrop(ctx, wallet.PublicKey(), solana.LAMPORTS_PER_SOL*solAmount, rpc.CommitmentFinalized)
	if err != nil {
		return nil, err
	}
	// m.logger.Info().Msg("airdrop sig :" + sig.String())

	return wallet, nil
}

// NewToken
func (m *SolanaManager) SetMintAccount(ctx context.Context, mintWallet, payerWallet *solana.Wallet, mintAuthority, freezeAuthority solana.PublicKey) (*solana.Signature, error) {

	min, err := m.rpc.GetMinimumBalanceForRentExemption(ctx, token.MINT_SIZE, "")
	if err != nil {
		return nil, err
	}

	// rent-fee 면제될 수 있는 lamport 양을 들고 있는 계정으로 생성.
	createAccountInst := system.NewCreateAccountInstruction(
		min,
		token.MINT_SIZE,
		token.ProgramID,
		payerWallet.PublicKey(),
		mintWallet.PublicKey(),
	).Build()

	mintInst := token.NewInitializeMintInstruction(
		9,                       // Decimal
		mintAuthority,           // Mint Authority
		freezeAuthority,         // Freeze Authority (선택)
		mintWallet.PublicKey(),  // Mint Address
		solana.SysVarRentPubkey, // SysVarRentPubkey
	).Build()

	sig, err := m.precessTransactions(ctx, []solana.Instruction{createAccountInst, mintInst}, payerWallet, []*solana.Wallet{payerWallet, mintWallet})
	if err != nil {
		return nil, err
	}

	return &sig, nil
}

func (m *SolanaManager) Mint(ctx context.Context, payerWallet *solana.Wallet, mintAccount, owner, mintAuthority solana.PublicKey, amount uint64) (*solana.Signature, solana.PublicKey, error) {

	instructions := make([]solana.Instruction, 0)
	ata, _, err := solana.FindAssociatedTokenAddress(owner, mintAccount)
	if err != nil {
		m.logger.Error().Err(err).Msg("Failed to find associated token address")
		return nil, solana.PublicKey{}, err
	}
	m.logger.Info().Msg(ata.String())

	// 이미 생성된 ata에 대해서는 생성 생략
	info, err := m.AccountInfo(ctx, ata)
	if err != nil && info == nil {
		createTokenAccountInst := associatedtokenaccount.NewCreateInstruction(
			payerWallet.PublicKey(),
			owner,
			mintAccount,
		).Build()
		instructions = append(instructions, createTokenAccountInst)
	}

	mintTokenInst := token.NewMintToInstruction(
		amount,
		mintAccount,
		ata,
		mintAuthority,
		nil,
	).Build()
	instructions = append(instructions, mintTokenInst)

	sig, err := m.precessTransactions(ctx, instructions, payerWallet, []*solana.Wallet{payerWallet})
	if err != nil {
		return nil, solana.PublicKey{}, err
	}

	return &sig, ata, nil
}

func (m *SolanaManager) SetMintAccountAndMint(ctx context.Context, payerWallet, mintWallet *solana.Wallet, owner, mintAuthority, freezeAuthority solana.PublicKey, amount uint64) (*solana.Signature, solana.PublicKey, error) {

	min, err := m.rpc.GetMinimumBalanceForRentExemption(ctx, token.MINT_SIZE, "")
	if err != nil {
		return nil, solana.PublicKey{}, err
	}

	mintAccount := mintWallet.PublicKey()
	// rent-fee 면제될 수 있는 lamport 양을 들고 있는 계정으로 생성.
	createAccountInst := system.NewCreateAccountInstruction(
		min,
		token.MINT_SIZE,
		token.ProgramID,
		payerWallet.PublicKey(),
		mintAccount,
	).Build()

	mintInst := token.NewInitializeMintInstruction(
		9,                       // Decimal
		mintAuthority,           // Mint Authority
		freezeAuthority,         // Freeze Authority (선택)
		mintAccount,             // Mint Address
		solana.SysVarRentPubkey, // SysVarRentPubkey
	).Build()

	ata, _, err := solana.FindAssociatedTokenAddress(owner, mintAccount)
	if err != nil {
		m.logger.Error().Err(err).Msg("Failed to find associated token address")
		return nil, solana.PublicKey{}, err
	}
	m.logger.Info().Msg(ata.String())

	createTokenAccountInst := associatedtokenaccount.NewCreateInstruction(
		payerWallet.PublicKey(),
		owner,
		mintAccount,
	).Build()

	mintTokenInst := token.NewMintToInstruction(
		amount,
		mintAccount,
		ata,
		mintAuthority,
		nil,
	).Build()

	sig, err := m.precessTransactions(ctx,
		[]solana.Instruction{createAccountInst, mintInst, createTokenAccountInst, mintTokenInst},
		payerWallet,
		[]*solana.Wallet{payerWallet, mintWallet})
	if err != nil {
		return nil, solana.PublicKey{}, err
	}

	return &sig, ata, nil
}

func (m *SolanaManager) TransferToken(ctx context.Context, payer, owner, receiver *solana.Wallet, mintAc solana.PublicKey, amount uint64) (*solana.Signature, error) {

	instructions := make([]solana.Instruction, 0)

	from, _, err := solana.FindAssociatedTokenAddress(owner.PublicKey(), mintAc)
	if err != nil {
		m.logger.Error().Err(err).Msg("Failed to find associated token address")
		return nil, err
	}

	// 이미 생성된 ata에 대해서는 생성 생략
	info, err := m.AccountInfo(ctx, from)
	if err != nil && info == nil {
		return nil, errors.New("미등록 token account에서 송신 시도")
	}

	to, _, err := solana.FindAssociatedTokenAddress(receiver.PublicKey(), mintAc)
	if err != nil {
		m.logger.Error().Err(err).Msg("Failed to find associated token address")
		return nil, err
	}

	// receiver가 이미 토큰 주소를 가지고 있다면 생략
	info, err = m.AccountInfo(ctx, to)
	if err != nil && info == nil {
		createTokenAccountInst := associatedtokenaccount.NewCreateInstruction(
			payer.PublicKey(),
			receiver.PublicKey(),
			mintAc,
		).Build()
		instructions = append(instructions, createTokenAccountInst)
	}

	transferInst := token.NewTransferInstruction(amount, from, to, owner.PublicKey(), nil).Build()
	instructions = append(instructions, transferInst)

	sig, err := m.precessTransactions(ctx, instructions, payer, []*solana.Wallet{payer, owner, receiver}) // todo. test - owner없이도 서명되는지

	return &sig, err
}

func (m *SolanaManager) WalletFromPK(pk string) (*solana.Wallet, error) {
	return solana.WalletFromPrivateKeyBase58(pk)
}

func (m *SolanaManager) PublicKeyFromAddr(addr string) (solana.PublicKey, error) {
	return solana.PublicKeyFromBase58(addr)
}

func (m *SolanaManager) TokenBalance(ctx context.Context, mintAc, ownerAc solana.PublicKey) (*string, error) {

	tokenAc, _, err := solana.FindAssociatedTokenAddress(ownerAc, mintAc)
	if err != nil {
		m.logger.Error().Err(err).Msg("Failed to find associated token address")
		return nil, err
	}

	balance, err := m.rpc.GetTokenAccountBalance(ctx, tokenAc, rpc.CommitmentFinalized)
	if err != nil {
		return nil, err
	}

	return &balance.Value.Amount, nil
}

func (m *SolanaManager) GetLargestTokenHolders(ctx context.Context, mintAddress solana.PublicKey) (map[string]bool, error) {
	// Fetch all token accounts for the mint
	tokenAccounts, err := m.rpc.GetTokenLargestAccounts(ctx, mintAddress, rpc.CommitmentFinalized)
	if err != nil {
		return nil, fmt.Errorf("failed to get token accounts: %w", err)
	}

	uniqueHolders := make(map[string]bool)

	for _, account := range tokenAccounts.Value {
		// fmt.Println(account)
		uniqueHolders[account.Address.String()] = true
	}

	return uniqueHolders, nil
}

func (m *SolanaManager) GetTokenHolders(ctx context.Context, mintAddress solana.PublicKey, commitment rpc.CommitmentType) (map[string]bool, error) {
	// Set up the filters for the mint
	config := rpc.GetProgramAccountsOpts{
		Filters: []rpc.RPCFilter{
			{DataSize: 165}, // Token account size
			{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: 0,                   // Mint address starts at offset 0
					Bytes:  mintAddress.Bytes(), // Base58-encoded mint address
				},
			},
		},
		Commitment: commitment,
	}

	// Call getProgramAccounts
	accounts, err := m.rpc.GetProgramAccountsWithOpts(ctx, solana.TokenProgramID, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch program accounts: %w", err)
	}

	// Extract unique owners
	holders := make(map[string]bool)
	for _, account := range accounts {
		// fmt.Println(account)
		holders[account.Pubkey.String()] = true
	}

	return holders, nil
}

/***************************************************************************************************************************************************
**********************************************************  Inner Function  ************************************************************************
***************************************************************************************************************************************************/

func (m *SolanaManager) precessTransactions(ctx context.Context, instructions []solana.Instruction, payerWallet *solana.Wallet, signers []*solana.Wallet) (solana.Signature, error) {

	// 트랜잭션 생성
	tx, err := m.buildTransactions(ctx, instructions, payerWallet.PublicKey())
	if err != nil {
		m.logger.Error().Err(err).Msg("Failed to create transaction")
		return solana.Signature{}, err
	}

	// 트랜잭션 서명
	tx, err = m.signTransaction(ctx, tx, signers)
	if err != nil {
		m.logger.Error().Err(err).Msg("Failed to sign transaction")
		return solana.Signature{}, err
	}

	// 트랜잭션 전송
	sig, err := m.sendTransaction(ctx, tx)
	if err != nil {
		m.logger.Error().Err(err).Msg("Failed to send transaction")
		return solana.Signature{}, err
	}

	return sig, nil
}

func (m *SolanaManager) buildTransactions(ctx context.Context, instructions []solana.Instruction, payer solana.PublicKey) (*solana.Transaction, error) {
	blockHash, err := m.latestBlockHash(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := solana.NewTransaction(
		instructions,
		blockHash,
		solana.TransactionPayer(payer),
	)
	if err != nil {
		m.logger.Error().Err(err).Msg("Failed to create transaction")
		return nil, err
	}

	return tx, nil
}

func (m *SolanaManager) signTransaction(ctx context.Context, tx *solana.Transaction, signers []*solana.Wallet) (*solana.Transaction, error) {
	_, err := tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			for _, signer := range signers {
				if signer.PublicKey().Equals(key) {
					return &signer.PrivateKey
				}
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (m *SolanaManager) sendTransaction(ctx context.Context, tx *solana.Transaction) (solana.Signature, error) {
	if m.conf.IsSync {
		sig, err := m.sendAndConfirmTransaction(ctx, m.rpc, m.ws, tx)
		if err != nil {
			m.logger.Error().Err(err).Msg("Failed to send transaction")
			return solana.Signature{}, err
		}

		return sig, nil
	} else {
		sig, err := m.rpc.SendTransactionWithOpts(ctx, tx, rpc.TransactionOpts{
			SkipPreflight:       true,
			PreflightCommitment: m.conf.Commitment,
		})
		if err != nil {
			m.logger.Error().Err(err).Msg("Failed to send transaction")
			return solana.Signature{}, err
		}
		return sig, nil
	}
}

func (m *SolanaManager) latestBlockHash(ctx context.Context) (solana.Hash, error) {
	recent, err := m.rpc.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return solana.Hash{}, err
	}
	return recent.Value.Blockhash, nil
}

func (m *SolanaManager) balance(ctx context.Context, pubKey solana.PublicKey) (*big.Float, error) {
	balance, err := m.rpc.GetBalance(ctx, pubKey, m.conf.Commitment)
	if err != nil {
		return nil, err
	}

	solBalance := new(big.Float).Quo(new(big.Float).SetUint64(balance.Value), new(big.Float).SetUint64(solana.LAMPORTS_PER_SOL))

	return solBalance, nil
}

// customize functions of sendandconfirmtransaction package
func (m *SolanaManager) sendAndConfirmTransaction(
	ctx context.Context,
	rpcClient *rpc.Client,
	wsClient *ws.Client,
	transaction *solana.Transaction,
) (signature solana.Signature, err error) {
	opts := rpc.TransactionOpts{
		SkipPreflight:       false,
		PreflightCommitment: m.conf.Commitment,
	}

	return m.sendAndConfirmTransactionWithOpts(
		ctx,
		rpcClient,
		wsClient,
		transaction,
		opts,
		nil,
	)
}

// Send and wait for confirmation of a transaction.
func (m *SolanaManager) sendAndConfirmTransactionWithOpts(
	ctx context.Context,
	rpcClient *rpc.Client,
	wsClient *ws.Client,
	transaction *solana.Transaction,
	opts rpc.TransactionOpts,
	timeout *time.Duration,
) (sig solana.Signature, err error) {
	sig, err = rpcClient.SendTransactionWithOpts(
		ctx,
		transaction,
		opts,
	)
	if err != nil {
		return sig, err
	}
	_, err = m.waitForConfirmation(
		ctx,
		wsClient,
		sig,
		timeout,
	)
	return sig, err
}

// WaitForConfirmation waits for a transaction to be confirmed.
// If the transaction was confirmed, but it failed while executing (one of the instructions failed),
// then this function will return an error (true, error).
// If the transaction was confirmed, and it succeeded, then this function will return nil (true, nil).
func (m *SolanaManager) waitForConfirmation(
	ctx context.Context,
	wsClient *ws.Client,
	sig solana.Signature,
	timeout *time.Duration,
) (confirmed bool, err error) {
	sub, err := wsClient.SignatureSubscribe(
		sig,
		m.conf.Commitment,
	)
	if err != nil {
		return false, err
	}
	defer sub.Unsubscribe()

	if timeout == nil {
		t := 2 * time.Minute // random default timeout
		timeout = &t
	}

	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		case <-time.After(*timeout):
			return false, fmt.Errorf("timeout")
		case resp, ok := <-sub.Response():
			if !ok {
				return false, fmt.Errorf("subscription closed")
			}
			if resp.Value.Err != nil {
				// The transaction was confirmed, but it failed while executing (one of the instructions failed).
				return true, fmt.Errorf("confirmed transaction with execution error: %v", resp.Value.Err)
			} else {
				// Success! Confirmed! And there was no error while executing the transaction.
				return true, nil
			}
		case err := <-sub.Err():
			return false, err
		}
	}
}
