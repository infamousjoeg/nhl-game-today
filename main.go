package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/infamousjoeg/nhl-game-today/pkg/nhlgametoday"
)

func getDate(r *http.Request) string {
	date := r.URL.Query().Get("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}
	return date
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the date query parameter.
		date := getDate(r)

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
		// Collect home team information.
		homeTeamInfo := nhlgametoday.GetTeamInfo("home", todaysGameInfo)
		// Collect away team information.
		awayTeamInfo := nhlgametoday.GetTeamInfo("away", todaysGameInfo)
		// Collect venue information.
		venue := todaysGameInfo.Dates[0].Games[0].Venue.Name

		// Print the game information.
		switch todaysGameInfo.Dates[0].Games[0].Status.DetailedState {
		case "Scheduled": // If the game has not been played yet, just display the game info.
			fmt.Fprintf(w, "Today's game is between the %s (%v-%v-%v) and the %s (%v-%v-%v) at %s.", awayTeamInfo.Name, awayTeamInfo.Wins, awayTeamInfo.Losses, awayTeamInfo.Ot, homeTeamInfo.Name, homeTeamInfo.Wins, homeTeamInfo.Losses, homeTeamInfo.Ot, venue)
		case "Final": // If the game has been played, display the score along with game info.
			score := nhlgametoday.GetScore(todaysGameInfo)
			fmt.Fprintf(w, "Today's game was between the %s (%v-%v-%v) and the %s (%v-%v-%v) at %s.\n\nThe final score was %v - %v.", awayTeamInfo.Name, awayTeamInfo.Wins, awayTeamInfo.Losses, awayTeamInfo.Ot, homeTeamInfo.Name, homeTeamInfo.Wins, homeTeamInfo.Losses, homeTeamInfo.Ot, venue, score.AwayTeamScore, score.HomeTeamScore)
		case "In Progress": // If the game is in progress, display the score along with game info.
			score := nhlgametoday.GetScore(todaysGameInfo)
			fmt.Fprintf(w, "Today's game is between the %s (%v-%v-%v) and the %s (%v-%v-%v) at %s.\n\nThe score is %v - %v.", awayTeamInfo.Name, awayTeamInfo.Wins, awayTeamInfo.Losses, awayTeamInfo.Ot, homeTeamInfo.Name, homeTeamInfo.Wins, homeTeamInfo.Losses, homeTeamInfo.Ot, venue, score.AwayTeamScore, score.HomeTeamScore)
		}
	})

	// Get the SSL cert and key from environment variables.
	// To generate self-signed certs, run the following commands:
	// openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365
	// To use Let's Encrypt, run the following commands:
	// sudo apt-get install certbot
	// sudo certbot certonly --standalone -d <domain>
	// The certificates will be located at /etc/letsencrypt/live/<domain>/
	if os.Getenv("SSL_CERT_FILE") == "" {
		log.Fatal("SSL_CERT_FILE environment variable not set.")
	}
	if os.Getenv("SSL_KEY_FILE") == "" {
		log.Fatal("SSL_KEY_FILE environment variable not set.")
	}
	// Get the port from the environment variable.
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "8443")
	}

	// Start the server.
	fmt.Println("Loading SSL cert and key from environment variables...")
	fmt.Printf("Server started on port %s.\n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServeTLS(":"+os.Getenv("PORT"), os.Getenv("SSL_CERT_FILE"), os.Getenv("SSL_KEY_FILE"), nil))
}
