package responses

// GameInfoResponse represents the JSON response from the NHL API.
type GameInfoResponse struct {
	TotalGames int     `json:"totalGames"`
	Dates      []Dates `json:"dates"`
}

// Dates represents the dates in the JSON response from the NHL API.
type Dates struct {
	Date  string  `json:"date"`
	Games []Games `json:"games"`
}

// Games represents the games in the JSON response from the NHL API.
type Games struct {
	Teams  Teams  `json:"teams"`
	Status Status `json:"status"`
	Venue  Venue  `json:"venue"`
}

// Teams represents the teams in the JSON response from the NHL API.
type Teams struct {
	Home Home `json:"home"`
	Away Away `json:"away"`
}

// Home represents the home team in the JSON response from the NHL API.
type Home struct {
	Team         Team         `json:"team"`
	LeagueRecord LeagueRecord `json:"leagueRecord"`
	Score        int          `json:"score"`
}

// Away represents the away team in the JSON response from the NHL API.
type Away struct {
	Team         Team         `json:"team"`
	LeagueRecord LeagueRecord `json:"leagueRecord"`
	Score        int          `json:"score"`
}

// Team represents the team in the JSON response from the NHL API.
type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// LeagueRecord represents the league record in the JSON response from the NHL API.
type LeagueRecord struct {
	Wins   int `json:"wins"`
	Losses int `json:"losses"`
	Ot     int `json:"ot"`
}

// Status represents the status in the JSON response from the NHL API.
type Status struct {
	DetailedState string `json:"detailedState"`
}

// Venue represents the venue in the JSON response from the NHL API.
type Venue struct {
	Name string `json:"name"`
}
