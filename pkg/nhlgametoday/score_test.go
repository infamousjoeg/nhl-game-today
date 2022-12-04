package nhlgametoday_test

import (
	"testing"

	"github.com/infamousjoeg/nhl-game-today/pkg/nhlgametoday"
)

// TestGetScore tests the GetScore function.
func TestGetScore(t *testing.T) {
	// Get the game information for the team.
	gameInfo, err := nhlgametoday.GetGameInfo(14, "2022-11-11")
	if err != nil {
		t.Errorf("GetGameInfo() returned an error: %v", err)
	}

	// Get the score information.
	score := nhlgametoday.GetScore(gameInfo)

	// Check if the score information is empty.
	if score.HomeTeamScore < 0 || score.AwayTeamScore < 0 {
		t.Errorf("GetScore() returned an invalid response")
	}
}
