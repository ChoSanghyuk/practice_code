package config

import (
	"strings"
	"workspace/api"
	"workspace/pkg/log"
	"workspace/pkg/solana"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/spf13/viper"
)

type Config struct {
	viper  viper.Viper
	logger log.Logger
}

func New(path, fileName, fileType string) (*Config, error) {
	lg := log.GetLogger("config")

	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		viper:  *viper.GetViper(),
		logger: lg,
	}, nil
}

func (c *Config) LogConfig() log.Config {
	path := "log"

	conf := log.Config{
		LogLevel: log.Level(c.viper.GetInt(path + ".level")),
	}

	c.logger.Info().
		Int("레벨 0: TRACE, 1: DEBUG, 2: INFO, 3: WARN, 4: ERROR, 5: FATAL, 6: PANIC, 7: NOLEVEL, 8: DISABLED", int(conf.LogLevel)).
		Msg("Log Config")

	return conf
}

func (c *Config) APIConfig() api.Config {
	path := "api"

	conf := api.Config{
		AppName: c.viper.GetString(path + ".app-name"),
		Port:    c.viper.GetUint(path + ".port"),
		Version: c.viper.GetString(path + ".version"),
	}

	c.logger.Info().
		Str("어플리케이션 이름", conf.AppName).
		Uint("어플리케이션 포트", conf.Port).
		Str("어플리케이션 버전", conf.Version).
		Msg("API Config")

	return conf
}

func (c *Config) SolanaConfig() solana.SolManagerConfig {
	path := "solana"

	conf := solana.SolManagerConfig{
		RPCURL:     c.viper.GetString(path + ".rpc_url"),
		WSURL:      c.viper.GetString(path + ".ws_url"),
		Commitment: rpc.CommitmentType(c.viper.GetString(path + ".commitment")),
		IsSync:     c.viper.GetBool(path + ".sync"),
	}

	c.logger.Info().
		Msg("Solana Config")

	return conf
}

func (c *Config) WalletConfig() solana.WalletManagerConfig {
	path := "wallet"

	n := c.viper.GetInt(path + ".N")
	m := c.viper.GetInt(path + ".M")

	return solana.WalletManagerConfig{
		N: solana.MintAccountNumber(n),
		M: solana.TargetAccountNumber(m),
	}
}
