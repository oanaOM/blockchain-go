//Package ledger implementes the double linked algorithm demonstrated on a blockchain
package ledger

import (
	"errors"
)

//Ledger represents the blockchain
type Ledger struct {
	Genesis *Block
}

//Add adds blocks on the chain and returns his index
func (l *Ledger) Add(h string) string {
	if l.Genesis == nil {
		l.Genesis = &Block{Hash: h, Next: nil, Previous: nil}
		return h
	}
	block := l.Genesis

	for block.Next != nil {
		block = block.Next
	}

	//we add the block to the end of the list
	block.Next = &Block{Hash: h, Previous: block, Next: nil}

	return h
}

//Get retrieves a block from the chain
func (l *Ledger) Get(h string) (*Block, error) {
	if l.Genesis == nil {
		return nil, errors.New("This chain is empty")
	}

	block := l.Genesis
	hash := l.Genesis.Hash

	for block != nil {
		if hash == h {
			return block, nil
		}

		block = block.Next
		hash = block.Hash
	}
	return nil, errors.New("Block not found")
}
