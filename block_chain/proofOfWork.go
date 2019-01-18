// proofOfWork
package main

type ProofOfWrok struct {
	block *Block

	//目标值
	target *big.Int
}

const targetBits = 24

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := ProofOfWork{block: block, target: target}
	return &pow
}

func (pow *ProofOfWork) PrepareData(nonce int64) []byte {
	block := pow.block
	temp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		block.MerKelRoot,
		IntToByte(block.TimeStamp),
		IntToByte(targetBits),
		IntToByte(nonce),
		block.Data}
	data := bytes.Join(temp, []byte{})
	return data
}

func (pow *ProofOfWrok) Run() (int64, []byte) {
	var nonce int64 = 0
	var hash [32]byte
	var hashInt big.Int
	for nonce < math.MaxInt64 {
		data := pow.PrepareData(nonce)
		hash = sha256.Sum256(data)

		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("found nonce, nonce:%d, hash:%x\n", nonce, hash)
			break
		} else {
			fmt.Printf("not found nonce, current nonce:%d, current hash:%x\n", nonce, hash)
			nonce++
		}
	}
	return nonce, hash[:]
}
