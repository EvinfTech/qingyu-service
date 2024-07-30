package util

import (
	"crypto"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"math/big"
	"strings"
)

// ECDH The main interface for ECDH key exchange.
type ECDH interface {
	GenerateKey(io.Reader) (crypto.PrivateKey, crypto.PublicKey, error)
	Marshal(crypto.PublicKey) []byte
	MarshalA(crypto.PublicKey) string
	Unmarshal([]byte) (crypto.PublicKey, bool)
	UnmarshalA(string) (crypto.PublicKey, bool)
	GenerateSharedSecret(crypto.PrivateKey, crypto.PublicKey) ([]byte, error)
}

type ellipticECDH struct {
	ECDH
	curve elliptic.Curve
}

type ellipticPublicKey struct {
	elliptic.Curve
	X, Y *big.Int
}

type ellipticPrivateKey struct {
	D []byte
}

// NewEllipticECDH creates a new instance of ECDH with the given elliptic.Curve curve
// to use as the elliptical curve for elliptical curve diffie-hellman.
func NewEllipticECDH(curve elliptic.Curve) ECDH {
	return &ellipticECDH{
		curve: curve,
	}
}

func (e *ellipticECDH) GenerateKey(rand io.Reader) (crypto.PrivateKey, crypto.PublicKey, error) {
	var d []byte
	var x, y *big.Int
	var priv *ellipticPrivateKey
	var pub *ellipticPublicKey
	var err error

	d, x, y, err = elliptic.GenerateKey(e.curve, rand)
	if err != nil {
		return nil, nil, err
	}

	priv = &ellipticPrivateKey{
		D: d,
	}
	pub = &ellipticPublicKey{
		Curve: e.curve,
		X:     x,
		Y:     y,
	}

	return priv, pub, nil
}

func (e *ellipticECDH) Marshal(p crypto.PublicKey) []byte {
	pub := p.(*ellipticPublicKey)
	return elliptic.Marshal(e.curve, pub.X, pub.Y)
}

func (e *ellipticECDH) MarshalA(p crypto.PublicKey) string {
	pub := p.(*ellipticPublicKey)
	xhex := "0x" + hex.EncodeToString(pub.X.Bytes())
	yhex := "0x" + hex.EncodeToString(pub.Y.Bytes())
	return xhex + "," + yhex
}

func (e *ellipticECDH) Unmarshal(data []byte) (crypto.PublicKey, bool) {
	var key *ellipticPublicKey
	var x, y *big.Int

	x, y = elliptic.Unmarshal(e.curve, data)
	if x == nil || y == nil {
		return key, false
	}
	key = &ellipticPublicKey{
		Curve: e.curve,
		X:     x,
		Y:     y,
	}
	return key, true
}

func (e *ellipticECDH) UnmarshalA(data string) (crypto.PublicKey, bool) {
	var key *ellipticPublicKey
	keyArr := strings.Split(data, ",")
	xhex := keyArr[0][2:]
	yhex := keyArr[1][2:]
	keyhex := "04" + xhex + yhex
	keybytes, err := hex.DecodeString(keyhex)
	if err != nil {
		return key, false
	}
	return e.Unmarshal(keybytes)
}

// GenerateSharedSecret takes in a public key and a private key
// and generates a shared secret.
//
// RFC5903 Section 9 states we should only return x.
func (e *ellipticECDH) GenerateSharedSecret(privKey crypto.PrivateKey, pubKey crypto.PublicKey) ([]byte, error) {
	priv := privKey.(*ellipticPrivateKey)
	pub := pubKey.(*ellipticPublicKey)

	x, y := e.curve.ScalarMult(pub.X, pub.Y, priv.D)
	xhex := "0x" + hex.EncodeToString(x.Bytes())
	yhex := "0x" + hex.EncodeToString(y.Bytes())
	seed := xhex + yhex
	w := sha256.New()
	io.WriteString(w, seed)
	bw := w.Sum(nil)
	return bw, nil
}
