package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.Equal(t, len(privKey.Bytes()), privKeyLen)

	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), pubKeyLen)
}

func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	msg := []byte("foo bar baz")

	sig := privKey.Sign(msg)
	assert.True(t, sig.Verify(pubKey, msg))

	// returns false when invalid message is used
	assert.False(t, sig.Verify(pubKey, []byte("foo")))

	// test with invalid public key
	// generate a new private key
	privKey2 := GeneratePrivateKey()
	// generate a new public key
	pubKey2 := privKey2.Public()
	// attempt to verify pubKey2 using a signature used to sign with
	// different private key
	// this should return false
	assert.False(t, sig.Verify(pubKey2, msg))
}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	address := pubKey.Address()
	assert.Equal(t, addressLen, len(address.Bytes()))
}

func TestNewPrivateKeyFromString(t *testing.T) {
	var (
		seed       = "5dc9b15df07f5f6c393c29bed0803f6168fb9affd70ebcb7aa904d17abfdbb84"
		privKey    = NewPrivateKeyFromString(seed)
		addressStr = "101c748a9d2afa77b766beb52c88bc5228fc643a"
	)

	assert.Equal(t, privKeyLen, len(privKey.Bytes()))
	address := privKey.Public().Address()
	assert.Equal(t, addressStr, address.String())
}
