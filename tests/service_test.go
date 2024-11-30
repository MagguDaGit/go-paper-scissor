package tests

import (
	"fmt"
	"rubrumcreation.com/go-paper-scissor/models"
	"rubrumcreation.com/go-paper-scissor/services"
	"testing"
)

// Objecte oriented static string performance: 1 000 000 ->  0.91s
// Object oriented static int performance : 1 00 000 -> 0.94s
func TestRandomGame(t *testing.T) {

	var player1Wins int = 0
	var player2Wins int = 0
	var draws int = 0
	for i := 0; i <= 1000000; i++ {
		var result models.Result = services.PlayRandomGame()
		if result.Winner == models.PLAYER_1 && result.Draw == false {
			player1Wins++

		} else if result.Winner == models.PLAYER_2 && result.Draw == false {
			player2Wins++
		} else if result.Draw {
			draws++
		} else {
			t.Error("No winner/ expected winner, and DRAW is false, this is a bug...")
		}
	}

	fmt.Println("--------------------------------")
	fmt.Println("Result of random games testing: ")
	fmt.Printf("Player 1 wins: %v", player1Wins)
	fmt.Println()
	fmt.Printf("Player 2 wins: %v ", player2Wins)
	fmt.Println()
	fmt.Printf("Draws: %v  ", draws)
	fmt.Println()
	fmt.Println("--------------------------------")
}
