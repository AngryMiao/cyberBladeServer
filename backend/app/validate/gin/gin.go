package gin

import (
	"angrymiao-ai/app/exception"
	"angrymiao-ai/app/log"
	"github.com/gin-gonic/gin"
)

func CheckBody(c *gin.Context, body interface{}) bool {
	if err := c.ShouldBindJSON(body); err != nil {
		log.Log.Error(err.Error())
		exception.BadRequestException(c, err.Error())
		return false
	}
	return true
}

func CheckQuery(c *gin.Context, query interface{}) bool {
	if err := c.BindQuery(query); err != nil {
		log.Log.Error(err.Error())
		exception.BadRequestException(c, err.Error())
		return false
	}
	return true
}

func CheckUri(c *gin.Context, query interface{}) bool {
	if err := c.BindUri(query); err != nil {
		log.Log.Error(err.Error())
		exception.BadRequestException(c, err.Error())
		return false
	}
	return true
}
