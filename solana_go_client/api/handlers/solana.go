package handlers

import (
	"context"
	"sync"
	"workspace/api/middlewares"
	"workspace/api/parameters"
	"workspace/api/types"
	"workspace/pkg/log"
	"workspace/pkg/solana"

	"github.com/gofiber/fiber/v2"
)

type SolanaHandler struct {
	solm   *solana.SolanaManager
	wm     *solana.WalletManager
	logger log.Logger
}

func NewSolanaHandler(solm *solana.SolanaManager, wm *solana.WalletManager) *SolanaHandler {
	return &SolanaHandler{
		solm:   solm,
		wm:     wm,
		logger: log.GetLogger("solana_handler"),
	}
}

func (h *SolanaHandler) Append(r fiber.Router) {
	r.Post("/deploy", h.DeployToken)
	r.Post("/deploy-with-mint", middlewares.Validate(&parameters.CreateTokenWithMintReq{}), h.DeployWithMint)
	r.Post("/set-mint-account", middlewares.Validate(&parameters.CreateTokenWithMintReq{}), h.SetMintAccount)
	r.Post("/mint", middlewares.Validate(&parameters.PerformanceMintReq{}), h.Mint)
	r.Post("/transfer", middlewares.Validate(&parameters.TransferTokenReq{}), h.TransferToken)
	r.Get("/query", h.TokenBalance)
}

// @Summary SPL Token 생성
// @Description SPL Token을 생성합니다
// @Tags /spl
// @Accept json
// @Produce json
// @Success 200 {object} parameters.CommonResponse{data=parameters.PerformanceCreateSPLTokenRes}
// @Router /spl/deploy [post]
func (h *SolanaHandler) DeployToken(c *fiber.Ctx) error {
	reqID := c.Locals(types.RequestID).(string)
	lg := h.logger.SetReqID(reqID)
	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

	lg.Debug().
		Msg("POST /spl 호출")

	_, initW := h.wm.NextMintInitWallet()

	mintWallet := h.solm.CreateAccount(context.Background())
	sig, err := h.solm.SetMintAccount(ctx, mintWallet, initW, initW.PublicKey(), initW.PublicKey())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).
		JSON(
			parameters.NewSuccessResponse(
				parameters.PerformanceCreateSPLTokenRes{
					Signature:   sig.String(),
					MintAccount: mintWallet.PrivateKey.String(),
				},
			),
		)
}

// @Summary Token Account 생성 후 민팅
// @Description Token Account를 생성한 뒤 민팅합니다.
// @Tags /spl
// @Accept json
// @Produce json
// @Param body body parameters.CreateTokenWithMintReq true "Performance Create SPL Token And Mint Request"
// @Success 200 {object} parameters.CommonResponse{data=parameters.PerformanceCreateSPLTokenRes}
// @Router /spl/deploy-with-mint [post]
func (h *SolanaHandler) DeployWithMint(c *fiber.Ctx) error {
	reqID := c.Locals(types.RequestID).(string)
	lg := h.logger.SetReqID(reqID)
	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

	req := &parameters.CreateTokenWithMintReq{}
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
		return err
	}

	return c.Status(fiber.StatusOK).
		JSON(
			parameters.NewSuccessResponse(
				parameters.PerformanceCreateSPLTokenRes{
					Signature:   sig.String(),
					MintAccount: mintWallet.PrivateKey.String(),
				},
			),
		)
}

