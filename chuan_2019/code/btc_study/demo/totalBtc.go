package main

import "fmt"

func main()  {
	fmt.Println("hello")
	total := 0.0
	blockInterval := 21.0
	currentReward := 50.0
	for currentReward > 0 {
		amount1 := blockInterval * currentReward
		currentReward *= 0.5
		total += amount1
	}
	fmt.Println("比特币总量：", total, "万")
}
