package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, err error) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    err.Error(),
	})
}


func ErrorForbiddenReponse(c *gin.Context, code int, errorStr string) {
	c.AbortWithStatusJSON(http.StatusForbidden, ResponseData{
		Code: code,
		Message: msg[code],
		Data: errorStr,
	})
}

func ErrorUnAuthorizedResponse(c *gin.Context, code int, errorStr string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, ResponseData{
		Code: code,
		Message: msg[code],
		Data: errorStr,
	})
}

func ErrorInternalServerError(c *gin.Context, code int, errorStr string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, ResponseData{
		Code: code,
		Message: msg[code],
		Data: errorStr,
	})
}