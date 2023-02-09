package repository

import (
	"context"

	"github.com/movie-api/model"
	"gorm.io/gorm"
)

func NewMovieRepository(sqlConnection *gorm.DB) MovieRepository {
	return &DB{
		repoPrefix: packagePrefix + "movieRepo",
		db:         sqlConnection,
	}
}

type MovieRepository interface {
	CreateMovie(ctx context.Context, payload model.Movie) (MovieID uint64, err error)
	UpdateMovieByFields(ctx context.Context, props model.Movie, fields []string, MovieID uint64) (err error)
	GetAllMovie(ctx context.Context, props model.PropsMovieGetAll) (output model.GetAllMovieOutput, err error)
}
