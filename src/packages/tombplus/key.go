package tombplus

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func PrivateKeyFromHex(privateKeyHex string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(privateKeyHex)
}

func AddressFromPriKey(privateKeyECDSA *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("error casting public key to ECDSA")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}
