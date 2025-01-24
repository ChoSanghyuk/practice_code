package handlers

import (
	"context"
	"sync"
	"workspace/api/middlewares"
	"workspace/api/parameters"
	"workspace/api/types"
	"workspace/pkg/log"
	"workspace/pkg/solana"

	solanaLib "github.com/gagliardetto/solana-go"
	"github.com/gofiber/fiber/v2"
)

type SplHandler struct {
	solm   *solana.SolanaManager
	wm     *solana.WalletManager
	logger log.Logger
}

func NewSplHandler(solm *solana.SolanaManager, wm *solana.WalletManager) *SplHandler {
	return &SplHandler{
		solm:   solm,
		wm:     wm,
		logger: log.GetLogger("solana_handler"),
	}
}

func (h *SplHandler) Append(r fiber.Router) {
	r.Post("/deploy", h.DeployToken)
	r.Post("/deploy-with-mint", middlewares.Validate(&parameters.MintReq{}), h.DeployWithMint)
	r.Post("/set-mint-account", middlewares.Validate(&parameters.MintReq{}), h.SetMintAccount)
	r.Post("/set-token-account", h.SetTokenAccount)
	r.Post("/mint", middlewares.Validate(&parameters.MintReq{}), h.Mint)
	r.Post("/transfer", middlewares.Validate(&parameters.TransferReq{}), h.TransferToken)
	r.Post("/query", h.TokenBalance)
	r.Post("/target/query", middlewares.Validate(&parameters.TargetTokenBalanceReq{}), h.TargetTokenBalance)
	r.Post("/target/transfer", middlewares.Validate(&parameters.TargetTransferTokenReq{}), h.TargetTransferToken)
}

// @Summary 테스트 사전 준비
// @Description 기존 생성한 Mint Account에 deploy 후 민팅
// @Tags /spl
// @Accept json
// @Produce json
// @Param body body parameters.MintReq true "Performance Create SPL Token And Mint Request"
// @Success 200
// @Router /spl/set-mint-account [post]
func (h *SplHandler) SetMintAccount(c *fiber.Ctx) error {
	reqID := c.Locals(types.RequestID).(string)
	lg := h.logger.SetReqID(reqID)
	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

	req := &parameters.MintReq{}
	err := c.BodyParser(req)
	if err != nil {
		return err
	}

	lg.Debug().
		Msg("POST /spl/set-mint-account 호출")

	n := h.wm.MintN()

	var wg sync.WaitGroup
	wg.Add(n)
	ch := make(chan error)
	for i := 0; i < n; i++ {
		go func(ch chan error) {
			defer wg.Done()
			mintWllt, initWllt := h.wm.NextMintInitWallet()
			auth := initWllt.PublicKey()

			_, _, err := h.solm.SetMintAccountAndMint(ctx, initWllt, mintWllt, initWllt.PublicKey(), auth, auth, uint64(req.Amount))
			if err != nil {
				lg.Error().
					Str("mint_address", mintWllt.PublicKey().String()).
					Str("init_holder_address", initWllt.PublicKey().String()).
					Err(err).
					Msg("set mint account 오류 발생")
				ch <- err
			}
		}(ch)
	}

	wg.Wait()
	close(ch)

	for err := range ch {
		if err != nil {
			return err
		}
	}
	return c.Status(fiber.StatusOK).
		JSON(
			parameters.NewSuccessResponse(nil),
		)
}

// @Summary 테스트 사전 준비
// @Description token n개에 대한 target m명의 token account 전체 생성
// @Tags /spl
// @Accept json
// @Produce json
// @Success 200
// @Router /spl/set-token-account [post]
func (h *SplHandler) SetTokenAccount(c *fiber.Ctx) error {
	reqID := c.Locals(types.RequestID).(string)
	lg := h.logger.SetReqID(reqID)
	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

	lg.Debug().
		Msg("POST /spl/set-token-account 호출")

	n := h.wm.MintN()
	m := h.wm.TargetN()

	var wg sync.WaitGroup
	wg.Add(n * m)
	ch := make(chan error)
	for i := 0; i < n; i++ {
		mintWllt, initWllt := h.wm.NextMintInitWallet()

		for j := 0; j < m; j++ {
			trgtWllt := h.wm.NextTrgtWallet()
			go func(ch chan error, mintWllt, initWllt, trgtWllt *solanaLib.Wallet) {
				defer wg.Done()

				_, err := h.solm.CreateAta(ctx, initWllt, trgtWllt, mintWllt.PublicKey())
				if err != nil {
					lg.Error().
						Str("mint_address", mintWllt.PublicKey().String()).
						Str("init_holder_address", initWllt.PublicKey().String()).
						Str("target_address", trgtWllt.PublicKey().String()).
						Err(err).
						Msg("create token address 오류 발생")
					ch <- err
				}
			}(ch, mintWllt, initWllt, trgtWllt)
		}
	}

	wg.Wait()
	close(ch)

	for err := range ch {
		if err != nil {
			return err
		}
	}
	return c.Status(fiber.StatusOK).
		JSON(
			parameters.NewSuccessResponse(nil),
		)
}

