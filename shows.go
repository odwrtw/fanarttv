package fanarttv

// ShowResult represents the result for a show
type ShowResult struct {
	Name   string `json:"name"`
	TvdbID string `json:"thetvdb_id"`

	CharacterArt   []*Resource `json:"characterart"`
	ClearArt       []*Resource `json:"clearart"`
	ClearLogo      []*Resource `json:"clearlogo"`
	HDClearArt     []*Resource `json:"hdclearart"`
	HDTVLogo       []*Resource `json:"hdtvlogo"`
	SeasonBanner   []*Resource `json:"seasonbanner"`
	SeasonPoster   []*Resource `json:"seasonposter"`
	SeasonThumb    []*Resource `json:"seasonthumb"`
	ShowBackground []*Resource `json:"showbackground"`
	TvBanner       []*Resource `json:"tvbanner"`
	TvPoster       []*Resource `json:"tvposter"`
	TvThumb        []*Resource `json:"tvthumb"`
}

// GetShowImages returns the images for a show
func (c *Client) GetShowImages(tvdbID string) (*ShowResult, error) {
	url := c.Endpoint + "/tv/" + tvdbID

	var sr ShowResult
	if err := c.get(url, &sr); err != nil {
		return nil, err
	}

	return &sr, nil
}
