package block

import (
	"crypto/sha512"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Block[T comparable] struct {
	Previous  *Block[T]
	Nonce     int64
	Data      T
	Timestamp int64
}

func New[T comparable](Previous *Block[T], data T) *Block[T] {
	block := new(Block[T])

	block.Previous = Previous
	block.Nonce = rand.Int63n(int64(math.Pow(2, 32)))
	block.Data = data
	block.Timestamp = time.Now().Unix()

	return block
}

func (block *Block[T]) Hash() string {
	return fmt.Sprintf("%x", sha512.Sum512(([]byte(block.String()))))
}

func (block *Block[T]) String() string {
	if block.Previous == nil {
		return fmt.Sprintf("%d:%v:%d",
			block.Nonce, block.Data, block.Timestamp)
	}

	return fmt.Sprintf("%s:%d:%v:%d",
		block.Previous.Hash(), block.Nonce, block.Data, block.Timestamp)
}

func (block *Block[T]) Show() {
	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)

	t.AppendHeader(table.Row{"BLOCK", "VALUE"})

	if block.Previous != nil {
		t.AppendRow(table.Row{"Previous",
			fmt.Sprintf("%s...", block.Previous.Hash()[:32])})
	}

	t.AppendRows([]table.Row{
		{"Nonce", fmt.Sprintf("%d", block.Nonce)},
		{"Hash", fmt.Sprintf("%s...", block.Hash()[:32])},
		{"Nonce", block.Nonce},
		{"Data", block.Data},
		{"Timestamp", fmt.Sprintf("%d", block.Timestamp)}})

	t.Render()
}
