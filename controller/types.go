package moviecontroller

import (
	movieservice "github.com/movie-api/service"
)

const (
	packagePrefix = "controller.rest.movie"
)

type movieController struct {
	prefix       string
	movieService movieservice.MovieService
}
