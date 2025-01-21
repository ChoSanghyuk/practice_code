package handlers

import (
	"context"
	"workspace/api/parameters"
	"workspace/api/types"
	"workspace/pkg/log"
	"workspace/pkg/solana"

	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	solm   *solana.SolanaManager
	logger log.Logger
}

func NewAccountHandler(solm *solana.SolanaManager) *AccountHandler {
	return &AccountHandler{
		solm:   solm,
		logger: log.GetLogger("solana_handler"),
	}
}

func (h *AccountHandler) Append(r fiber.Router) {
	r.Post("/balance", h.SolBalance)

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
