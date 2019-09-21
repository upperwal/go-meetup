package structs

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"time"
)

type Block struct {
	Nonce        int
	PreviousHash string
	Transactions []Transaction
	Timestamp    int64
}

func NewBlock(cnt int) *Block {
	b := &Block{
		Nonce:     cnt,
		Timestamp: time.Now().Unix(),
	}

	return b
}

func (b *Block) SetPrevHash(prevBlock Block) error {
	m, err := json.Marshal(prevBlock)
	if err != nil {
		return err
	}

	s := md5.Sum(m)
	b.PreviousHash = hex.EncodeToString(s[:])
	return nil
}

func (b *Block) AddTransactions(t []Transaction) {
	b.Transactions = t
}