// @Summary SPL Token 생성
// @Description SPL Token을 생성합니다
// @Tags /spl
// @Accept json
// @Produce json
// @Success 200 {object} parameters.CommonResponse{data=parameters.DeployRes}
// @Router /spl/deploy [post]
func (h *SplHandler) DeployToken(c *fiber.Ctx) error {
	reqID := c.Locals(types.RequestID).(string)
	lg := h.logger.SetReqID(reqID)
	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

	lg.Debug().
		Msg("POST /spl 호출")

	_, initW := h.wm.NextMintInitWallet()

	mintWallet := h.solm.CreateAccount(context.Background())
	sig, err := h.solm.SetMintAccount(ctx, mintWallet, initW, initW.PublicKey(), initW.PublicKey())
	if err != nil {
		lg.Error().
			Str("mint_address", mintWallet.PublicKey().String()).
			Str("init_holder_address", initW.PublicKey().String()).
			Err(err).
			Msg("deploy 오류 발생")
		return err
	}

	return c.Status(fiber.StatusOK).
		JSON(
			parameters.NewSuccessResponse(
				parameters.DeployRes{
					Signature:  sig.String(),
					MintWallet: mintWallet.PrivateKey.String(),
				},
			),
		)
}

// @Summary Token Account 생성 후 민팅
// @Description Token Account를 생성한 뒤 민팅합니다.
// @Tags /spl
// @Accept json
// @Produce json
// @Param body body parameters.MintReq true "Performance Create SPL Token And Mint Request"
// @Success 200 {object} parameters.CommonResponse{data=parameters.DeployRes}
// @Router /spl/deploy-with-mint [post]
func (h *SplHandler) DeployWithMint(c *fiber.Ctx) error {
	reqID := c.Locals(types.RequestID).(string)
	lg := h.logger.SetReqID(reqID)
	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

	req := &parameters.MintReq{}
	err := c.BodyParser(req)
	if err != nil {
		return err
	}

	lg.Debug().
		Msg("POST /spl/deploy-with-mint 호출")

	initWllt := h.wm.NextUserWallet()
	auth := initWllt.PublicKey()
	mintWallet := h.wm.NewWallet()

	sig, _, err := h.solm.SetMintAccountAndMint(ctx, initWllt, mintWallet, initWllt.PublicKey(), auth, auth, uint64(req.Amount))
	if err != nil {
		lg.Error().
			Str("mint_address", mintWallet.PublicKey().String()).
			Str("init_holder_address", initWllt.PublicKey().String()).
			Err(err).
			Msg("deploy with mint 오류 발생")
		return err
	}

	return c.Status(fiber.StatusOK).
		JSON(
			parameters.NewSuccessResponse(
				parameters.DeployRes{
					Signature:  sig.String(),
					MintWallet: mintWallet.PrivateKey.String(),
				},
			),
		)
}

// @Summary 민팅
// @Description 배포 완료 상태에서 target wallet에 추가 민팅합니다.
// @Tags /spl
// @Accept json
// @Produce json
// @Param body body parameters.MintReq true "Performance SPL Token Mint Request"
// @Success 200 {object} parameters.CommonResponse{data=parameters.TransactionRes}
// @Router /spl/mint [post]
func (h *SplHandler) Mint(c *fiber.Ctx) error {
	reqID := c.Locals(types.RequestID).(string)
	lg := h.logger.SetReqID(reqID)
	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

	req := &parameters.MintReq{}
	err := c.BodyParser(req)
	if err != nil {
		return err
	}

	lg.Debug().
		Msg("POST /spl/mint 호출")

	mintWllt, initWllt := h.wm.NextMintInitWallet()
	trgtWllt := h.wm.NextTrgtWallet()
	auth := initWllt.PublicKey() // mint authority

	sig, _, err := h.solm.Mint(ctx, initWllt, mintWllt.PublicKey(), trgtWllt.PublicKey(), auth, uint64(req.Amount))
	if err != nil {
		lg.Error().
			Str("mint_address", mintWllt.PublicKey().String()).
			Str("init_address", initWllt.PublicKey().String()).
			Str("target_address", trgtWllt.PublicKey().String()).
			Err(err).
			Msg("Mint 오류 발생")
		return err
	}

	return c.Status(fiber.StatusOK).
		JSON(
			parameters.NewSuccessResponse(
				parameters.TransactionRes{
					Signature: sig.String(),
				},
			),
		)
}

