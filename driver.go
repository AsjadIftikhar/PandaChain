package main

import (
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
	return ""
}

func InsertBlock(transactionsToInsert []string, chainHead *Block) *Block {

	//insert new block and return head pointer
	newBlock := Block{
		transactions: transactionsToInsert,
		prevPointer:  chainHead,
		prevHash:     "",
		currentHash:  "",
	}

	return &newBlock
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {

	//change transaction data inside block
}

func ListBlocks(chainHead *Block) {

	//display the data(transaction) inside all blocks
	currPointer := chainHead
	for i := 0; currPointer != nil; i++ {
		fmt.Print(currPointer.transactions)
		fmt.Print("->")

		currPointer = currPointer.prevPointer
	}

}

func VerifyChain(chainHead *Block) {

	//check whether "Block chain is compromised" or "Block chain is unchanged"

}
func main() {

	transactions := []string{"Alice Sent $100 to Bob", "Bob Received $50"}
	transactions1 := []string{"A"}
	transactions2 := []string{"B"}
	transactions3 := []string{"C"}
	transactions4 := []string{"D"}

	chainHead := InsertBlock(transactions, nil)
	chainHead = InsertBlock(transactions1, chainHead)
	chainHead = InsertBlock(transactions2, chainHead)
	chainHead = InsertBlock(transactions3, chainHead)
	chainHead = InsertBlock(transactions4, chainHead)

	ListBlocks(chainHead)

}
