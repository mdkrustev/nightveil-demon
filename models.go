package main

type RenderRequest struct {
	Version  int       `json:"version"`
	Output   string    `json:"output"`
	Settings Settings  `json:"settings"`
	Segments []Segment `json:"segments"`
}

type Settings struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Fps    int `json:"fps"`
}

type Segment struct {
	Image    string  `json:"image"`
	Audio    string  `json:"audio"`
	Duration float64 `json:"duration"`
}