// @Summary 테스트 사전 준비
// @Description 기존 생성한 Mint Account에 deploy 후 민팅
// @Tags /spl
// @Accept json
// @Produce json
// @Param body body parameters.CreateTokenWithMintReq true "Performance Create SPL Token And Mint Request"
// @Success 200 {object} parameters.CommonResponse{data=parameters.PerformanceMintRes}
// @Router /spl/set-mint-account [post]
func (h *SolanaHandler) SetMintAccount(c *fiber.Ctx) error {
	reqID := c.Locals(types.RequestID).(string)
	lg := h.logger.SetReqID(reqID)
	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

	req := &parameters.CreateTokenWithMintReq{}
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

// @Summary 민팅
// @Description 배포 완료 상태에서 target wallet에 추가 민팅합니다.
// @Tags /spl
// @Accept json
// @Produce json
// @Param body body parameters.PerformanceMintReq true "Performance Create SPL Token Request"
// @Success 200 {object} parameters.CommonResponse{data=parameters.PerformanceMintRes}
// @Router /spl/mint [post]
func (h *SolanaHandler) Mint(c *fiber.Ctx) error {
	reqID := c.Locals(types.RequestID).(string)
	lg := h.logger.SetReqID(reqID)
	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

	req := &parameters.PerformanceMintReq{}
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
			Str("mint_wallet", mintWllt.PublicKey().String()).
			Str("init_wallet", initWllt.PublicKey().String()).
			Str("target_wallet", trgtWllt.PublicKey().String()).
			Err(err).
			Msg("Mint 오류 발생")
		return err
	}

	return c.Status(fiber.StatusOK).
		JSON(
			parameters.NewSuccessResponse(
				parameters.PerformanceMintRes{
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
// @Param body body parameters.TransferTokenReq true "Transfer SPL Token"
// @Success 200 {object} parameters.CommonResponse{data=parameters.PerformanceMintRes}
// @Router /spl/transfer [post]
func (h *SolanaHandler) TransferToken(c *fiber.Ctx) error {
	reqID := c.Locals(types.RequestID).(string)
	lg := h.logger.SetReqID(reqID)
	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

	req := &parameters.TransferTokenReq{}
	err := c.BodyParser(req)
	if err != nil {
		return err
	}

	lg.Debug().
		Msg("POST /spl/mint 호출")

	mintWl, sender := h.wm.NextMintInitWallet()
	receiver := h.wm.NextTrgtWallet()

	sig, err := h.solm.TransferToken(ctx, sender, sender, receiver, mintWl.PublicKey(), uint64(req.Amount))
	if err != nil {
		lg.Error().
			Str("mint_wallet", mintWl.PublicKey().String()).
			Str("sender_wallet", sender.PublicKey().String()).
			Str("target_wallet", receiver.PublicKey().String()).
			Err(err).
			Msg("transfer 오류 발생")
		return err
	}

	return c.Status(fiber.StatusOK).
		JSON(
			parameters.NewSuccessResponse(
				parameters.PerformanceMintRes{
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
func (h *SolanaHandler) TokenBalance(c *fiber.Ctx) error {

	reqID := c.Locals(types.RequestID).(string)
	lg := h.logger.SetReqID(reqID)
	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

	lg.Debug().
		Msg("GET /spl/query 호출")

	mintWl, _ := h.wm.NextMintInitWallet()
	trgtWllt := h.wm.NextTrgtWallet()

	balance, err := h.solm.TokenBalance(ctx, mintWl.PublicKey(), trgtWllt.PublicKey())
	if err != nil {
		lg.Error().
			Str("mint_wallet", mintWl.PublicKey().String()).
			Str("target_wallet", trgtWllt.PublicKey().String()).
			Err(err).
			Msg("query 오류 발생")
		return err
	}

	return c.Status(fiber.StatusOK).
		JSON(
			parameters.NewSuccessResponse(
				parameters.TokenBalanceRes{
					MintAccount:  mintWl.PrivateKey.String(),
					OwnerAccount: trgtWllt.PrivateKey.String(),
					Balance:      *balance,
				},
			),
		)

}

// // @Summary 전송
// // @Description 배포 및 민팅 완료 된 상태에서 transfer 진행
// // @Tags /spl
// // @Accept json
// // @Produce json
// // @Param body body parameters.CreateTokenWithMintReq true "Performance Create SPL Token And Mint Request"
// // @Success 200 {object} parameters.CommonResponse{data=parameters.PerformanceMintRes}
// // @Router /sol/performance/set-transfer [post]
// func (h *SolanaHandler) SetTransfer(c *fiber.Ctx) error {
// 	reqID := c.Locals(types.RequestID).(string)
// 	lg := h.logger.SetReqID(reqID)
// 	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

// 	req := c.Locals("validatedData").(*parameters.CreateTokenWithMintReq)

// 	lg.Debug().
// 		Msg("POST /sol/performance/mint 호출")

// 	owner := h.wm.NextUserWallet()
// 	payerWallet := owner
// 	auth := payerWallet.PublicKey()
// 	mintWallet := h.wm.NewWallet()

// 	sig, _, err := h.solm.SetMintAccountAndMint(ctx, payerWallet, mintWallet, owner.PublicKey(), auth, auth, uint64(req.Amount))
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(
// 			parameters.NewErrorResponse(
// 				err.Error(),
// 			),
// 		)
// 	}

// 	return c.Status(fiber.StatusOK).
// 		JSON(
// 			parameters.NewSuccessResponse(
// 				parameters.PerformanceMintRes{
// 					Signature: sig.String(),
// 				},
// 			),
// 		)
// }
