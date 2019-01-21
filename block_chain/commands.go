package main

import "fmt"

func (cli *CLI) AddBlock(data string) {
	//bc := GetBlockChainHandler()
	//bc.AddBlock(data) //TODO
}

func (cli *CLI) PrintChain() {
	bc := GetBlockChainHandler()
	defer bc.db.Close()
	it := bc.NewIterator()
	for {
		block := it.Next()
		fmt.Printf("Version: %d\n", block.Version)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("TimeStamp: %d\n", block.TimeStamp)
		fmt.Printf("Bits: %x\n", block.Bits)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		//fmt.Printf("Data: %s\n", block.Data) //TODO
		fmt.Printf("Isvalid: %v\n", NewProofOfWork(block).IsValid())

		if len(block.PrevBlockHash) == 0 {
			fmt.Println("print over!")
			break
		}
	}
}

func (cli *CLI)CreateChain(address string)  {
	bc := InitBlockChain(address)
	defer bc.db.Close()
	fmt.Println("Create blockchain successfully")
}

func (cli *CLI)GetBalance(address string) {
	bc := GetBlockChainHandler()
	defer bc.db.Close()

	utxos := bc.FindUTXO(address)
	//总金额
	var total float64 = 0
	//遍历所有的utxo获取金额总数
	for _, utxo := range utxos {
		total += utxo.Value
	}
	fmt.Printf("The balance of %s is %f\n", address, total)
}
