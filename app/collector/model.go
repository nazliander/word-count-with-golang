package collector

// TextData contains Text URL and Counted Words data for collection call
type TextData struct {
	URL        string         `json:"URL"`
	WordCounts map[string]int `json:"WordCounts"`
}

// TextRequest contains URL data for request call
type TextRequest struct {
	URL string `json:"URL"`
}
