package tests

import (
	"rubrumcreation.com/go-paper-scissor/services"
	"testing"
)

func TestRandomGame(t *testing.T) {
	services.PlayRandomGames(100)
}
