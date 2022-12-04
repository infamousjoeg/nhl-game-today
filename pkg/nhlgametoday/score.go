package nhlgametoday

import "github.com/infamousjoeg/nhl-game-today/pkg/nhlgametoday/responses"

// GetScore returns the score for today's game.
func GetScore(gameInfo responses.GameInfoResponse) responses.ScoreResponse {
	// Get the score information.
	score := responses.ScoreResponse{
		HomeTeamScore: gameInfo.Dates[0].Games[0].Teams.Home.Score,
		AwayTeamScore: gameInfo.Dates[0].Games[0].Teams.Away.Score,
	}

	// Return the score information.
	return score
}
