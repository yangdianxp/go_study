package main

import "fmt"

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

func (cli *CLI) Send(from, to string, amount float64, miner, data string) {
	fmt.Printf("from : %s\n", from)
	fmt.Printf("to : %s\n", to)
	fmt.Printf("amount : %f\n", amount)
	fmt.Printf("miner : %s\n", miner)
	fmt.Printf("data : %s\n", data)

	// 1. 创建挖矿交易
	coinbase := NewCoinbaseTX(miner, data)
	// 2. 创建一个普通交易
	tx := NewTransaction(from, to, amount, cli.bc)
	if tx == nil {
		fmt.Printf("无效的交易")
		return
	}
	// 3. 添加到区块
	cli.bc.AddBlock([]*Transaction{coinbase, tx})
	fmt.Printf("转账成功！")
}

func (cli *CLI)NewWallet()  {
	//ws := CreateWallet()
	//ws.Print()
	ws := NewWallets()
	ws.CreateWallet()
	ws.Print()
}

func (cli *CLI)List()  {
	ws := NewWallets()
	ws.Print()
}