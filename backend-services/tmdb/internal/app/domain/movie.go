package domain

type Movie struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	OriginalTitle string  `json:"original_title"`
	ReleaseDate   string  `json:"release_date"`
	HomePage      string  `json:"homepage"`
	Popularity    float64 `json:"popularity"`
}
