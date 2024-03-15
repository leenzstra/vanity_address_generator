package types

import "errors"

type AddressType int

const (
	evm AddressType = iota
	solana 
	move
)

func TypeNum(typeName string) (AddressType, error) {
	switch typeName {
	case "sol":
		return solana, nil
	case "evm":
		return evm, nil
	case "move":
		return move, nil
	default:
		return -1, errors.New("no such type " + typeName)
	}
}

type Params struct {
	Sequence      string
	SeqPos        int
	Type          AddressType
	CaseSensetive bool
}
