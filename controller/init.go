package moviecontroller

import (
	"github.com/gin-gonic/gin"
	movieservice "github.com/movie-api/service"
)

func NewMovieController(movieService movieservice.MovieService) MovieController {
	return &movieController{
		prefix:       packagePrefix + ".programmeController",
		movieService: movieService,
	}
}

type MovieController interface {
	CreateMovie(c *gin.Context)
}
