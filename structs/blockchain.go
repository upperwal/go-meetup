package structs

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/upperwal/go-meetup/utils"
)

type Blockchain struct {
	Blocks              []Block
	PendingTransactions []Transaction
	BlocksealDuration   time.Duration
	State               int
}

func NewBlockchain() *Blockchain {

	b := &Blockchain{
		Blocks: []Block{
			getGenesisBlock(),
		},
		PendingTransactions: []Transaction{},
		BlocksealDuration:   5 * time.Second,
	}

	go b.Sealer()

	return b
}

func getGenesisBlock() Block {
	return Block{
		Nonce:        0,
		Timestamp:    time.Now().Unix(),
		Transactions: nil,
		PreviousHash: "0000000000000000000000000000000000",
	}
}

func (bc *Blockchain) AddBlock(b Block) {
	bc.Blocks = append(bc.Blocks, b)
}

func (bc *Blockchain) AddPendingTrans(t Transaction) {
	bc.PendingTransactions = append(bc.PendingTransactions, t)
}

func (bc *Blockchain) Sealer() {
	for {
		time.Sleep(bc.BlocksealDuration)

		/* if len(bc.PendingTransactions) == 0 {
			continue
		} */

		for _, t := range bc.PendingTransactions {
			bc.State += t.Value
		}

		l := len(bc.Blocks)
		b := NewBlock(l)
		b.SetPrevHash(bc.Blocks[l-1])
		b.AddTransactions(bc.PendingTransactions)

		bc.Blocks = append(bc.Blocks, *b)

		bc.PendingTransactions = []Transaction{}

		fmt.Println("Sealed a new Block")
		utils.GetPrettyJSON(b)

		fmt.Println("New State:", bc.State)
	}
}

func (bc *Blockchain) Marshal() ([]byte, error) {
	return json.Marshal(bc.Blocks)
}

func (bc *Blockchain) GetState() string {
	return strconv.Itoa(bc.State)
}
