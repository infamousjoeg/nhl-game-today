package nhlgametoday_test

import (
	"testing"
	"time"

	"github.com/infamousjoeg/nhl-game-today/pkg/nhlgametoday"
)

// TestGetGameInfo tests the GetGameInfo function.
func TestGetGameInfo(t *testing.T) {
	// Set the date to today.
	date := time.Now().Format("2006-01-02")
	// Get the game information for the team.
	gameInfo, err := nhlgametoday.GetGameInfo(1, date)
	if err != nil {
		t.Errorf("GetGameInfo() returned an error: %v", err)
	}

	// Check if the game information is empty.
	if gameInfo.TotalGames < 0 || gameInfo.TotalGames > 1 {
		t.Errorf("GetGameInfo() returned an invalid response")
	}
}
