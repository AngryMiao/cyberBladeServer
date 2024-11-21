package request

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func CurrentUserID(c *gin.Context) (int, error) {
	key := "user_id"
	if value, exists := c.Get(key); exists {
		return value.(int), nil
	}

	return 0, errors.New("NotFound")
}

func CurrentUserRole(c *gin.Context) (string, error) {
	key := "user_role"
	if value, exists := c.Get(key); exists {
		return value.(string), nil
	}

	return "", errors.New("NotFound")
}
