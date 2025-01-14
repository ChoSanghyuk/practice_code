package api

import (
	"context"
	"encoding/json"
	"fmt"
	"workspace/api/handlers"
	"workspace/api/middlewares"
	"workspace/pkg/solana"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	conf Config
	app  *fiber.App
}

func NewServerWith(conf Config,
	solm *solana.SolanaManager,
	wm *solana.WalletManager,
) (*Server, error) {

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		AppName:     conf.AppName,
	})

	router := app.Group("/api/v1")
	middlewares.SetupMiddlewares(router)

	spl := router.Group("spl")
	handlers.NewSolanaHandler(solm, wm).Append(spl)

	app.Get("/docs/*", swagger.HandlerDefault)
	router.Get("/docs/*", swagger.HandlerDefault)

	return &Server{
		conf: conf,
		app:  app,
	}, nil
}

func (s *Server) Start() error {
	return s.app.Listen(fmt.Sprintf(":%d", s.conf.Port))
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}
