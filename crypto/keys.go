package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"io"
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

const (
	privKeyLen = 64 // length of the private key
	pubKeyLen  = 32 // length of the public key
	seedLen    = 32
)

func GeneratePrivateKey() *PrivateKey {
	seed := make([]byte, seedLen)
	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		panic(err)
	}

	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}

// Sign signs/authorizes transaction on the blockchain network
// using the private key and a message
// the message can be e.g. the transaction information in the case of bitcoin
func (p *PrivateKey) Sign(msg []byte) *Signature {
	return &Signature{value: ed25519.Sign(p.key, msg)}
}

type PublicKey struct {
	key ed25519.PublicKey
}

func (p *PrivateKey) Public() *PublicKey {
	b := make([]byte, pubKeyLen)
	copy(b, p.key[32:])

	return &PublicKey{
		key: b,
	}
}

func (p *PublicKey) Bytes() []byte {
	return p.key
}

type Signature struct {
	value []byte
}

func (s *Signature) Bytes() []byte {
	return s.value
}

func (s *Signature) Verify(pubKey *PublicKey, msg []byte) bool {
	return ed25519.Verify(pubKey.key, msg, s.value)
}
