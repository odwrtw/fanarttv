package fanarttv

import (
	"encoding/json"
	"sort"
	"strconv"
)

// Resource represent a result from the API
type Resource struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Lang  string `json:"lang"`
	Likes int    `json:"likes"`

	// Only for the shows
	Season string `json:"season,omitempty"`
}

// UnmarshalJSON is a custom unmarshal function to handle likes as ints
func (r *Resource) UnmarshalJSON(data []byte) error {
	aux := struct {
		ID     *string `json:"id"`
		URL    *string `json:"url"`
		Lang   *string `json:"lang"`
		Season *string `json:"season"`
		Likes  string  `json:"likes"`
	}{
		ID:     &r.ID,
		URL:    &r.URL,
		Lang:   &r.Lang,
		Season: &r.Season,
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Convert the likes into int
	likes, err := strconv.Atoi(aux.Likes)
	if err != nil {
		return err
	}
	r.Likes = likes

	return nil
}

// Best returns the best resource based on the likes
func Best(res []*Resource) *Resource {
	size := len(res)
	switch size {
	case 0:
		return nil
	case 1:
		return res[0]
	default:
		sort.Slice(res, func(i, j int) bool { return res[i].Likes > res[j].Likes })
		return res[0]
	}
}
