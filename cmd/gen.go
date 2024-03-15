package main

import (
	"log"
	"os"

	"github.com/leenzstra/vanity_address_generator/gentype"
	"github.com/leenzstra/vanity_address_generator/types"
	flag "github.com/spf13/pflag"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("err: no sequence provided")
	}

	seq := flag.StringP("seq", "s", "", "sequence to find")
	startPos := flag.IntP("seqpos", "p", 0, "start position of sequence in address")
	addrType := flag.StringP("type", "t", "", "address type [evm, solana, move]")
	caseSens := flag.BoolP("casesense", "c", false, "is case sensetive")
	flag.Parse()

	typeNum, err := types.TypeNum(*addrType)
	if err != nil {
		log.Fatal(err)
	}

	p := types.Params{
		Sequence:      *seq,
		SeqPos:        *startPos,
		Type:          typeNum,
		CaseSensetive: *caseSens,
	}

	gen := gentype.NewSolanaGen(p)

	for {
		pub, pk, err := gen.Generate()
		if err != nil {
			log.Fatal(err)
		}

		if gen.Check(pub) {
			spub, spk := gen.Encode(pub,pk)
			log.Printf("generated: %s %s", spub, spk)
			return
		}
	}
}
