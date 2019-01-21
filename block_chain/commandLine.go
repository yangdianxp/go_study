package main

import (
	"flag"
	"fmt"
	"os"
)

const usage = `
	createChain --address ADDRESS	"create a blockchain"
	addBlock --data DATA            "add a block to blockchain"
	send --from FROM --to TO --amount AMOUNT	"send coin from FROM to TO"
	getBalance --address ADRESS		"get balance of the addressd"
	printChain						"print all blocks"
`

const AddBlockCmdString = "addBlock"
const PrintChainCmdString = "printChain"
const CreateChainCmdString = "createChain"
const GetBalanceCmdString = "getBalance"

type CLI struct {
	//bc *BlockChain
}

func (cli *CLI)printUsage()  {
	fmt.Println(usage)
	os.Exit(1)
}

func (cli *CLI)parameterCheck()  {
	if len(os.Args) < 2 {
		fmt.Println("invalid input!")
		cli.printUsage()
	}
}

func (cli *CLI)Run()  {
	cli.parameterCheck()
	addBlockCmd := flag.NewFlagSet(AddBlockCmdString, flag.ExitOnError)
	createChainCmd := flag.NewFlagSet(CreateChainCmdString, flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet(GetBalanceCmdString, flag.ExitOnError)
	printChainCmd := flag.NewFlagSet(PrintChainCmdString, flag.ExitOnError)

	addBlockCmdPara := addBlockCmd.String("data", "", "block transaction info!")
	createChainCmdPara := createChainCmd.String("address", "", "address info!")
	getBalanceCmdPara := getBalanceCmd.String("address", "", "address info!")

	switch os.Args[1] {
	case CreateChainCmdString:
		err := createChainCmd.Parse(os.Args[2:])
		CheckErr("Run1()", err)
		if createChainCmd.Parsed() {
			if *createChainCmdPara == "" {
				cli.printUsage()
			}
			cli.CreateChain(*createChainCmdPara)
		}
	case AddBlockCmdString:
		err := addBlockCmd.Parse(os.Args[2:])
		CheckErr("Run2()", err)
		if addBlockCmd.Parsed() {
			if *addBlockCmdPara == "" {
				cli.printUsage()
			}
			cli.AddBlock(*addBlockCmdPara)
		}
	case GetBalanceCmdString:
		err := getBalanceCmd.Parse(os.Args[2:])
		CheckErr("Run4()", err)
		if getBalanceCmd.Parsed() {
			if *getBalanceCmdPara == "" {
				cli.printUsage()
			}
			cli.GetBalance(*getBalanceCmdPara)
		}
	case PrintChainCmdString:
		err := printChainCmd.Parse(os.Args[2:])
		CheckErr("Run3()", err)
		if printChainCmd.Parsed() {
			cli.PrintChain()
		}
	default:
		cli.printUsage()
	}
}