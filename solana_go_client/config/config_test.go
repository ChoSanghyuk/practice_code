package config

import (
	"log"
	"testing"

	"github.com/test-go/testify/require"
)

func TestWalletConfig(t *testing.T) {

	conf, err := New("local", "config", "yaml")
	require.NoError(t, err)

	config := conf.WalletConfig()

	log.Printf("Wallet Manager Config %v", config)
}
