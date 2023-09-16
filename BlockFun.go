package Blockchain



import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
	"strings"
)

type Block struct {
    Transaction  string
    Nonce        int
    PreviousHash string
    Hash         string
}

var Blockchain []Block

func NewBlock(transaction string, nonce int, previousHash string) *Block {
    block := &Block{
        Transaction:  transaction,
        Nonce:        nonce,
        PreviousHash: previousHash,
    }
    block.Hash = CalculateHash(block)
    Blockchain = append(Blockchain, *block)
    return block
}

func ListBlocks() {
    for i, block := range Blockchain {
        fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 25), i , strings.Repeat("=", 25))
        fmt.Printf("Transaction: %s\n", block.Transaction)
        fmt.Printf("Nonce: %d\n", block.Nonce)
        fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
        fmt.Printf("Hash: %s\n", block.Hash)
        fmt.Println()
    }
}

func ChangeBlock(blockIndex int, newTransaction string) {
    if blockIndex >= 0 && blockIndex < len(Blockchain) {
        Blockchain[blockIndex].Transaction = newTransaction
        Blockchain[blockIndex].Hash = CalculateHash(&Blockchain[blockIndex])
    } else {
		fmt.Printf("Please enter a valid index !")
	}
}

func VerifyChain() bool {
    for i := 1; i < len(Blockchain); i++ {
        currentBlock := Blockchain[i]
        previousBlock := Blockchain[i-1]

        if currentBlock.Hash != CalculateHash(&currentBlock) {
            return false
        }

        if currentBlock.PreviousHash != previousBlock.Hash {
            return false
        }
    }
    return true
}

func CalculateHash(block *Block) string {
    data := fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash)
    hashBytes := sha256.Sum256([]byte(data))
    return hex.EncodeToString(hashBytes[:])
}
