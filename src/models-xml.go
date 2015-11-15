package main

type xmlSectionsContainer struct {
	Sections []xmlSection `xml:"Directory"`
}

type xmlSection struct {
	Key   string `xml:"key,attr"`
	Type  string `xml:"type,attr"`
	Title string `xml:"title,attr"`
}

type xmlMoviesContainer struct {
	Movies []xmlMovieKey `xml:"Video"`
}

type xmlMovieKey struct {
	Key   string `xml:"ratingKey,attr"`
	Title string `xml:"title,attr"`
}

type xmlMovieContainer struct {
	Movie xmlMovie `xml:"Video"`
}

type xmlMovie struct {
	Key                   string  `xml:"ratingKey,attr"`
	Type                  string  `xml:"type,attr"`
	Title                 string  `xml:"title,attr"`
	Studio                string  `xml:"studio,attr"`
	ContentRating         string  `xml:"contentRating,attr"`
	Summary               string  `xml:"summary,attr"`
	Tagline               string  `xml:"tagline,attr"`
	Rating                float32 `xml:"rating,attr"`
	ViewCount             uint    `xml:"viewCount,attr"`
	Year                  uint    `xml:"year,attr"`
	Duration              uint    `xml:"duration,attr"`
	LastViewedAt          string  `xml:"lastViewedAt,attr"`
	OriginallyAvailableAt string  `xml:"originallyAvailableAt,attr"`
	AddedAt               string  `xml:"addedAt,attr"`
	UpdatedAt             string  `xml:"updatedAt,attr"`
}

type xmlShowsContainer struct {
	Shows []xmlShowKey `xml:"Directory"`
}

type xmlShowKey struct {
	Key   string `xml:"ratingKey,attr"`
	Title string `xml:"title,attr"`
}

type xmlShowContainer struct {
	Show xmlShow `xml:"Directory"`
}

type xmlShow struct {
	Key                   string  `xml:"ratingKey,attr"`
	Type                  string  `xml:"type,attr"`
	Title                 string  `xml:"title,attr"`
	Studio                string  `xml:"studio,attr"`
	TitleSort             string  `xml:"titleSort,attr"`
	ContentRating         string  `xml:"contentRating,attr"`
	Summary               string  `xml:"summary,attr"`
	Rating                float32 `xml:"rating,attr"`
	ViewCount             uint    `xml:"viewCount,attr"`
	Year                  uint    `xml:"year,attr"`
	Duration              uint    `xml:"duration,attr"`
	NumSeasons            uint    `xml:"childCount,attr"`
	NumEpisodes           uint    `xml:"leafCount,attr"`
	NumViewedEpisodes     uint    `xml:"viewedLeafCount,attr"`
	LastViewedAt          string  `xml:"lastViewedAt,attr"`
	OriginallyAvailableAt string  `xml:"originallyAvailableAt,attr"`
	AddedAt               string  `xml:"addedAt,attr"`
	UpdatedAt             string  `xml:"updatedAt,attr"`
}

type xmlSeasonsContainer struct {
	Seasons []xmlSeason `xml:"Directory"`
}

type xmlSeason struct {
	Key               string `xml:"ratingKey,attr"`
	ShowKey           string `xml:"parentRatingKey,attr"`
	Type              string `xml:"type,attr"`
	Title             string `xml:"title,attr"`
	Summary           string `xml:"summary,attr"`
	SeasonIndex       int    `xml:"parentIndex,attr"`
	ViewCount         uint   `xml:"viewCount,attr"`
	NumEpisodes       uint   `xml:"leafCount,attr"`
	NumViewedEpisodes uint   `xml:"viewedLeafCount,attr"`
}

type xmlEpisodesContainer struct {
	Episodes []xmlEpisodeKey `xml:"Video"`
}

type xmlEpisodeKey struct {
	Key   string `xml:"ratingKey,attr"`
	Title string `xml:"title,attr"`
}

type xmlEpisodeContainer struct {
	Episode xmlEpisode `xml:"Video"`
}

type xmlEpisode struct {
	Key                   string  `xml:"ratingKey,attr"`
	ShowKey               string  `xml:"grandparentRatingKey,attr"`
	SeasonKey             string  `xml:"parentRatingKey,attr"`
	SeasonIndex           int     `xml:"parentIndex,attr"`
	EpisodeIndex          int     `xml:"index,attr"`
	Type                  string  `xml:"type,attr"`
	Title                 string  `xml:"title,attr"`
	ContentRating         string  `xml:"contentRating,attr"`
	Summary               string  `xml:"summary,attr"`
	Rating                float32 `xml:"rating,attr"`
	ViewCount             uint    `xml:"viewCount,attr"`
	Year                  uint    `xml:"year,attr"`
	Duration              uint    `xml:"duration,attr"`
	OriginallyAvailableAt string  `xml:"originallyAvailableAt,attr"`
	LastViewedAt          string  `xml:"lastViewedAt,attr"`
	AddedAt               string  `xml:"addedAt,attr"`
	UpdatedAt             string  `xml:"updatedAt,attr"`
}
