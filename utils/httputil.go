package httputil

import "github.com/gin-gonic/gin"

// https://github.com/swaggo/swag/blob/6cdaaf5c77457e82d9e0f8fccd303fefb8dc8072/example/celler/httputil/error.go#L1
// NewError creates a new HTTPError and sends it as a JSON response.
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError represents an HTTP error.
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
