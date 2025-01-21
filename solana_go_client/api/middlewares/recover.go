package middlewares

import (
	"fmt"
	"workspace/api/parameters"
	"workspace/api/types"
	"workspace/pkg/log"

	"github.com/gofiber/fiber/v2"
)

func globalRecover(c *fiber.Ctx) error {
	lg := log.GetLogger("recover-middleware").SetReqID(c.Locals(types.RequestID).(string))

	defer func() {
		if r := recover(); r != nil {
			lg.Error().Msgf("Recover 미들웨어가 실행되었습니다. [%v]", r)
			c.Status(fiber.StatusInternalServerError).JSON(
				parameters.NewErrorResponse(
					fmt.Sprintf("서버패닉오류발생 [%v]", r),
				))
		}
	}()

	return c.Next()
}
