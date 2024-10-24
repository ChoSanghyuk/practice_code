package crypto

import (
	"fmt"
	"testing"
)

func TestCrypt(t *testing.T) {
	key := []byte("") // 16, 24, 32 bytes key for AES-128, AES-192, AES-256
	text := ""

	// 암호화
	encrypted, err := encrypt(key, []byte(text))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted: %s\n", encrypted)

	// 복호화
	decrypted, err := decrypt(key, encrypted)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}
