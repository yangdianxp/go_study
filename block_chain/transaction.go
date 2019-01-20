package main

type Transaction struct {
	//交易ID
	TXID []byte
	//输入
	TXInputs []TXInput
	//输出
	TXOutputs []TXOutput
}

type TXInput struct {
	//所引用输出的交易ID
	TXID []byte
	//所引用output的索引值
	Vout int64
	//解锁脚本，指明可以使用某个output的条件
	ScriptSig string
}

type TXOutput struct {
	//支付给收款方的金额
	Value float64
	//锁定脚本， 指定收款方的地址
	ScriptPubKey string
}