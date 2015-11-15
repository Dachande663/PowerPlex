package main

type PlexLibrary struct {
	Sections []*PlexSection `json:"sections"`
}

type PlexSection struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Title string `json:"title"`
}

type PlexItem struct {
	Type    string       `json:"type"`
	Key     string       `json:"key"`
	Movie   *PlexMovie   `json:"movie,omitempty"`
	Show    *PlexShow    `json:"show,omitempty"`
	Season  *PlexSeason  `json:"season,omitempty"`
	Episode *PlexEpisode `json:"episode,omitempty"`
}

type PlexMovie struct {
	Key           string    `json:"key"`
	Studio        string    `json:"studio"`
	Title         string    `json:"title"`
	ContentRating string    `json:"content_rating"`
	Summary       string    `json:"summary"`
	Tagline       string    `json:"tagline"`
	Thumb         string    `json:"thumb"`
	Art           string    `json:"art"`
	ReleaseDate   string    `json:"release_date"`
	Rating        float32   `json:"rating"`
	UserRating    float32   `json:"user_rating"`
	ViewCount     uint32    `json:"view_count"`
	ReleaseYear   uint32    `json:"release_year"`
	Duration      uint32    `json:"duration"`
	LastViewed    uint32    `json:"last_viewed"`
	Created       uint32    `json:"created"`
	Updated       uint32    `json:"updated"`
	Media         PlexMedia `json:"media"`
}

type PlexShow struct {
	Key             string  `json:"key"`
	Studio          string  `json:"studio"`
	Title           string  `json:"title"`
	ContentRating   string  `json:"content_rating"`
	Summary         string  `json:"summary"`
	Tagline         string  `json:"tagline"`
	Guid            string  `json:"guid"`
	Thumb           string  `json:"thumb"`
	Art             string  `json:"art"`
	Banner          string  `json:"banner"`
	ReleaseDate     string  `json:"release_date"`
	Rating          float32 `json:"rating"`
	UserRating      float32 `json:"user_rating"`
	ViewCount       uint32  `json:"view_count"`
	ReleaseYear     uint32  `json:"release_year"`
	Duration        uint32  `json:"duration"`
	LastViewed      uint32  `json:"last_viewed"`
	Created         uint32  `json:"created"`
	Updated         uint32  `json:"updated"`
	TotalSeasons    uint32  `json:"total_seasons"`
	TotalEpisodes   uint32  `json:"total_episodes"`
	WatchedEpisodes uint32  `json:"watched_episodes"`
}

type PlexSeason struct {
	Key         string `json:"key"`
	Title       string `json:"title"`
	ShowKey     string `json:"show_key"`
	SeasonIndex uint32 `json:"season_index"`
}

type PlexEpisode struct {
	Key          string    `json:"key"`
	Title        string    `json:"title"`
	ShowKey      string    `json:"showkey"`
	SeasonKey    string    `json:"season_key"`
	SeasonIndex  uint32    `json:"season_index"`
	EpisodeIndex uint32    `json:"episode_index"`
	Summary      string    `json:"summary"`
	Thumb        string    `json:"thumb"`
	Rating       float32   `json:"rating"`
	UserRating   float32   `json:"user_rating"`
	ViewCount    uint32    `json:"viewc_ount"`
	ReleaseYear  uint32    `json:"release_year"`
	Duration     uint32    `json:"duration"`
	LastViewed   uint32    `json:"last_viewed"`
	Created      uint32    `json:"created"`
	Updated      uint32    `json:"updated"`
	Media        PlexMedia `json:"media"`
}

type PlexMedia struct {
	Resolution string `json:"resolution"`
	Filesize   string `json:"filesize"`
}
