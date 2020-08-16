//Package ledger implementes the double linked algorithm demonstrated on a blockchain
package ledger

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Blockchain initialise a blockchain
var Blockchain []Block

//calculateHash creates a HASH256 string has of a block
func calculateHash(b Block) string {
	record := string(b.Index) + b.Timestamp + string(b.BPM) + b.PreviousHash

	h := sha256.New()

	h.Write([]byte(record))

	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

//CreateBlock will create a new block
func CreateBlock(b Block, bpm int) (Block, error) {

	var newB Block

	t := time.Now()

	newB.Index = b.Index + 1
	newB.Timestamp = t.String()
	newB.BPM = bpm
	newB.Hash = calculateHash(newB)
	newB.PreviousHash = b.Hash

	return newB, nil

}

// isValidBlock will check the block hasn't been tampered
func isValidBlock(newB Block, prevB Block) bool {
	if prevB.Index+1 != newB.Index {
		return false
	}

	if prevB.Hash != newB.PreviousHash {
		return false
	}

	if calculateHash(newB) != newB.Hash {
		return false
	}

	return true
}

// ReplaceChain will compare the length of the chain.
// when dealing with 2 blockchains usually the longest one gets picked
func ReplaceChain(newB []Block) {
	if len(newB) > len(Blockchain) {
		Blockchain = newB
	}
}
