package board

import (
	"testing"
)

func BenchmarkNewBoard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newBoard()
	}
}
