package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"io"
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

func NewPrivateKeyFromSeed(seed []byte) *PrivateKey {
	if len(seed) != seedLen {
		panic("invalid seed length, must be 32")
	}

	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}
}

func NewPrivateKeyFromString(s string) *PrivateKey {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return NewPrivateKeyFromSeed(b)
}

const (
	privKeyLen = 64 // length of the private key
	pubKeyLen  = 32 // length of the public key
	seedLen    = 32
	addressLen = 20
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

// Address returns last 20 characters on the public address
func (p *PublicKey) Address() Address {
	return Address{
		value: p.key[len(p.key)-addressLen:],
	}
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

// Address
type Address struct {
	value []byte
}

func (a *Address) Bytes() []byte {
	return a.value
}

// String converts Address to string
func (a *Address) String() string {
	return hex.EncodeToString(a.value)
}
