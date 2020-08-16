package main

import (
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/oanaOM/blockchain-go/ledger"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := ledger.Block{Index: 0, Timestamp: t.String(), BPM: 0, Hash: "", PreviousHash: ""}
		spew.Dump(genesisBlock)
		ledger.Blockchain = append(ledger.Blockchain, genesisBlock)

	}()

	log.Fatal(ledger.Run())
}
