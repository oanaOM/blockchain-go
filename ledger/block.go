package ledger

//Block is a definition of the block part from a blockchain
type Block struct {
	Hash     string
	Previous *Block
	Next     *Block
}
