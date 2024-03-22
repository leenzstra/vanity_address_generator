package main

import (
	"context"

	"log"
	"os"

	"github.com/leenzstra/vanity_address_generator/gentype"
	"github.com/leenzstra/vanity_address_generator/multigen"
	"github.com/leenzstra/vanity_address_generator/types"
	flag "github.com/spf13/pflag"
)

func main() {
	seq := flag.StringP("seq", "s", "", "sequence to find")
	startPos := flag.IntP("seqPos", "p", 0, "start position of sequence in address")
	addrType := flag.StringP("type", "t", "", "address type [evm, solana, move]")
	caseSens := flag.BoolP("caseSensetive", "c", false, "is case sensetive")
	workers := flag.IntP("workers", "w", 10, "parallel workers count")
	logFile := flag.StringP("logFile", "f", "logs", "log file")
	generateCount := flag.IntP("generateCount", "g", 1, "address count to generate")
	flag.Parse()

	// pipe logs
	f, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	// get address type number
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

	// create generator
	gen, err := gentype.SelectGen(typeNum, p)
	if err != nil {
		log.Fatal(err)
	}

	// create parallel generator
	multigen := multigen.New(gen, *workers)

	toGenerate := *generateCount
	if toGenerate < 1 {
		log.Fatal("generateCount must bo > 0")
	}

	// start generation
	for keys := range multigen.Start(context.Background()) {
		log.Printf("pub: %s pk: %s\n", keys.PublicKey, keys.PrivateKey)

		toGenerate--
		if toGenerate == 0 {
			break
		}
	}

	log.Printf("end working")
}
