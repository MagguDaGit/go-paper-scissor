package models

import (
	"fmt"
	"math/rand"
)

// MOVES
type MoveType int8

const (
	ROCK    MoveType = 1
	PAPER   MoveType = 2
	SCISSOR MoveType = 3
)

// Move interface to define behavior for each move
type Move interface {
	GetType() MoveType // Capitalized to make it public
}

// Implementations of Move interface for each type of move
type Rock struct{}
type Paper struct{}
type Scissor struct{}

func (r Rock) GetType() MoveType {
	return ROCK
}

func (p Paper) GetType() MoveType {
	return PAPER
}

func (s Scissor) GetType() MoveType {
	return SCISSOR
}

//  Player

type PlayerName string

const (
	PLAYER_1 PlayerName = "PLAYER 1"
	PLAYER_2 PlayerName = "PLAYER 2"
)

type Player interface {
	PlayMove() Move
	GetName() PlayerName
}

// Generic player, has a chosen move
type GenericPlayer struct {
	Name       PlayerName
	ChosenMove Move
}

func (p GenericPlayer) GetName() PlayerName {
	return p.Name
}
func (p GenericPlayer) PlayMove() Move {
	return p.ChosenMove
}

// Only plays random moves
type RandomPlayer struct {
	Name PlayerName
}

func (p RandomPlayer) GetName() PlayerName {
	return p.Name
}

var availableMoves = [3]Move{Rock{}, Paper{}, Scissor{}}

func (p RandomPlayer) PlayMove() Move {
	return availableMoves[rand.Intn(3)]
}

// Result
type Result struct {
	Winner PlayerName
	Loser  PlayerName
	Draw   bool
}

func (r Result) getWinner() string {
	return string(r.Winner)
}

func (r Result) getLoser() string {
	return string(r.Loser)
}

func (r Result) isDraw() bool {
	return r.Draw
}

// A single object to summarize a series of games, uses the list of results to construct the summary
type GameSummary struct {
	Results     []Result
	Duration    float64
	Winner      string
	Player1Wins int
	Player2Wins int
	Draws       int
}

func ConstructSummary(results []Result, duration float64) GameSummary {
	summary := GameSummary{}
	summary.Duration = duration
	for i := 0; i < len(results); i++ {
		if results[i].isDraw() {
			summary.Draws = summary.Draws + 1
		} else if results[i].getWinner() == string(PLAYER_1) {
			summary.Player1Wins = summary.Player1Wins + 1
		} else if results[i].getWinner() == string(PLAYER_2) {
			summary.Player2Wins = summary.Player2Wins + 1
		} else {
			fmt.Println("This should never be called... ")
		}
	}
	if summary.Player1Wins > summary.Player2Wins {
		summary.Winner = string(PLAYER_1)
	} else if summary.Player2Wins > summary.Player1Wins {
		summary.Winner = string(PLAYER_2)
	} else {
		summary.Winner = "DRAW"
	}
	summary.Results = results
	return summary

}

func (summary GameSummary) PrintSummary() {
	fmt.Println("--- GAME SUMMARY ---")
	fmt.Printf("Number of games played: %v \n", len(summary.Results))
	fmt.Printf("Duration of games: %v seconds \n", summary.Duration)
	fmt.Printf("Player 1 wins: %v \n", summary.Player1Wins)
	fmt.Printf("Player 2 wins: %v \n", summary.Player2Wins)

	fmt.Printf("Draws: %v \n", summary.Draws)
	fmt.Printf("Winner: %s \n", summary.Winner)
	fmt.Println("--- END OF SUMMARY ---")
}
