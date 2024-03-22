package types

import "errors"

type AddressType int

const (
	EvmAddr AddressType = iota
	SolanaAddr 
	MoveAddr
)

func TypeNum(typeName string) (AddressType, error) {
	switch typeName {
	case "sol":
		return SolanaAddr, nil
	case "evm":
		return EvmAddr, nil
	case "move":
		return MoveAddr, nil
	default:
		return -1, errors.New("no type provided or no such type " + typeName)
	}
}

type Params struct {
	Sequence      string
	SeqPos        int
	Type          AddressType
	CaseSensetive bool
	Workers int
}

type Keypair struct {
	PublicKey  string
	PrivateKey string
}
