package model

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Meta struct {
		Message string `json:"message"`
		Status  string `json:"status"`
		Code    int    `json:"code"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

// SuccessData write response
func (r *Response) SuccessData(data interface{}, dataAdditional interface{}) (int, *Response) {
	r.Meta.Message = "data berhasil diproses"
	r.Meta.Status = "success"
	r.Meta.Code = http.StatusOK

	r.Data = data
	return http.StatusOK, r
}

func (r *Response) GinSuccessData(c *gin.Context, data interface{}) {
	c.JSON(r.SuccessData(data, nil))
}

func (r *Response) GinSuccessDataWithAdditional(c *gin.Context, data interface{}, dataAdditional interface{}) {
	c.JSON(r.SuccessData(data, dataAdditional))
}

// SuccessCreated write response
func (r *Response) SuccessCreated(data interface{}) (int, *Response) {
	r.Meta.Message = "data berhasil dibuat"
	r.Meta.Status = "success"
	r.Meta.Code = http.StatusCreated

	r.Data = data
	return http.StatusCreated, r
}

// SuccessUpdated write response
func (r *Response) SuccessUpdated(data interface{}) (int, *Response) {
	r.Meta.Message = "data berhasil diubah"
	r.Meta.Status = "success"
	r.Meta.Code = http.StatusOK

	r.Data = data
	return http.StatusOK, r
}

// ErrorNotAuthorized write response
func (r *Response) ErrorNotAuthorized() (int, *Response) {
	r.Meta.Message = "tidak terautentikasi, silahkan login kembali"
	r.Meta.Status = "Not Authorized"
	r.Meta.Code = http.StatusUnauthorized

	r.Data = nil
	return http.StatusUnauthorized, r
}

func (r *Response) GinErrorNotAuthorized(c *gin.Context) {
	c.AbortWithStatusJSON(r.ErrorNotAuthorized())
}

func (r *Response) GinErrorNotAuthorizedWithError(c *gin.Context, err error) {
	c.Error(err)
	c.AbortWithStatusJSON(r.ErrorNotAuthorized())
}

// ErrorForbidden write response
func (r *Response) ErrorForbidden(message string) (int, *Response) {
	r.Meta.Message = message
	r.Meta.Status = "Forbidden"
	r.Meta.Code = http.StatusForbidden

	r.Data = nil
	return http.StatusForbidden, r
}

func (r *Response) GinErrorForbidden(c *gin.Context, message string) {
	c.Error(errors.New(message))
	c.AbortWithStatusJSON(r.ErrorForbidden(message))
}

// ErrorDataNotFound write response
func (r *Response) ErrorDataNotFound(message string) (int, *Response) {
	r.Meta.Message = "data tidak ditemukan"
	r.Meta.Status = "not found"
	r.Meta.Code = http.StatusNotFound

	r.Data = nil
	return http.StatusNotFound, r
}

func (r *Response) ErrorDataNotFoundFromError(err error) (int, *Response) {
	return r.ErrorDataNotFound(err.Error())
}

func (r *Response) GinErrorDataNotFound(c *gin.Context, err error) {
	c.Error(err)
	c.AbortWithStatusJSON(r.ErrorDataNotFoundFromError(err))
}

// ErrorInternalServer write response
func (r *Response) ErrorInternalServer(err error) (int, *Response) {
	log.Printf("Error Response: %v \n", err)

	r.Meta.Message = err.Error()
	r.Meta.Status = "internal server error"
	r.Meta.Code = http.StatusInternalServerError

	r.Data = nil
	return http.StatusInternalServerError, r
}

func (r *Response) GinErrorInternalServer(c *gin.Context, err error) {
	c.Error(err)
	c.AbortWithStatusJSON(r.ErrorInternalServer(err))
}

// ErrorBadRequestFromError write response with error argument
func (r *Response) ErrorBadRequest(message string) (int, *Response) {
	r.Meta.Message = message
	r.Meta.Status = "Bad Request"
	r.Meta.Code = http.StatusBadRequest

	r.Data = nil
	return http.StatusBadRequest, r
}

func (r *Response) ErrorBadRequestFromError(err error) (int, *Response) {
	return r.ErrorBadRequest(err.Error())
}

func (r *Response) GinErrorBadRequest(c *gin.Context, err error) {
	c.Error(err)
	c.AbortWithStatusJSON(r.ErrorBadRequestFromError(err))
}
