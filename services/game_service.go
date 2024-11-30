package services

import (
	"time"

	"rubrumcreation.com/go-paper-scissor/models"
)

func PlayRandomGames(numberOfGames int) models.GameSummary {
	start := time.Now()
	player1 := models.RandomPlayer{}
	player1.Name = models.PLAYER_1

	player2 := models.RandomPlayer{}
	player2.Name = models.PLAYER_2
	results := make([]models.Result, numberOfGames, numberOfGames)
	for i := 0; i < numberOfGames; i++ {
		result := play(player1, player2)
		results[i] = result
	}
	var elapsed float64 = time.Since(start).Abs().Seconds()
	summary := models.ConstructSummary(results, elapsed)
	return summary
}

// Direct comparrison with predefined rules. Could use map lookup, modulo or bitmasking.
// But the rules of rock paper scissor is so simple this is a fast way to determine a winner given a predefined ruleset
func determineWinner(m1 models.Move, m2 models.Move) int {
	switch {
	case m1 == m2:
		return 0 // draw
	case (m1.GetType() == models.ROCK && m2.GetType() == models.SCISSOR) ||
		(m1.GetType() == models.PAPER && m2.GetType() == models.ROCK) ||
		(m1.GetType() == models.SCISSOR && m2.GetType() == models.PAPER):
		return 1 // move 1 wins
	default:
		return 2 // move 2 wins
	}
}

func play(p1 models.Player, p2 models.Player) models.Result {
	player1Move := p1.PlayMove()
	player2Move := p2.PlayMove()
	result := models.Result{}
	winner := determineWinner(player1Move, player2Move)

	switch {
	case winner == 0:
		{
			result.Draw = true
			return result
		}
	case winner == 1:
		{
			result.Winner = p1.GetName()
			result.Loser = p2.GetName()
			result.Draw = false
			return result
		}
	case winner == 2:
		{
			result.Winner = p2.GetName()
			result.Loser = p1.GetName()
			result.Draw = false
			return result
		}
	default:
		return result
	}

}
