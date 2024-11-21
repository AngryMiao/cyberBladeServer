package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	successMsg = "Success"
)

type RespCommon struct {
	Message string `json:"message"`
}

func BaseResponse(c *gin.Context, statusCode int, obj interface{}) {
	c.JSON(statusCode, obj)
}

func OK(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}

func Create(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusCreated, obj)
}

func Delete(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func OkSuccess(c *gin.Context) {
	OK(c, RespCommon{
		Message: successMsg,
	})
}
