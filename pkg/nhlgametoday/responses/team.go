package responses

// TeamInfoResponse represents the JSON response from the NHL API.
type TeamInfoResponse struct {
	Name   string `json:"name"`
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`
	Ot     int    `json:"ot"`
}
