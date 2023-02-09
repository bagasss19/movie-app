package movieservice

import (
	"context"

	"github.com/movie-api/model"
	"github.com/movie-api/repository"
)

func NewMovieUsecase(movieRepo repository.MovieRepository) MovieService {
	return &movieService{
		prefix:          packagePrefix + ".MovieService",
		movieRepository: movieRepo,
	}
}

type MovieService interface {
	// Admin
	CreateMovie(ctx context.Context, props model.Movie) (resp uint64, err error)
	UpdateMovieByFields(ctx context.Context, props PropsMovieUpdateByFields, MovieID uint64) error

	// User
	UserGetAllMovie(ctx context.Context, props model.PropsMovieGetAll) (resp model.GetAllMovieOutput, err error)
}
