package tests

import (
	"rubrumcreation.com/go-paper-scissor/services"
	"testing"
)

func BenchmarkPlayRandomGames(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.PlayRandomGames(1000) // Adjust the number of games as needed
	}
}
