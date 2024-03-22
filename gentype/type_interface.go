package gentype

import (
	"errors"

	"github.com/leenzstra/vanity_address_generator/types"
)

func SelectGen(t types.AddressType, p types.Params) (GenInterface, error) {
	switch t {
	case types.EvmAddr:
		return NewEvmGen(p), nil
	case types.SolanaAddr:
		return NewSolanaGen(p), nil
	case types.MoveAddr:
		return nil, errors.New("MOVE not implemented")
	default:
		return nil, errors.New("not exists")
	}

}

type GenInterface interface {
	Generate() ([]byte, []byte, error)
	Check(pub []byte) bool
	Encode(pub []byte, pk []byte) (string, string)
}
