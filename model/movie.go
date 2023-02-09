package model

type Movie struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    uint64 `json:"duration"`
	Artist      string `json:"artist"`
	Genre       string `json:"genre"`
	WatchUrl    string `json:"watch_url"`
}

const TableMovie = "movies"

func (Movie) TableName() string {
	return TableMovie
}

type PropsMovieGetAll struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Artist      string `json:"artist"`
	Genre       string `json:"genre"`

	// pagination
	Page           int64
	ContentPerPage int64
}

type GetAllMovieOutput struct {
	Data []*Movie

	// pagination
	Count   int64
	MaxPage int64
}
