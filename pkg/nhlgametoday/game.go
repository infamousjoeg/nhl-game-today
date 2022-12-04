package nhlgametoday

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/infamousjoeg/nhl-game-today/pkg/nhlgametoday/responses"
)

// GetGameInfo returns today's game information for an NHL team.
func GetGameInfo(teamID int, date string) (responses.GameInfoResponse, error) {
	// Get the game information for the team.
	resp, err := http.Get("https://statsapi.web.nhl.com/api/v1/schedule?date=" + date + "&teamId=" + strconv.Itoa(teamID))
	if err != nil {
		return responses.GameInfoResponse{}, err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return responses.GameInfoResponse{}, err
	}

	// Parse the JSON response.
	var gameInfo responses.GameInfoResponse
	err = json.Unmarshal(body, &gameInfo)
	if err != nil || gameInfo.TotalGames == 0 {
		return responses.GameInfoResponse{}, err
	}

	// Return the game information.
	return gameInfo, nil
}
