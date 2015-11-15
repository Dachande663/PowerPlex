package main

import (
	"encoding/xml"
	"fmt"
	"sync"
)

type exporter struct {
	Reporter func(string)
	Url      string
	ch       chan *PlexItem
	metadata *exporterData
	wg       *sync.WaitGroup
}

type exporterData struct {
	sync.RWMutex
	data map[string]interface{}
}

func makeExporter() exporter {

	return exporter{
		Reporter: func(msg string) {},
		Url:      "http://localhost:32400",
		ch:       make(chan *PlexItem),
		metadata: &exporterData{data: make(map[string]interface{})},
		wg:       &sync.WaitGroup{},
	}

}

func (e *exporter) Export() map[string]interface{} {

	e.Reporter("Starting exporter...")
	e.getSections()
	e.wg.Wait()
	e.Reporter("Export completed")

	e.metadata.RLock()
	data := e.metadata.data
	e.metadata.RUnlock()

	return data

}

func (e *exporter) getSections() {

	resp, err := makeRequest("http://europa.local:32400/library/sections")
	if err != nil {
		e.Reporter(err.Error())
		return
	}

	var c xmlSectionsContainer
	err = xml.Unmarshal(resp, &c)
	if err != nil {
		e.Reporter(err.Error())
		return
	}

	e.Reporter(fmt.Sprintf("Found %d sections", len(c.Sections)))

	for _, section := range c.Sections {
		e.wg.Add(1)
		go e.getSection(section)
	}

}

func (e *exporter) getSection(section xmlSection) {

	switch section.Type {
	case "movie":
		go e.getMovies(section)
	case "show":
		go e.getShows(section)
	default:
		e.Reporter(fmt.Sprintf("Unknown section type: %s", section.Type))
		e.wg.Done()
	}

}

func (e *exporter) getMovies(section xmlSection) {

	defer e.wg.Done()
	e.Reporter(fmt.Sprintf("Getting movies in %s", section.Title))

	resp, err := makeRequest("http://europa.local:32400/library/sections/" + section.Key + "/all")
	if err != nil {
		e.Reporter(err.Error())
		return
	}

	var movies xmlMoviesContainer
	err = xml.Unmarshal(resp, &movies)
	if err != nil {
		e.Reporter(err.Error())
		return
	}

	e.Reporter(fmt.Sprintf("Found %d movies in %s", len(movies.Movies), section.Title))

	for _, movie := range movies.Movies {
		e.wg.Add(1)
		go e.getMovie(movie)
	}

}

func (e *exporter) getMovie(key xmlMovieKey) {

	defer e.wg.Done()
	e.Reporter(fmt.Sprintf("Getting movie %s", key.Title))

	resp, err := makeRequest("http://europa.local:32400/library/metadata/" + key.Key)
	if err != nil {
		e.Reporter(err.Error())
		return
	}

	var container xmlMovieContainer
	err = xml.Unmarshal(resp, &container)
	if err != nil {
		e.Reporter(err.Error())
		return
	}

	if container.Movie.Key == "" {
		e.Reporter(fmt.Sprintf("Movie not found: %s", key.Title))
		return
	}

	movie := container.Movie

	e.Reporter(fmt.Sprintf("Got %s", movie.Title))

	e.metadata.Lock()
	e.metadata.data[movie.Key] = movie
	e.metadata.Unlock()

}

func (e *exporter) getShows(section xmlSection) {

	defer e.wg.Done()
	e.Reporter(fmt.Sprintf("Getting shows in %s", section.Title))

	resp, err := makeRequest("http://europa.local:32400/library/sections/" + section.Key + "/all")
	if err != nil {
		e.Reporter(err.Error())
		return
	}

	var shows xmlShowsContainer
	err = xml.Unmarshal(resp, &shows)
	if err != nil {
		e.Reporter(err.Error())
		return
	}

	e.Reporter(fmt.Sprintf("Found %d shows in %s", len(shows.Shows), section.Title))

	for _, show := range shows.Shows {
		e.wg.Add(3)
		go e.getShow(show)
		go e.getSeasons(show)
		go e.getEpisodes(show)
	}

}

func (e *exporter) getShow(key xmlShowKey) {

	defer e.wg.Done()
	e.Reporter(fmt.Sprintf("Getting show %s", key.Title))

	resp, err := makeRequest("http://europa.local:32400/library/metadata/" + key.Key)
	if err != nil {
		e.Reporter(err.Error())
		return
	}

	var container xmlShowContainer
	err = xml.Unmarshal(resp, &container)
	if err != nil {
		e.Reporter(err.Error())
		return
	}

	if container.Show.Key == "" {
		e.Reporter(fmt.Sprintf("Show not found: %s", key.Title))
		return
	}

	show := container.Show
	e.Reporter(fmt.Sprintf("Got %s", show.Title))

	e.metadata.Lock()
	e.metadata.data[show.Key] = show
	e.metadata.Unlock()

}

func (e *exporter) getSeasons(key xmlShowKey) {

	defer e.wg.Done()
	e.Reporter(fmt.Sprintf("Getting seasons for show %s", key.Title))

	resp, err := makeRequest("http://europa.local:32400/library/metadata/" + key.Key + "/children")
	if err != nil {
		e.Reporter(err.Error())
		return
	}

	var seasons xmlSeasonsContainer
	err = xml.Unmarshal(resp, &seasons)
	if err != nil {
		e.Reporter(err.Error())
		return
	}

	e.Reporter(fmt.Sprintf("Found %d seasons for %s", len(seasons.Seasons), key.Title))

	for _, season := range seasons.Seasons[1:] {
		e.metadata.Lock()
		e.metadata.data[season.Key] = season
		e.metadata.Unlock()
	}

}

func (e *exporter) getEpisodes(key xmlShowKey) {

	defer e.wg.Done()
	e.Reporter(fmt.Sprintf("Getting episodes for show %s", key.Title))

	resp, err := makeRequest("http://europa.local:32400/library/metadata/" + key.Key + "/allLeaves")
	if err != nil {
		e.Reporter(err.Error())
		return
	}

	var episodes xmlEpisodesContainer
	err = xml.Unmarshal(resp, &episodes)
	if err != nil {
		e.Reporter(err.Error())
		return
	}

	e.Reporter(fmt.Sprintf("Found %d episodes for %s", len(episodes.Episodes), key.Title))

	for _, episode := range episodes.Episodes {
		e.wg.Add(1)
		go e.getEpisode(episode)
	}

}

func (e *exporter) getEpisode(key xmlEpisodeKey) {

	defer e.wg.Done()
	e.Reporter(fmt.Sprintf("Getting episode %s", key.Title))

	resp, err := makeRequest("http://europa.local:32400/library/metadata/" + key.Key)
	if err != nil {
		e.Reporter(err.Error())
		return
	}

	var container xmlEpisodeContainer
	err = xml.Unmarshal(resp, &container)
	if err != nil {
		e.Reporter(err.Error())
		return
	}

	if container.Episode.Key == "" {
		e.Reporter(fmt.Sprintf("Episode not found: %s", key.Title))
		return
	}

	episode := container.Episode
	e.Reporter(fmt.Sprintf("Got %s", episode.Title))

	e.metadata.Lock()
	e.metadata.data[episode.Key] = episode
	e.metadata.Unlock()

}
