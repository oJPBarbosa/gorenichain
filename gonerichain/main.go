package main

import "gonerichain/chain"

func main() {
	chain := chain.New[int]()

	chain.AddBlock(1)
	chain.AddBlock(2)
	chain.AddBlock(3)

	chain.List()

	chain.FindBlockByData(0)
	chain.FindBlockByData(1)
}
