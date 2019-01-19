// common
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func IntToByte(num int64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	CheckErr("IntToByte", err)
	return buffer.Bytes()
}

func CheckErr(info string, err error) {
	if err != nil {

		fmt.Println(info, " err occur: ", err)
		os.Exit(1)
	}
}
