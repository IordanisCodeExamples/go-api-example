package transportkafka

// Movie represents the movie object send through kafka
type Movie struct {
	Title            string   `json:"title"`
	Year             int      `json:"year"`
	Duration         int      `json:"duration"`
	Director         string   `json:"director"`
	Cast             []string `json:"cast"`
	Genre            []string `json:"genre"`
	Synopsis         string   `json:"synopsis"`
	BoxOfficeRevenue float64  `json:"box_office_revenue"`
}
