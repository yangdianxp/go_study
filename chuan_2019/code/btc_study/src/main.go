package main

func main()  {
	bc := NewBlockChain("张三")
	cli := CLI{bc:bc}
	cli.Run()
}
