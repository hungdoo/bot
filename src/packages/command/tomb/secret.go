package tomb

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hungdoo/bot/src/packages/dotenv"
)

// Encrypt encrypts the plaintext using AES-GCM encryption.
func Encrypt(plaintext []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the ciphertext using AES-GCM decryption.
func Decrypt(ciphertext string, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(decodedCiphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext is too short")
	}

	nonce, cipher := decodedCiphertext[:nonceSize], decodedCiphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipher, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func LoadSecrets(idx int, key string) (*ecdsa.PrivateKey, error) {
	encryptedContent, err := os.ReadFile(dotenv.GetEnv("SECRET_FILE"))
	if err != nil {
		return nil, err
	}

	encryptedSecrets := strings.Split(string(encryptedContent), ",")
	if len(encryptedSecrets) <= idx {
		return nil, fmt.Errorf("encryptedSecrets out-of-bound: length[%v], idx[%v]", len(encryptedSecrets), idx)
	}

	// Convert the byte slice to a string (assuming the secret is text-based)
	decrypted, err := Decrypt(string(encryptedSecrets[idx]), []byte(key))
	if err != nil {
		return nil, err
	}

	pk, err := crypto.ToECDSA(decrypted)
	if err != nil {
		return nil, err
	}
	return pk, nil
}
