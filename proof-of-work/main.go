package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

func calculateHash(block Block) string {
	record := block.Timestamp + block.Data + block.PrevHash + strconv.Itoa(block.Nonce)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func mineBlock(data string, prevHash string, difficulty int) Block {
	var newBlock Block
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash = prevHash
	newBlock.Nonce = 0
	for {
		newBlock.Hash = calculateHash(newBlock)
		if strings.HasPrefix(newBlock.Hash, strings.Repeat("0", difficulty)) {
			break
		}
		newBlock.Nonce++
	}
	fmt.Println("newBlock", newBlock)
	return newBlock
}

func main() {
	genesis := mineBlock("Genesis Block", "", 4)
	fmt.Println("Genesis Block:")
	fmt.Printf("Hash: %s\nNonce: %d\n\n", genesis.Hash, genesis.Nonce)

	second := mineBlock("İkinci Block", genesis.Hash, 4)
	fmt.Println("İkinci Blok:")
	fmt.Printf("Hash: %s\nNonce: %d\n", second.Hash, second.Nonce)
}
