package fanarttv

// MovieResult represents the result for a movie
type MovieResult struct {
	Name   string `json:"name"`
	TmdbID string `json:"tmdb_id"`
	ImdbID string `json:"imdb_id"`

	HDMovieClearArt []*Resource `json:"hdmovieclearart"`
	HDMovieLogo     []*Resource `json:"hdmovielogo"`
	MovieArt        []*Resource `json:"movieart"`
	MovieBackground []*Resource `json:"moviebackground"`
	MovieBanner     []*Resource `json:"moviebanner"`
	MovieDisc       []*Resource `json:"moviedisc"`
	MovieLogo       []*Resource `json:"movielogo"`
	MoviePoster     []*Resource `json:"movieposter"`
	MovieThumb      []*Resource `json:"moviethumb"`
}

// GetMovieImages returns the images for a movie
func (c *Client) GetMovieImages(imdbID string) (*MovieResult, error) {
	url := c.Endpoint + "/movies/" + imdbID

	var mr MovieResult
	if err := c.get(url, &mr); err != nil {
		return nil, err
	}

	return &mr, nil
}
