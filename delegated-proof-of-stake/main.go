package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Voter struct {
	Name  string
	Stake int
}

type Delegate struct {
	Name string
}

type Block struct {
	Index     int
	Timestamp string
	Data      string
	Producer  string
}

func electDelegates(voters []Voter, count int) []Delegate {
	sort.Slice(voters, func(i, j int) bool {
		return voters[i].Stake > voters[j].Stake
	})

	var delegates []Delegate
	for i := 0; i < count && i < len(voters); i++ {
		delegates = append(delegates, Delegate{Name: voters[i].Name})
	}
	return delegates
}

func createBlock(index int, data string, producer string) Block {
	return Block{
		Index:     index,
		Timestamp: time.Now().Format(time.RFC3339),
		Data:      data,
		Producer:  producer,
	}
}

func runDPoS() {
	voters := []Voter{
		{"Alice", 50},
		{"Bob", 30},
		{"Charlie", 20},
		{"Dave", 10},
	}

	delegates := electDelegates(voters, 3)

	fmt.Println("🗳️ Seçilen Delegeler (En çok stake edenler):")
	for i, d := range delegates {
		fmt.Printf("%d. %s\n", i+1, d.Name)
	}

	var blockchain []Block

	fmt.Println("\n⛏️ Blok Üretimi Başlıyor:")
	for i := 0; i < 6; i++ {
		producer := delegates[i%len(delegates)].Name
		newBlock := createBlock(i, fmt.Sprintf("İşlem #%d", i), producer)
		blockchain = append(blockchain, newBlock)

		fmt.Printf("Blok %d | Üreten: %s | Veri: %s | Zaman: %s\n",
			newBlock.Index, newBlock.Producer, newBlock.Data, newBlock.Timestamp)
	}

	fmt.Println("\n📦 Final Blockchain:")
	for _, block := range blockchain {
		fmt.Printf("Blok %d - %s tarafından üretildi - %s\n", block.Index, block.Producer, block.Data)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	runDPoS()
}
