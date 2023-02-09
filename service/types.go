package movieservice

import "github.com/movie-api/repository"

const (
	packagePrefix = "movieService"
)

type movieService struct {
	prefix          string
	movieRepository repository.MovieRepository
}
