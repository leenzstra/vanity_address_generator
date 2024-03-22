package gentype

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/dustinxie/ecc"
	"golang.org/x/crypto/sha3"

	"github.com/leenzstra/vanity_address_generator/types"
)

var _ GenInterface = (*EvmGen)(nil)

type EvmGen struct {
	params types.Params
}

// Check implements GenInterface.
func (e *EvmGen) Check(pub []byte) bool {
	pubhex := hex.EncodeToString(pub)
	seq := e.params.Sequence

	if !e.params.CaseSensetive {
		pubhex = strings.ToLower(pubhex)
		seq = strings.ToLower(seq)
	}

	return strings.Index(pubhex, seq) == e.params.SeqPos
}

// Encode implements GenInterface.
func (e *EvmGen) Encode(pub []byte, pk []byte) (string, string) {
	pubhex := hex.EncodeToString(pub)
	pkhex := hex.EncodeToString(pk)

	return "0x" + pubhex, "0x" + pkhex
}

// Generate implements GenInterface.
func (e *EvmGen) Generate() ([]byte, []byte, error) {
	p256k1 := ecc.P256k1()
	priv, err := ecdsa.GenerateKey(p256k1, rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	addr, err := e.addrFromPrivate(priv)
	if err != nil {
		return nil, nil, err
	}

	return addr, priv.D.Bytes(), nil
}

func (e *EvmGen) addrFromPrivate(pk *ecdsa.PrivateKey) ([]byte, error) {
	// sign message
	hash := sha256.Sum256([]byte("gen"))
	sig, err := ecc.SignEthereum(hash[:], pk)
	if err != nil {
		return nil, err
	}

	pub, err := ecc.RecoverEthereum(hash[:], sig)
	if err != nil {
		return nil, err
	}

	k256 := sha3.NewLegacyKeccak256()
	k256.Write(pub[1:])
	addr := k256.Sum(nil)

	return addr[12:], nil
}

func NewEvmGen(p types.Params) *EvmGen {
	return &EvmGen{
		params: p,
	}
}
