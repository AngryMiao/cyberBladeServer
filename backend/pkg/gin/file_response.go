package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FileResponse(c *gin.Context, fileName, filePath, contentType string) {
	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	c.Writer.Header().Add("Content-Type", contentType)
	c.File(filePath)
}
