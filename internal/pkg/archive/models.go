package archive

type StatsResponse struct {
	Player Player         `json:"player"`
	Values map[string]any `json:"values"`
}

type DogtagsResponse struct {
	Player  Player         `json:"player"`
	Records []DogtagRecord `json:"records"`
}

type Player struct {
	Pid       int    `json:"pid"`
	Name      string `json:"name"`
	Platform  string `json:"platform"`
	Namespace string `json:"namespace"`
	Added     string `json:"added"`
	Updated   string `json:"updated"`
}

type DogtagRecord struct {
	Player    Player  `json:"player"`
	Timestamp float64 `json:"timestamp"`
	Rank      float64 `json:"rank"`
	Bronze    int     `json:"bronze"`
	Silver    int     `json:"silver"`
	Gold      int     `json:"gold"`
	Total     int     `json:"total"`
	Raw       string  `json:"raw"`
}
