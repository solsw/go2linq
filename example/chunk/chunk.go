//go:build go1.18

package main

import (
	"fmt"

	"github.com/solsw/go2linq/v2"
)

// https://docs.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/partitioning-data#example

func main() {
	chunkNumber := 0
	enrChunks := go2linq.ChunkMust(go2linq.RangeMust(0, 8), 3).GetEnumerator()
	for enrChunks.MoveNext() {
		chunkNumber++
		fmt.Printf("Chunk %d\n", chunkNumber)
		chunk := enrChunks.Current()
		enrItems := go2linq.NewEnSlice(chunk...).GetEnumerator()
		for enrItems.MoveNext() {
			item := enrItems.Current()
			fmt.Printf("    %d\n", item)
		}
		fmt.Println()
	}
}
