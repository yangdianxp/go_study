package main

import (
	"fmt"
	"os"
)

func main()  {
	len1 := len(os.Args)
	fmt.Printf("len1:%d", len1)
	for i, cmd := range os.Args {
		fmt.Printf("arg[%d] : %s\n", i, cmd)
	}
}
