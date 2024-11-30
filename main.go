package main

import (
	"fmt"
	"rubrumcreation.com/go-paper-scissor/services"
)

// Currently before further optimizations, 1 million games are played and summarized in 0,118 seconds
// After optimizations the random array of moves to a global static array, the new runtime is 0,112
func main() {
	fmt.Println("Starting ....")
	numOfGames := 1000000
	summary := services.PlayRandomGames(numOfGames)
	summary.PrintSummary()
}

