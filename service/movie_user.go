package movieservice

import (
	"context"

	"github.com/movie-api/model"
	"github.com/pkg/errors"
)

func (p *movieService) UserGetAllMovie(ctx context.Context, props model.PropsMovieGetAll) (resp model.GetAllMovieOutput, err error) {
	prefix := p.prefix + ".UserGetAllMovie"

	resp, err = p.movieRepository.GetAllMovie(ctx, props)
	if err != nil {
		err = errors.Wrapf(err, "[%s]", prefix)
		return
	}

	return
}
