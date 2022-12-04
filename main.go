package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/infamousjoeg/nhl-game-today/pkg/nhlgametoday"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the date query parameter.
		date := r.URL.Query().Get("date")
		if date == "" {
			// Set the date to today.
			date = time.Now().Format("2006-01-02")
		}
		// Get the game information for the team. 14 = Tampa Bay Lightning.
		todaysGameInfo, err := nhlgametoday.GetGameInfo(14, date)
		if err != nil {
			log.Fatalf("Error getting game info: %v", err)
		}
		// Check if there is a game today.
		if todaysGameInfo.TotalGames == 0 {
			fmt.Fprintf(w, "There is no game scheduled for today.")
			return
		}
		// Home team
		homeTeamInfo := nhlgametoday.GetTeamInfo("home", todaysGameInfo)
		// Away team
		awayTeamInfo := nhlgametoday.GetTeamInfo("away", todaysGameInfo)
		// Venue
		venue := todaysGameInfo.Dates[0].Games[0].Venue.Name
		if todaysGameInfo.Dates[0].Games[0].Status.DetailedState != "Final" {
			// If the game has not been played yet, just display the game info.
			fmt.Fprintf(w, "Today's game is between the %s (%v-%v-%v) and the %s (%v-%v-%v) at %s.", awayTeamInfo.Name, awayTeamInfo.Wins, awayTeamInfo.Losses, awayTeamInfo.Ot, homeTeamInfo.Name, homeTeamInfo.Wins, homeTeamInfo.Losses, homeTeamInfo.Ot, venue)
		} else {
			// If the game has been played, display the score along with game info.
			score := nhlgametoday.GetScore(todaysGameInfo)
			fmt.Fprintf(w, "Today's game was between the %s (%v-%v-%v) and the %s (%v-%v-%v) at %s.\n\nThe final score was %v - %v.", awayTeamInfo.Name, awayTeamInfo.Wins, awayTeamInfo.Losses, awayTeamInfo.Ot, homeTeamInfo.Name, homeTeamInfo.Wins, homeTeamInfo.Losses, homeTeamInfo.Ot, venue, score.AwayTeamScore, score.HomeTeamScore)
		}
	})
	http.ListenAndServe(":8080", nil)
}
