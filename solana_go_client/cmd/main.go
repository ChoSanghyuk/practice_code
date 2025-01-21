package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"workspace/api"
	_ "workspace/api/docs"
	"workspace/config"
	"workspace/pkg/log"
	"workspace/pkg/solana"
)

// @title SOL API
// @version 1.0.0
// @description This is a rest api server for solapi
// @BasePath /api/v1
func main() {
	lg := log.GetLogger("workspace")

	defer func() {
		if r := recover(); r != nil {
			lg.Error().Msgf("어플리케이션 패닉 발생: %v\n", r)
		}
	}()

	conf, err := loadConfig()
	if err != nil {
		lg.Fatal().Err(err).Msg("Config 초기화 실패")
	}

	log.Configure(conf.LogConfig())
	////////////////////////////////////////////////////
	// 패키지 초기화
	solm, err := solana.NewSolanaManager(conf.SolanaConfig())
	if err != nil {
		lg.Fatal().Err(err).Msg("Solana Manager 초기화 실패")
	}

	path, _ := os.Getwd()
	wm, err := solana.NewWalletManager(path+"/wallets.txt", solm, conf.WalletConfig())
	if err != nil {
		lg.Fatal().Err(err).Msg("Wallet Manager 초기화 실패")
	}

	// err = logecordWallet(wm.AllAddress())
	// if err != nil {
	// 	lg.Fatal().Err(err).Msg("Wallet list 저장 실패")
	// }
	////////////////////////////////////////////////////
	// API서버 초기화
	server, err := api.NewServerWith(conf.APIConfig(), solm, wm)
	if err != nil {
		lg.Fatal().Err(err).Msg("API서버 초기화 실패")
	}

	////////////////////////////////////////////////////
	// API서버 실행
	go func() {
		defer func() {
			if r := recover(); r != nil {
				lg.Error().Msgf("API 서버 패닉 발생: %v", r)
			}
		}()

		if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			lg.Fatal().Err(err).Msg("API 서버 실행 실패")
		}
	}()
	go func() {
		time.Sleep(1 * time.Second)
		if err := server.Init_Wallet(); err != nil {
			lg.Fatal().Err(err).Msg("API 서버 초기화 실패")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-quit

		lg.Info().Str("signal", sig.String()).Msg("OS Signal 수신")

		timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer timeoutCancel()
		if err := server.ShutDown(timeoutCtx); err != nil {
			lg.Error().Err(err).Msg("API 서버 종료 실패")
		}
		os.Exit(0)

	}()

	select {}

}

func loadConfig() (*config.Config, error) {
	var configFileName, configFilePath string

	flag.StringVar(&configFilePath, "conf-path", "config/local", "Config File Path")
	flag.StringVar(&configFileName, "conf-file", "config", "Config File Name")
	flag.Parse()

	conf, err := config.New(configFilePath, configFileName, "yaml")
	if err != nil {
		return nil, err
	}
	fmt.Printf("config file : %s/%s\n", configFilePath, configFileName)
	return conf, nil
}
