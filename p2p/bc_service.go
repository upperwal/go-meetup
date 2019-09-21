package p2p

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	inet "github.com/libp2p/go-libp2p-core/network"
	service "github.com/pravahio/go-mesh/interface/service"
	"github.com/upperwal/go-meetup/structs"
	"github.com/upperwal/go-meetup/utils"
)

// Name and Version of this service.
const (
	NAME    = "blockchain"
	VERSION = 1
	TRANS   = "TRANS_TOPIC"
)

type BCService struct {
	service.ApplicationContainer

	BlockchainStore *structs.Blockchain
}

func NewBCService() *BCService {
	ss := &BCService{
		BlockchainStore: structs.NewBlockchain(),
	}

	ss.SetNameVersion(NAME, VERSION)
	go ss.IncomingTrans()

	return ss
}

func (bc *BCService) SendTransaction(t structs.Transaction) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	err = bc.GetPubSub().Publish(TRANS, b)
	if err != nil {
		return err
	}
	bc.BlockchainStore.AddPendingTrans(t)
	fmt.Println("\nNew Transaction")
	fmt.Println(utils.GetPrettyJSON(t))
	return nil
}

func (bc *BCService) IncomingTrans() {
	time.Sleep(3 * time.Second)
	fmt.Println("Ready")
	s, err := bc.GetPubSub().Subscribe(TRANS)
	if err != nil {
		panic(err)
	}

	for {
		m, err := s.Next(context.Background())
		if err != nil {
			panic(err)
		}

		if m.GetFrom() == bc.GetHost().ID() {
			continue
		}

		var t structs.Transaction
		err = json.Unmarshal(m.GetData(), &t)
		if err != nil {
			panic(err)
		}
		bc.BlockchainStore.AddPendingTrans(t)

		fmt.Println("Got new trans from remote")
		fmt.Println(utils.GetPrettyJSON(t))
	}

}

func (bc *BCService) MarshalBlockchain() ([]byte, error) {
	return bc.BlockchainStore.Marshal()
}

func (bcService *BCService) Start(ctx context.Context) error           { return nil }
func (bcService *BCService) Stop() error                               { return nil }
func (bcService *BCService) Run(stream inet.Stream)                    {}
func (bcService *BCService) Get(name string) (chan interface{}, error) { return nil, nil }
func (bcService *BCService) Set(name string, value interface{}) error  { return nil }
