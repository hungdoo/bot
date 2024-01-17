package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/hungdoo/bot/src/packages/command/tomb"
	"github.com/hungdoo/bot/src/packages/tombplus"
)

func TestLoadSecretFile(t *testing.T) {
	pk, err := tomb.LoadSecrets(0, "16-byte-key12345")
	if err != nil {
		log.Fatal("LoadSecret error:", err)
	}
	addess, err := tombplus.AddressFromPriKey(pk)
	if err != nil {
		log.Fatal("AddressFromPriKey error:", err)
	}
	t.Log(addess)
	if addess.String() != "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" {
		t.Errorf("Not same address")
	}

	pk, err = tomb.LoadSecrets(1, "16-byte-key12399")
	if err != nil {
		log.Fatal("LoadSecret error:", err)
	}
	addess, err = tombplus.AddressFromPriKey(pk)
	if err != nil {
		log.Fatal("AddressFromPriKey error:", err)
	}
	if addess.String() != "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" {
		t.Errorf("Not same address")
	}
}
func TestSecret(t *testing.T) {
	// Example plaintext secret
	plaintextSecrets := map[string]string{ // [16-byte-key12345]:"k"
		"16-byte-key12345": "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		"16-byte-key12399": "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
	}

	for k, s := range plaintextSecrets {
		// Encrypt the secret
		encryptedSecret, err := tomb.Encrypt([]byte(s), []byte(k))
		if err != nil {
			log.Fatal("Encryption error:", err)
		}

		// Print the encrypted secret (this is what you would store or transmit)
		fmt.Println("Encrypted Secret:", encryptedSecret)

		// Decrypt the secret
		decryptedSecret, err := tomb.Decrypt(encryptedSecret, []byte(k))
		if err != nil {
			log.Fatal("Decryption error:", err)
		}

		// Print the decrypted secret
		fmt.Println("Decrypted Secret:", string(decryptedSecret))
	}
}
