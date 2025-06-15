package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Validator string
}

type Validator struct {
	Name  string
	Stake int
}

// hash hesapalama
func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.Data, block.PrevHash)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

// yeni blok oluşturma
func generateBlock(prevBlock Block, data string, validatorName string) Block {
	newBlock := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevBlock.PrevHash,
		Validator: validatorName,
	}
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

// Stake oranına göre validator seçme
func selectValidator(validators []Validator) Validator {
	var pool []Validator
	for _, v := range validators {
		for i := 0; i < v.Stake; i++ {
			pool = append(pool, v)
		}
	}
	rand.Seed(time.Now().UnixNano())
	return pool[rand.Intn(len(pool))]
}

func main() {
	genesis := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		PrevHash:  "",
		Validator: "System",
	}
	genesis.Hash = calculateHash(genesis)
	blockchain := []Block{genesis}
	validators := []Validator{
		{"Esra", 5},
		{"Fatih", 3},
		{"Veli", 2},
	}
	for i := 0; i < 5; i++ {
		chosen := selectValidator(validators)
		data := fmt.Sprintf("Transaction Data %d", i+1)
		newBlock := generateBlock(blockchain[len(blockchain)-1], data, chosen.Name)
		blockchain = append(blockchain, newBlock)
		fmt.Printf("Blok #%d %s tarafından eklendi (Stake: %d)\n", newBlock.Index, chosen.Name, chosen.Stake)
	}
	fmt.Println("\nBlok Zinciri:")
	for _, block := range blockchain {
		fmt.Printf("Blok #%d | Validator: %s | Veri: %s | Hash: %.10s...\n",
			block.Index, block.Validator, block.Data, block.Hash)
	}
}
