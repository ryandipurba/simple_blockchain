package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	timestamp    time.Time
	transactions []string
	prevhash     []byte
	Hash         []byte
}

func main() {
	ryan := []string{"ryandi sent 50 BTC to Reza"}
	block_1 := Blocks(ryan, []byte{})
	fmt.Println("This is our first block")
	Print(block_1)
	Transaction(block_1)

	jack := []string{"jack send 20 BTC to Reynold"}
	block_2 := Blocks(jack, block_1.Hash)
	fmt.Println("This is our second block")
	Print(block_2)
	Transaction(block_2)
}

func Blocks(transactions []string, prevhash []byte) *Block {
	currentTime := time.Now()
	return &Block{
		timestamp:    time.Time{},
		transactions: transactions,
		prevhash:     prevhash,
		Hash:         NewHash(currentTime, transactions, prevhash),
	}
}

func NewHash(time time.Time, transactions []string, prevhash []byte) []byte {
	input := append(prevhash, time.String()...)
	for transaction := range transactions {
		input = append(input, string(rune(transaction))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

func Print(block *Block) {
	fmt.Printf("\ttime: %s\n", block.timestamp.String())
	fmt.Printf("\tprevhash: %x\n", block.prevhash)
	fmt.Printf("\thash: %x\n", block.Hash)
}
func Transaction(block *Block) {
	fmt.Println("\tTransactions:")
	for i, transaction := range block.transactions {
		fmt.Printf("\t\t%v: %q\n", i, transaction)
	}
}
