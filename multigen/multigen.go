package multigen

import (
	"context"
	"log"
	"sync"

	"github.com/alitto/pond"
	"github.com/leenzstra/vanity_address_generator/gentype"
	"github.com/leenzstra/vanity_address_generator/types"
)

type Multigen struct {
	gen gentype.GenInterface

	mut     sync.Mutex
	pool    *pond.WorkerPool
	workers int
}

func New(gen gentype.GenInterface, workers int) *Multigen {
	return &Multigen{
		gen:     gen,
		workers: workers,
	}
}

func (m *Multigen) Start(ctx context.Context) <-chan types.Keypair {
	m.pool = pond.New(m.workers, m.workers*10, pond.MinWorkers(m.workers))

	keypair := make(chan types.Keypair)
	errors := make(chan error)

	ctx, cancel := context.WithCancel(ctx)

	// catch errors
	go func() {
		select {
		case <-ctx.Done():
			return
		case err, _ := <-errors:
			cancel()
			log.Println(err)
		}
	}()

	// generate addresses
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("done")
				return
			default:
				m.pool.Submit(func() {
					pub, pk, err := m.gen.Generate()
					if err != nil {
						errors <- err
					}

					if m.gen.Check(pub) {
						spub, spk := m.gen.Encode(pub, pk)
						keypair <- types.Keypair{
							PublicKey:  spub,
							PrivateKey: spk,
						}
					}
				})
			}
		}
	}()

	log.Println("start working")

	return keypair
}
