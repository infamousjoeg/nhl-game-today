package nhlgametoday

import "github.com/infamousjoeg/nhl-game-today/pkg/nhlgametoday/responses"

// GetTeamInfo returns either the home or away team's information for today's game.
func GetTeamInfo(teamType string, gameInfo responses.GameInfoResponse) responses.TeamInfoResponse {
	if teamType == "home" {
		// Get the home team information.
		teamInfo := responses.TeamInfoResponse{
			Name:   gameInfo.Dates[0].Games[0].Teams.Home.Team.Name,
			Wins:   gameInfo.Dates[0].Games[0].Teams.Home.LeagueRecord.Wins,
			Losses: gameInfo.Dates[0].Games[0].Teams.Home.LeagueRecord.Losses,
			Ot:     gameInfo.Dates[0].Games[0].Teams.Home.LeagueRecord.Ot,
		}

		// Return the home team information.
		return teamInfo
	}

	// Get the away team information.
	teamInfo := responses.TeamInfoResponse{
		Name:   gameInfo.Dates[0].Games[0].Teams.Away.Team.Name,
		Wins:   gameInfo.Dates[0].Games[0].Teams.Away.LeagueRecord.Wins,
		Losses: gameInfo.Dates[0].Games[0].Teams.Away.LeagueRecord.Losses,
		Ot:     gameInfo.Dates[0].Games[0].Teams.Away.LeagueRecord.Ot,
	}

	// Return the away team information.
	return teamInfo

}
