package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
	handlers.NewSplHandler(solm, wm).Append(spl)
	account := router.Group("account")
	handlers.NewAccountHandler(solm).Append(account)

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

func (s *Server) InitMinitAccount() error {
	body := map[string]int{
		"amount": 100000000,
	}
	bodyBytes, _ := json.Marshal(body)
	resp, err := http.Post(fmt.Sprintf("http://localhost:%d/api/v1/spl/set-mint-account", s.conf.Port), "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Printf("Init Request Response status: %s\n", resp.Status)
	return nil
}

func (s *Server) InitTokenAccount() error {
	resp, err := http.Post(fmt.Sprintf("http://localhost:%d/api/v1/spl/set-token-account", s.conf.Port), "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Printf("Init Request Response status: %s\n", resp.Status)
	return nil
}
