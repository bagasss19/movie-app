package repository

import (
	"context"
	"fmt"

	"github.com/movie-api/model"
	"github.com/movie-api/util"
	"github.com/pkg/errors"
)

func (p *DB) CreateMovie(ctx context.Context, payload model.Movie) (MovieID uint64, err error) {
	prefix := p.repoPrefix + ".CreateMovie"
	query := p.db.WithContext(ctx)

	db := query.Create(&payload)
	if err = db.Error; err != nil {
		return 0, fmt.Errorf("[%s] %+v", prefix, err)
	}
	return uint64(payload.ID), nil
}

func (p *DB) UpdateMovieByFields(ctx context.Context, props model.Movie, fields []string, MovieID uint64) (err error) {
	prefix := p.repoPrefix + ".UpdateMovieByFields"

	query := p.db.WithContext(ctx).Select(fields)
	query = filterMovieByID(query, MovieID)

	if err := query.Updates(&props).Error; err != nil {
		return errors.Wrapf(err, "[%s] error while update", prefix)
	}

	return nil
}

func (p *DB) GetAllMovie(ctx context.Context, props model.PropsMovieGetAll) (output model.GetAllMovieOutput, err error) {
	var (
		prefix = p.repoPrefix + ".GetAllMovie"
		query  = p.db.WithContext(ctx)
		rows   []*struct {
			model.Movie
			Total int64
		}
	)

	query = query.Select("*, COUNT(1) OVER() AS total")
	query = query.Order("id DESC")

	// Selector
	if props.Page != 0 && props.ContentPerPage != 0 {
		query = util.FilterPaginate(query, props.Page, props.ContentPerPage)
	}

	if props.Title != "" {
		query = filterMovieSearchByTitle(query, props.Title)
	}

	if props.Description != "" {
		query = filterMovieSearchByDescription(query, props.Description)
	}

	if props.Artist != "" {
		query = filterMovieSearchByArtist(query, props.Artist)
	}

	if props.Genre != "" {
		query = filterMovieSearchByGenre(query, props.Genre)
	}

	// execute
	if err = query.Find(&rows).Error; err != nil {
		err = errors.Wrapf(err, "[%s]", prefix)
		return
	}

	// prepare output
	for i, row := range rows {
		if i == 0 {
			output.Count = row.Total
			output.MaxPage = util.CountMaxPage(output.Count, props.ContentPerPage)
		}

		output.Data = append(output.Data, &row.Movie)
	}

	return
}
