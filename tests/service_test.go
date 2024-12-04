package tests

import (
	"rubrumcreation.com/go-paper-scissor/services"
	"testing"
)

func TestHundredRandomGame(t *testing.T) {
	services.PlayRandomGames(100)
}

func TestTenMillionRandomGames(t *testing.T) {
	services.PlayRandomGames(10000000)
}
