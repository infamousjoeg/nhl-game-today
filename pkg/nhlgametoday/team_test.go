package nhlgametoday_test

import (
	"testing"

	"github.com/infamousjoeg/nhl-game-today/pkg/nhlgametoday"
)

// TestGetTeamInfo tests the GetTeamInfo function.
func TestGetTeamInfo(t *testing.T) {

	// Get the game information.
	gameInfo, err := nhlgametoday.GetGameInfo(14, "2022-11-11")
	if err != nil {
		t.Errorf("GetGameInfo() returned an error: %v", err)
	}

	// Get the home team information.
	homeTeamInfo := nhlgametoday.GetTeamInfo("home", gameInfo)

	// Get the away team information.
	awayTeamInfo := nhlgametoday.GetTeamInfo("away", gameInfo)

	// Check if the home team name is correct.
	if homeTeamInfo.Name != "Washington Capitals" {
		t.Errorf("GetTeamInfo() returned an incorrect home team name: %v", homeTeamInfo.Name)
	}

	// Check if the away team name is correct.
	if awayTeamInfo.Name != "Tampa Bay Lightning" {
		t.Errorf("GetTeamInfo() returned an incorrect away team name: %v", awayTeamInfo.Name)
	}

	if homeTeamInfo.Wins < 0 || homeTeamInfo.Losses < 0 || homeTeamInfo.Ot < 0 {
		t.Errorf("GetTeamInfo() returned an invalid response")
	}

	if awayTeamInfo.Wins < 0 || awayTeamInfo.Losses < 0 || awayTeamInfo.Ot < 0 {
		t.Errorf("GetTeamInfo() returned an invalid response")
	}
}
