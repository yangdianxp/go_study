package main

func (cli *CLI) AddBlock(data string)  {
	cli.bc.AddBlock(data)
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