// @Summary Token Transfer
// @Description Token을 전송합니다.
// @Tags /spl
// @Accept json
// @Produce json
// @Param body body parameters.TransferReq true "Transfer SPL Token"
// @Success 200 {object} parameters.CommonResponse{data=parameters.TransactionRes}
// @Router /spl/transfer [post]
func (h *SplHandler) TransferToken(c *fiber.Ctx) error {
	reqID := c.Locals(types.RequestID).(string)
	lg := h.logger.SetReqID(reqID)
	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

	req := &parameters.TransferReq{}
	err := c.BodyParser(req)
	if err != nil {
		return err
	}

	lg.Debug().
		Msg("POST /spl/transfer 호출") // todo. 이거 url 자동화

	mintWl, sender := h.wm.NextMintInitWallet()
	receiver := h.wm.NextTrgtWallet()

	sig, err := h.solm.TransferToken(ctx, sender, sender, receiver, mintWl.PublicKey(), uint64(req.Amount))
	if err != nil {
		lg.Error().
			Str("mint_address", mintWl.PublicKey().String()).
			Str("sender_address", sender.PublicKey().String()).
			Str("target_address", receiver.PublicKey().String()).
			Err(err).
			Msg("transfer 오류 발생")
		return err
	}

	return c.Status(fiber.StatusOK).
		JSON(
			parameters.NewSuccessResponse(
				parameters.TransactionRes{
					Signature: sig.String(),
				},
			),
		)
}

// @Summary Token Query
// @Description Token Balance를 조회합니다
// @Tags /spl
// @Accept json
// @Produce json
// @Success 200 {object} parameters.CommonResponse{data=parameters.TokenBalanceRes}
// @Router /spl/query [get]
func (h *SplHandler) TokenBalance(c *fiber.Ctx) error {

	reqID := c.Locals(types.RequestID).(string)
	lg := h.logger.SetReqID(reqID)
	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

	lg.Debug().
		Msg("GET /spl/query 호출")

	mintWl, trgtWllt := h.wm.NextMintInitWallet()
	// trgtWllt := h.wm.NextTrgtWallet()

	balance, err := h.solm.TokenBalance(ctx, mintWl.PublicKey(), trgtWllt.PublicKey())
	if err != nil {
		lg.Error().
			Str("mint_address", mintWl.PublicKey().String()).
			Str("target_address", trgtWllt.PublicKey().String()).
			Err(err).
			Msg("query 오류 발생")
		return err
	}

	return c.Status(fiber.StatusOK).
		JSON(
			parameters.NewSuccessResponse(
				parameters.TokenBalanceRes{
					MintAddress:  mintWl.PublicKey().String(),
					OwnerAddress: trgtWllt.PublicKey().String(),
					Balance:      *balance,
				},
			),
		)
}

// @Summary Token Query
// @Description Token Balance를 조회합니다
// @Tags /spl
// @Accept json
// @Produce json
// @Param body body parameters.TargetTokenBalanceReq true "Target Transfer SPL Token"
// @Success 200 {object} parameters.CommonResponse{data=parameters.TokenBalanceRes}
// @Router /spl/target/query [post]
func (h *SplHandler) TargetTokenBalance(c *fiber.Ctx) error {

	reqID := c.Locals(types.RequestID).(string)
	lg := h.logger.SetReqID(reqID)
	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

	req := &parameters.TargetTokenBalanceReq{}
	err := c.BodyParser(req)
	if err != nil {
		return err
	}

	lg.Debug().
		Msg("GET /spl/target/query 호출")

	mintAddr, _ := h.solm.PublicKeyFromAddr(req.MintAddress)
	trgtAddr, _ := h.solm.PublicKeyFromAddr(req.OwnerAddress)

	balance, err := h.solm.TokenBalance(ctx, mintAddr, trgtAddr)
	if err != nil {
		lg.Error().
			Str("mint_address", req.MintAddress).
			Str("target_address", req.OwnerAddress).
			Err(err).
			Msg("target query 오류 발생")
		return err
	}

	return c.Status(fiber.StatusOK).
		JSON(
			parameters.NewSuccessResponse(
				parameters.TargetBalanceRes{
					Balance: *balance,
				},
			),
		)
}

// @Summary Target Token Transfer
// @Description Token을 전송합니다.
// @Tags /spl
// @Accept json
// @Produce json
// @Param body body parameters.TargetTransferTokenReq true "Transfer SPL Token"
// @Success 200 {object} parameters.CommonResponse{data=parameters.TransactionRes}
// @Router /spl/target/transfer [post]
func (h *SplHandler) TargetTransferToken(c *fiber.Ctx) error {
	reqID := c.Locals(types.RequestID).(string)
	lg := h.logger.SetReqID(reqID)
	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

	req := &parameters.TargetTransferTokenReq{}
	err := c.BodyParser(req)
	if err != nil {
		return err
	}

	lg.Debug().
		Msg("POST /spl/target/transfer 호출")

	mintWl, _ := h.solm.WalletFromPK(req.MintWallet)
	sender, _ := h.solm.WalletFromPK(req.OwnerWallet)
	receiver, _ := h.solm.WalletFromPK(req.TargetWallet)

	sig, err := h.solm.TransferToken(ctx, sender, sender, receiver, mintWl.PublicKey(), uint64(req.Amount))
	if err != nil {
		lg.Error().
			Str("mint_address", mintWl.PublicKey().String()).
			Str("sender_address", sender.PublicKey().String()).
			Str("target_address", receiver.PublicKey().String()).
			Err(err).
			Msg("transfer 오류 발생")
		return err
	}

	return c.Status(fiber.StatusOK).
		JSON(
			parameters.NewSuccessResponse(
				parameters.TransactionRes{
					Signature: sig.String(),
				},
			),
		)
}
