// main
package main

import (
	"fmt"
)

func main() {
	bc := NewBlockChain()
	//	bc.AddBlock(data:"A send B 1BTC")
	//	bc.AddBlock(data:"B send C 1BTC")

	for _, block := range bc.blocks {
		//		fmt.Printf(fomat:"Version: %d\n", block.Version)
		//		fmt.Printf(fomat:"PrevBlockHash: %x\n", block.PreBlockHash)
		//		fmt.Printf(fomat:"Hash: %x\n", block.Hash)
		//		fmt.Printf(fomat:"TimeStamp: %d\n", block.TimeStamp)
		//		fmt.Printf(fomat:"Bits: %x\n", block.Bits)
		//		fmt.Printf(fomat:"Nonce: %d\n", block.Nonce)
		//		fmt.Printf(fomat:"Data: %s\n", block.Data)
	}
}
