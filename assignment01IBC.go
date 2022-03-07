package main

//import (
//	"fmt"
//)
//
//type Block struct {
//	transactions []string
//	prevPointer  *Block
//	prevHash     string
//	currentHash  string
//}
//
//func CalculateHash(inputBlock *Block) string {
//	//Calculate Hash of a Block
//	return ""
//}
//
//func InsertBlock(transactionsToInsert []string, chainHead *Block) *Block {
//
//	//insert new block and return head pointer
//	newBlock := Block{
//		transactions: transactionsToInsert,
//		prevPointer:  chainHead,
//		prevHash:     "",
//		currentHash:  "",
//	}
//	fmt.Println(newBlock)
//
//	return &newBlock
//}
//
//func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
//
//	//change transaction data inside block
//}
//
//func ListBlocks(chainHead *Block) {
//
//	//display the data(transaction) inside all blocks
//	currPointer := chainHead
//	for i := 0; currPointer != nil; i++ {
//		fmt.Print(currPointer.transactions)
//		fmt.Println("->")
//
//		currPointer = currPointer.prevPointer
//	}
//
//}
//
//func VerifyChain(chainHead *Block) {
//
//	//check whether "Block chain is compromised" or "Block chain is unchanged"
//
//}
