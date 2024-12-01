package models

type GameSummaryResponse struct {
	Title   string
	Message string
	Data    GameSummary
	Success bool
}
