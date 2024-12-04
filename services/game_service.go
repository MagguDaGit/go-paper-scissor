package services

import (
	"sync"
	"time"

	"rubrumcreation.com/go-paper-scissor/models"
	"rubrumcreation.com/go-paper-scissor/routines"
)

func PlayRandomGames(numberOfGames int) models.GameSummary {
	start := time.Now()
	player1 := models.RandomPlayer{}
	player1.Name = models.PLAYER_1

	player2 := models.RandomPlayer{}
	player2.Name = models.PLAYER_2
	var elapsed float64 = time.Since(start).Abs().Seconds()
	var results []models.Result
	if numberOfGames < 100000 {
		results = SimulateRandomGamesSync(numberOfGames, &player1, &player2)
	} else {
		results = SimulateRandomGamesASync(numberOfGames, &player1, &player2)
	}
	summary := models.ConstructSummary(results, elapsed)
	return summary
}

func SimulateRandomGamesSync(numberOfGames int, player1 *models.RandomPlayer, player2 *models.RandomPlayer) []models.Result {
	results := make([]models.Result, numberOfGames, numberOfGames)
	for i := 0; i < numberOfGames; i++ {
		result := Play(player1, player2)
		results[i] = result
	}
	return results
}

func SimulateRandomGamesASync(numberOfGames int, player1 *models.RandomPlayer, player2 *models.RandomPlayer) []models.Result {
	resultsBuffer := make(chan []models.Result, numberOfGames) // Change channel to accept slices
	var waitGroup sync.WaitGroup

	// Get intervals for distributing games to workers
	intervals, err := routines.GetIntervals(numberOfGames)
	if err != nil {
		panic(err) // Or handle error appropriately
	}

	// Start workers for each interval
	for _, interval := range intervals {
		waitGroup.Add(1)
		go func(interval models.Interval) {
			defer waitGroup.Done()
			var batch []models.Result
			for i := interval.Start; i < interval.End; i++ {
				result := Play(player1, player2)
				batch = append(batch, result)
				if len(batch) >= 100 { // batch size for bulk sending
					resultsBuffer <- batch
					batch = nil
				}
			}
			if len(batch) > 0 {
				resultsBuffer <- batch
			}
		}(interval)
	}

	// Await all workers and close the results channel
	go func() {
		waitGroup.Wait()
		close(resultsBuffer)
	}()

	// Collect results
	var results []models.Result
	for batch := range resultsBuffer {
		results = append(results, batch...)
	}

	return results
}

func storeResultInChannel(p1 models.Player, p2 models.Player, results chan<- models.Result, wg *sync.WaitGroup) {
	defer wg.Done()
	results <- Play(p1, p2)
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

func Play(p1 models.Player, p2 models.Player) models.Result {
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
