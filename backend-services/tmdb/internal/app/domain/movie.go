package domain

type Movie struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	OriginalTitle string  `json:"original_title"`
	ReleaseDate   string  `json:"release_date"`
	Popularity    float64 `json:"popularity"`
	PosterPath    string  `json:"poster_path"`
	BackdropPath  string  `json:"backdrop_path"`
}

type MovieSearchResponse struct {
	Page         int      `json:"page"`
	Results      []*Movie `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
