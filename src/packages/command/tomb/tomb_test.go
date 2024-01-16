package tomb

import (
	"fmt"
	"log"
	"testing"
)

func TestSecret(t *testing.T) {
	// Example plaintext secret
	plaintextSecrets := map[string]string{ // [16-byte-key12345]:"k"
		"16-byte-key12345": "mysecret1",
		"16-byte-key12399": "mysecret2",
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
