package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	transactions []string
	prevPointer  *Block
	prevHash     string
	currentHash  string
}

func CalculateHash(inputBlock *Block) string {
	//Calculate Hash of a Block
	message := ""
	for i := 0; i < len(inputBlock.transactions); i++ {
		message += inputBlock.transactions[i]
	}
	message += inputBlock.prevHash
	converted := []byte(message)
	hashTool := sha256.New()
	hashTool.Write(converted)

	return hex.EncodeToString(hashTool.Sum(nil))
}

func InsertBlock(transactionsToInsert []string, chainHead *Block) *Block {

	//insert new block and return head pointer
	newBlock := Block{
		transactions: transactionsToInsert,
		prevPointer:  chainHead,
		prevHash:     "",
		currentHash:  "",
	}

	if chainHead != nil {
		newBlock.prevHash = chainHead.currentHash
	}

	newBlock.currentHash = CalculateHash(&newBlock)

	return &newBlock
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {

	//change transaction data inside block
	currPointer := chainHead
	for i := 0; currPointer != nil; i++ {
		for j := 0; j < len(currPointer.transactions); j++ {
			if currPointer.transactions[j] == oldTrans {
				currPointer.transactions[j] = newTrans
			}
		}

		currPointer = currPointer.prevPointer
	}

}

func ListBlocks(chainHead *Block) {

	//display the data(transaction) inside all blocks
	fmt.Println()
	currPointer := chainHead
	for i := 0; currPointer != nil; i++ {
		fmt.Print(currPointer.transactions, " | ", currPointer.currentHash, " | ", currPointer.prevHash)
		fmt.Println("<--")

		currPointer = currPointer.prevPointer
	}

}

func VerifyChain(chainHead *Block) {

	//check whether "Block chain is compromised" or "Block chain is unchanged"
	currPointer := chainHead.prevPointer
	prevHash := chainHead.prevHash

	for i := 0; currPointer != nil; i++ {
		currHash := CalculateHash(currPointer)
		if currHash != prevHash {
			fmt.Println("Chain Has Been Compromised")
			return
		}
		prevHash = currPointer.prevHash
		currPointer = currPointer.prevPointer
	}
	fmt.Println("Chain is unchanged")

}
func main() {

	chainHead := InsertBlock([]string{"Alice Sent $100 to Bob", "Bob Sent $50 to Chad"}, nil)
	chainHead = InsertBlock([]string{"Alice Minted an NFT"}, chainHead)
	chainHead = InsertBlock([]string{"Bob is Alive"}, chainHead)
	chainHead = InsertBlock([]string{"Why am I Doing this?", "Such a bad assignment tbh!", "Lets finish it ASAP"}, chainHead)
	chainHead = InsertBlock([]string{"Stupid Course"}, chainHead)

	//ChangeBlock("Bob is Alive", "Bob is dead", chainHead)
	ListBlocks(chainHead)
	VerifyChain(chainHead)

}
