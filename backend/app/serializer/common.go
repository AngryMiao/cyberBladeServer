package serializer

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

type UriUUID struct {
	UUID string `uri:"uuid" binding:"required,uuid"`
}

type UriID struct {
	ID int `uri:"id" binding:"required"`
}

type UriStringID struct {
	ID string `uri:"id" binding:"required"`
}

func SuccessResponse(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}

func CreateResponse(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusCreated, obj)
}

func CommonSuccessResponse(c *gin.Context) {
	c.JSON(http.StatusOK, RespCommon{
		Message: successMsg,
	})
}
