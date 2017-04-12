package fanarttv

import (
	"encoding/json"
	"sort"
	"strconv"
)

// ImageInfo represent an image with its associated infos
type ImageInfo struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Lang  string `json:"lang"`
	Likes int    `json:"likes"`

	// Only for the shows
	Season string `json:"season,omitempty"`
}

// UnmarshalJSON is a custom unmarshal function to handle likes as ints
func (i *ImageInfo) UnmarshalJSON(data []byte) error {
	aux := struct {
		ID     *string `json:"id"`
		URL    *string `json:"url"`
		Lang   *string `json:"lang"`
		Season *string `json:"season"`
		Likes  string  `json:"likes"`
	}{
		ID:     &i.ID,
		URL:    &i.URL,
		Lang:   &i.Lang,
		Season: &i.Season,
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Convert the likes into int
	likes, err := strconv.Atoi(aux.Likes)
	if err != nil {
		return err
	}
	i.Likes = likes

	return nil
}

// Best returns the best image based on the likes
func Best(imgs []*ImageInfo) *ImageInfo {
	size := len(imgs)
	switch size {
	case 0:
		return nil
	case 1:
		return imgs[0]
	default:
		sort.Slice(imgs, func(i, j int) bool { return imgs[i].Likes > imgs[j].Likes })
		return imgs[0]
	}
}
