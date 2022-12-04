package responses

// ScoreResponse represents the JSON response from the NHL API.
type ScoreResponse struct {
	HomeTeamScore int `json:"homeTeamScore"`
	AwayTeamScore int `json:"awayTeamScore"`
}
