package models

type QueryTerm struct {
	Term string `json:"term"`
}

type QuerySearchRequest struct {
	SearchType string    `json:"search_type"`
	Query      QueryTerm `json:"query"`
	From       int       `json:"from"`
	MaxResult  int       `json:"max_results"`
	Source     []string  `json:"_source"`
}

type QuerySearchResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		}
		Hits []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				MessageID string   `json:"message_id"`
				From      string   `json:"from"`
				To        []string `json:"to"`
				Subject   string   `json:"subject"`
				Date      string   `json:"date"`
				Body      string   `json:"body"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
