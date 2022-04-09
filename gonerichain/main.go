package main

import (
	"gonerichain/chain"
)

func main() {
	chain := chain.New[int]()

	chain.AddBlock(1)
	chain.AddBlock(2)
	chain.AddBlock(3)

	chain.List()

	chain.PrintBlockByData(0)
	chain.PrintBlockByData(1)
}
