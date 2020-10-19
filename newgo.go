package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type Block struct {
	transactions []string
	prevPointer  *Block
	prevHash     string
	currentHash  string
}

func CalculateHash(inputBlock *Block) string {
	str := strings.Join(inputBlock.transactions[:], ",")
	str += inputBlock.prevHash
	h := sha256.Sum256([]byte(str))
	return hex.EncodeToString(h[:])
}

func InsertBlock(transactionsToInsert []string, chainHead *Block) *Block {
	if chainHead == nil {
		chainHead = new(Block)
		for i := 0; i < len(transactionsToInsert); i++ {
			chainHead.transactions = append(chainHead.transactions, transactionsToInsert[i])
		}
		chainHead.prevHash = ""
		chainHead.prevPointer = nil
		chainHead.currentHash = CalculateHash(chainHead)
		return chainHead
	}
	a := chainHead
	for a.prevPointer != nil {
		a = a.prevPointer
	}
	var b *Block
	b = new(Block)
	for i := 0; i < len(transactionsToInsert); i++ {
		b.transactions = append(b.transactions, transactionsToInsert[i])
	}
	b.currentHash = CalculateHash(b)
	a.prevPointer = b
	b.prevPointer = nil
	b.prevHash = a.currentHash
	return chainHead
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	if chainHead != nil {
		for chainHead.prevPointer != nil {
			for i := 0; i < len(chainHead.transactions); i++ {
				if chainHead.transactions[i] == oldTrans {
					chainHead.transactions[i] = newTrans
				}
			}
			chainHead = chainHead.prevPointer
		}
	}
}

func ListBlocks(chainHead *Block) {
	if chainHead != nil {
		fmt.Println(chainHead.currentHash)
		for chainHead.prevPointer != nil {
			chainHead = chainHead.prevPointer
			fmt.Println(chainHead.currentHash)
		}
	}
}

func VerifyChain(chainHead *Block) {
	if chainHead != nil {
		i := 0
		for chainHead.prevPointer != nil {
			s := CalculateHash(chainHead)
			if s != chainHead.currentHash {
				fmt.Printf("Chain compromised at %d block\n", i)
			}
			i++
			chainHead = chainHead.prevPointer
		}
	}
}

func main() {
	var chainHead *Block
	genesis := []string{"S2E", "S2Z"}
	chainHead = InsertBlock(genesis, chainHead)

	firstBlock := []string{"E2Alice", "E2Bob", "S2John"}
	chainHead = InsertBlock(firstBlock, chainHead)

	ListBlocks(chainHead)
	ChangeBlock("S2E", "S2Trudy", chainHead)

	ListBlocks(chainHead)
	VerifyChain(chainHead)

	fmt.Println("OkY")

}
