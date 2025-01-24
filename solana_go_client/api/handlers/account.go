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

type AccountHandler struct {
	solm   *solana.SolanaManager
	wm     *solana.WalletManager
	logger log.Logger
}

// todo. interface로 연결
func NewAccountHandler(solm *solana.SolanaManager, wm *solana.WalletManager) *AccountHandler {
	return &AccountHandler{
		solm:   solm,
		wm:     wm,
		logger: log.GetLogger("solana_handler"),
	}
}

func (h *AccountHandler) Append(r fiber.Router) {
	r.Post("/balance", h.SolBalance)
	r.Post("/fill-balance", middlewares.Validate(&parameters.AridropReq{}), h.FillBalance)

}

// @Summary solana balance query
// @Description SOL Balance를 조회합니다
// @Tags /account
// @Accept json
// @Produce json
// @Param body body parameters.TargetSolBalanceReq true "Get SOL Balance"
// @Success 200 {object} parameters.CommonResponse{data=parameters.TargetBalanceRes}
// @Router /account/balance [post]
func (h *AccountHandler) SolBalance(c *fiber.Ctx) error {

	reqID := c.Locals(types.RequestID).(string)
	lg := h.logger.SetReqID(reqID)
	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

	req := &parameters.TargetSolBalanceReq{}
	err := c.BodyParser(req)
	if err != nil {
		return err
	}

	lg.Debug().
		Msg("GET /spl/target_query 호출")

	trgtWllt, _ := h.solm.PublicKeyFromAddr(req.OwnerAddress)

	balance, err := h.solm.Balance(ctx, trgtWllt)
	if err != nil {
		lg.Error().
			Str("target_wallet", req.OwnerAddress).
			Err(err).
			Msg("target solana query 오류 발생")
		return err
	}

	return c.Status(fiber.StatusOK).
		JSON(
			parameters.NewSuccessResponse(
				parameters.TargetBalanceRes{
					Balance: balance.String(),
				},
			),
		)
}

// @Summary fill solana balance
// @Description SOL Balance를 충전합니다
// @Tags /account
// @Accept json
// @Produce json
// @Param body body parameters.AridropReq true "Get SOL Balance"
// @Success 200
// @Router /account/fill-balance [post]
func (h *AccountHandler) FillBalance(c *fiber.Ctx) error {

	reqID := c.Locals(types.RequestID).(string)
	lg := h.logger.SetReqID(reqID)
	ctx := context.WithValue(c.Context(), types.RequestID, reqID)

	req := &parameters.AridropReq{}
	err := c.BodyParser(req)
	if err != nil {
		return err
	}

	lg.Debug().
		Msg("POST /account/fill-balance 호출")

	n := h.wm.MintN()

	var wg sync.WaitGroup
	wg.Add(n)

	ch := make(chan error)
	for i := 0; i < n; i++ {
		go func(ch chan error) {
			defer wg.Done()
			_, initWllt := h.wm.NextMintInitWallet()

			amount, err := h.solm.Balance(ctx, initWllt.PublicKey())
			if err != nil {
				lg.Error().
					Str("address", initWllt.PublicKey().String()).
					Err(err).
					Msg("balance 조회 오류 발생")
				ch <- err
			}
			bal, _ := amount.Float32()

			if bal == 0 {
				_, err = h.solm.RequestAirdrop(ctx, initWllt, uint64(req.Amount))
				if err != nil {
					lg.Error().
						Str("address", initWllt.PublicKey().String()).
						Err(err).
						Msg("airdrop 오류 발생")
					ch <- err
				}
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

	lg.Debug().Msg("airdrop 완료")

	return c.Status(fiber.StatusOK).
		JSON(
			parameters.NewSuccessResponse(nil),
		)
}
