package address

import (
	"fmt"
	"testing"
)

func TestCreateAddressFromPrivateKey(t *testing.T) {
	address, err := CreateAddressFromPrivateKey()
	if err != nil {
		return
	}
	fmt.Println(address.PrivateKey, address.PublicKey, address.Address)
}

// bec3229437866e1ac4bc5092c5fbba7bf155010516cdc50da6276fed1b52cdf5
//04fe032dac102b3cc0359d5dec61f11e302e342dbd4c6f6c2747e0c5be19f2c53c34e7bdd493a6f6763e6157f58552f3c8a2c6eec1f63333aa075983bc78a6c00e
//0xf066D46B124FBF820B5E3Bc08bE3A9F5c6Bbc513

func TestCreateAddressFromPublicKey(t *testing.T) {
	address, err := PublicKeyToAddress("04fe032dac102b3cc0359d5dec61f11e302e342dbd4c6f6c2747e0c5be19f2c53c34e7bdd493a6f6763e6157f58552f3c8a2c6eec1f63333aa075983bc78a6c00e")
	if err != nil {
		return
	}
	fmt.Println("address=", address)
}
