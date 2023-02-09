package moviecontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/movie-api/model"
	"github.com/pkg/errors"
)

// Programme godoc
// @Security 		AdminAuth
// @Summary      Admin create programme
// @Description  Admin create programme
// @Tags         Admin
// @Accept json
// @Param        payload    body   model.Movie  true  "body payload"
// @Router       /admin/movie  [post]
func (p *movieController) CreateMovie(c *gin.Context) {
	prefix := p.prefix + ".CreateMovie"
	var (
		requestJson model.Movie
		response    model.Response
	)

	if err := c.BindJSON(&requestJson); err != nil {
		response.GinErrorBadRequest(c, err)
		return
	}

	resp, err := p.movieService.CreateMovie(c, requestJson)
	if err != nil {
		err = errors.Wrapf(err, "[%s] error while create movie", prefix)
		response.GinErrorInternalServer(c, err)
		return
	}

	response.GinSuccessData(c, resp)
}
