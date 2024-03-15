package gentype

import (
	"crypto/ed25519"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/leenzstra/vanity_address_generator/types"
)

var _ GenInterface = (*SolanaGen)(nil)

type SolanaGen struct {
	addrLen int
	params  types.Params
}

// Encode implements GenInterface.
func (s *SolanaGen) Encode(pub []byte, pk []byte) (string, string) {
	return base58.Encode(pub), base58.Encode(pk)
}

// Generate pubk and pk
func (s *SolanaGen) Generate() ([]byte, []byte, error) {
	pub, pk, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, nil, err
	}

	return pub, pk, nil
}

func (s *SolanaGen) Check(pub []byte) bool {
	pubb58 := base58.Encode(pub)
	seq := s.params.Sequence

	if !s.params.CaseSensetive {
		pubb58 = strings.ToLower(pubb58)
		seq = strings.ToLower(seq)
	}

	// log.Println(seq, pubb58)

	return strings.Index(pubb58, seq) == s.params.SeqPos
}

func NewSolanaGen(p types.Params) *SolanaGen {
	return &SolanaGen{
		addrLen: 44,
		params:  p,
	}
}
