package main

func main()  {
	bc := NewBlockChain("13KLq8gS8z4n8tUuQ1YG63149HgwebR1L7")
	cli := CLI{bc:bc}
	cli.Run()
}
