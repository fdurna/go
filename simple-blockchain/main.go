package main

import (
	"crypto/sha256" // hash hesaplamak için kullanılır
	"encoding/hex"  // has'i okunabilir hala getirmek için kullanılır.
	"encoding/json" // json veri işleme
	"fmt"           // Yazdırmak için (fmt.PrintLn vs)
	"io"            // HTTP'den gelen veri okumak için
	"net/http"      // HTTP sunucusu kurmak için
	"strconv"       // Sayı -> string dönüşümü
	"time"          // Zaman bilgisini almak için
)

type Block struct {
	Index     int    `json:"index"`     // Blok sırası
	Timestamp string `json:"timestamp"` // Oluşturulma zamanı
	Data      string `json:"data"`      // İçerik
	PrevHash  string `json:"prev_hash"` // Önceki bloğun hash'i
	Hash      string `json:"hash"`      // Bu bloğun hash'i
}

// Tüm bloklar burada tutulur.
var blockchain []Block

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, data string) Block {
	newBlock := Block{
		Index:     oldBlock.Index + 1,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  oldBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func isBlockValid(newBlock, oldBlock Block) bool {
	return newBlock.Index == oldBlock.Index+1 &&
		newBlock.PrevHash == oldBlock.Hash &&
		newBlock.Hash == calculateHash(newBlock)
}

func getBlocksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blockchain)
}

func addBlockHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Data string `json:"data"`
	}
	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &body)

	lastBlock := blockchain[len(blockchain)-1]
	newBlock := generateBlock(lastBlock, body.Data)
	if isBlockValid(newBlock, lastBlock) {
		blockchain = append(blockchain, newBlock)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newBlock)
	} else {
		http.Error(w, "Geçersiz blok", http.StatusBadRequest)
	}
}

func main() {
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		PrevHash:  "",
	}
	genesisBlock.Hash = calculateHash(genesisBlock)
	blockchain = append(blockchain, genesisBlock)
	fmt.Println("blockchain", blockchain)
	http.HandleFunc("/blocks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getBlocksHandler(w, r)
		} else if r.Method == "POST" {
			addBlockHandler(w, r)
		}
	})
	fmt.Println("Sunucu : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	// genesisBlock.Hash = "123z412vad1esd123dqx"
}
