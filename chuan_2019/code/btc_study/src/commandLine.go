package main

import "fmt"

func (cli *CLI) AddBlock(data string)  {
	//cli.bc.AddBlock(data)
}

func (cli *CLI) PrintBlockChain()  {
	bc := cli.bc
	it := bc.NewIterator()
	for {
		block := it.Next()
		block.Print()
		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CLI) GetBalance(address string) {
	utxos := cli.bc.FindUTXOs(address)

	total := 0.0
	for _, utxo := range utxos {
		total += utxo.Value
	}
	fmt.Printf("[%s]的余额为: %f\n", address, total)
}