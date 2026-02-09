// shared/server/response.go
package server

import "github.com/gin-gonic/gin"

type Response struct {
	Data  any    `json:"data,omitempty"`
	Error *Error `json:"error,omitempty"`
	Meta  *Meta  `json:"meta,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}

func OK(c *gin.Context, status int, data any) {
	c.JSON(status, Response{Data: data})
}

func OKWithMeta(c *gin.Context, status int, data any, meta Meta) {
	c.JSON(status, Response{Data: data, Meta: &meta})
}

func Fail(c *gin.Context, status int, code string, message string) {
	c.JSON(status, Response{Error: &Error{Code: code, Message: message}})
}

func FailWithDetails(c *gin.Context, status int, code string, message string, details any) {
	c.JSON(status, Response{Error: &Error{Code: code, Message: message, Details: details}})
}
