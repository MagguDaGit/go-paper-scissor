package routines

import (
	"errors"

	"rubrumcreation.com/go-paper-scissor/models"
)

const MAX_WORKERS int = 4

func GetIntervals(numberOfGames int) ([]models.Interval, error) {

	if numberOfGames < 1 {
		return nil, errors.New("Cannot create intervals for less than 1 game")
	}
	details := make([]models.Interval, MAX_WORKERS, MAX_WORKERS)
	// Check for even distribution among workers
	remainder := numberOfGames % MAX_WORKERS
	totalProcessed := 0
	chunkSize := numberOfGames / MAX_WORKERS
	start := totalProcessed
	// if there is a remainder, the first worker will include this in their interval
	end := start + chunkSize + remainder
	chunkPartition := 0
	for totalProcessed <= numberOfGames {
		totalProcessed++
		if totalProcessed == end && chunkPartition < MAX_WORKERS {
			details[chunkPartition] = models.Interval{Start: start, End: end}
			chunkPartition++
			start = totalProcessed
			end = start + chunkSize
		}
	}
	return details, nil
}

func PlayGame(listOfResults []models.Result, start int, end int) {

}
