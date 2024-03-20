package tomb

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hungdoo/bot/src/packages/tombplus"
)

func TestLoadSecretFile(t *testing.T) {
	pk, err := LoadSecrets(0, "16-byte-key12345")
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

	pk, err = LoadSecrets(1, "16-byte-key12399")
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
		encryptedSecret, err := Encrypt([]byte(s), []byte(k))
		if err != nil {
			log.Fatal("Encryption error:", err)
		}

		// Print the encrypted secret (this is what you would store or transmit)
		fmt.Println("Encrypted Secret:", encryptedSecret)

		// Decrypt the secret
		decryptedSecret, err := Decrypt(encryptedSecret, []byte(k))
		if err != nil {
			log.Fatal("Decryption error:", err)
		}

		// Print the decrypted secret
		fmt.Println("Decrypted Secret:", string(decryptedSecret))
	}
}

func TestTombCommand(t *testing.T) {
	cmd := TombCommand{
		Rpc:      "https://rpc3.fantom.network",
		Contract: "0xA979F47480b4B598bf6a8bFA73aC0B6aEccBa505",
		Up:       true,
		PkIdx:    1,
		Key:      "aaa",
		User:     "0xA979F47480b4B598bf6a8bFA73aC0B6aEccBa505",
	}

	// Set time
	res, err := cmd.Execute(true, "")
	if err != nil {
		t.Error(err)
	}
	t.Logf("set time: %v\nTimestamps: %v/%v", res, cmd.VoteEndTimestamp, cmd.NextEpochTimestamp)

	// Early
	res, err = cmd.Execute(true, "")
	if err != nil {
		t.Error(err)
	}
	t.Logf("early: %v", res)

	// Mock VoteEnd 4m before now -> vote too late
	cmd.VoteEndTimestamp = time.Now().Add(-4 * time.Minute)
	res, err = cmd.Execute(true, "")
	if err != nil {
		t.Error(err)
	}
	t.Logf("VoteEnd 4m before now: %v", res)

	// Mock VoteEnd 4m after now -> vote window
	cmd.VoteEndTimestamp = time.Now().Add(4 * time.Minute)
	res, err = cmd.Execute(true, "")
	if err != nil {
		t.Error(err)
	}
	t.Logf("VoteEnd 4m after now: %v", res)
}
