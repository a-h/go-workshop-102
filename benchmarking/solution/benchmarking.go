package benchmarking

import (
	_ "embed"
	"encoding/json"
	"errors"
	"strings"
)

//go:embed movies-1990s.json
var jsonFile []byte

type Film struct {
	Title   string
	Extract string
}

var ErrFilmNotFound = errors.New("film not found")

func NewFilmSlice() ([]Film, error) {
	var films []Film

	err := json.Unmarshal(jsonFile, &films)
	if err != nil {
		return nil, err
	}

	return films, nil
}

func NewFilmMap() (map[string]string, error) {
	s, err := NewFilmSlice()
	if err != nil {
		return nil, err
	}
	films := make(map[string]string)
	for _, film := range s {
		films[film.Title] = film.Extract
	}

	return films, nil
}

func SearchFilmMap(films map[string]string, title string) (Film, error) {
	v, ok := films[title]
	if !ok {
		return Film{}, ErrFilmNotFound
	}

	title = strings.ReplaceAll(title, "-", "")
	title = strings.ReplaceAll(title, ":", "")

	title = strings.ToLower(title)

	f := Film{
		Title:   title,
		Extract: v,
	}
	return f, nil
}

func SearchFilmSlice(films []Film, title string) (Film, error) {
	for _, film := range films {
		if film.Title == title {
			film.Title = strings.ReplaceAll(film.Title, "-", "")
			film.Title = strings.ReplaceAll(film.Title, ":", "")

			film.Title = strings.ToLower(film.Title)
			return film, nil
		}
	}
	return Film{}, ErrFilmNotFound
}
