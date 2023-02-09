package movieservice

import (
	"context"
	"fmt"

	"github.com/movie-api/model"
	"github.com/spf13/cast"
)

func (p *movieService) CreateMovie(ctx context.Context, props model.Movie) (resp uint64, err error) {
	prefix := p.prefix + ".CreateMovie"

	resp, err = p.movieRepository.CreateMovie(ctx, props)
	if err != nil {
		return 0, fmt.Errorf("[%s] error while create movie: %+v", prefix, err)
	}

	return resp, nil
}

type PropsMovieUpdateByFields struct {
	Data map[string]interface{}
}

func (p *movieService) UpdateMovieByFields(ctx context.Context, props PropsMovieUpdateByFields, MovieID uint64) error {
	prefix := p.prefix + ".UpdateMovieByFields"
	var (
		fields       []string
		updatedMovie model.Movie
	)

	for field, val := range props.Data {
		switch field {
		case "title":
			newValue, err := cast.ToStringE(val)
			if err != nil {
				continue
			}
			updatedMovie.Title = newValue
			fields = append(fields, field)

		case "description":
			newValue, err := cast.ToStringE(val)
			if err != nil {
				continue
			}
			updatedMovie.Description = newValue
			fields = append(fields, field)

		case "artist":
			newValue, err := cast.ToStringE(val)
			if err != nil {
				continue
			}
			updatedMovie.Artist = newValue
			fields = append(fields, field)

		case "genre":
			newValue, err := cast.ToStringE(val)
			if err != nil {
				continue
			}
			updatedMovie.Genre = newValue
			fields = append(fields, field)

		case "duration":
			newValue, err := cast.ToUint64E(val)
			if err != nil {
				continue
			}
			updatedMovie.Duration = newValue
			fields = append(fields, field)
		}
	}

	if len(fields) > 0 {
		if err := p.movieRepository.UpdateMovieByFields(ctx, updatedMovie, fields, MovieID); err != nil {
			return fmt.Errorf("[%s] error while UpdateMovieByFields movie: %+v", prefix, err)
		}
	}

	return nil
}
