package archive

type StatsResponse struct {
	Player Player         `json:"player"`
	Values map[string]any `json:"values"`
}

type Player struct {
	Pid       int    `json:"pid"`
	Name      string `json:"name"`
	Platform  string `json:"platform"`
	Namespace string `json:"namespace"`
	Added     string `json:"added"`
	Updated   string `json:"updated"`
}
