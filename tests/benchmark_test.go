package tests

import (
	"rubrumcreation.com/go-paper-scissor/services"
	"testing"
)

/*func BenchmarkFewPlayRandomGames(b *testing.B) {
	services.PlayRandomGames(100) // Adjust the number of games as needed
}

func BenchmarkPlayManyRandomGamesParallel(b *testing.B) {
	numberOfGames := 10000000
	services.PlayRandomGames(numberOfGames)
}
*/
func BenchmarkBillionRandomGamesParallel(b *testing.B) {
	numberOfGames := 1000000000
	services.PlayRandomGames(numberOfGames)
}
