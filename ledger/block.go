package ledger

//Block is a definition of the block part from a blockchain
type Block struct {
	Index        int
	Timestamp    string
	BPM          int
	Hash         string
	PreviousHash string
}
