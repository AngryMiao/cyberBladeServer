package exception

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	// 400
	BadRequestCode = "invalid"
	BadRequestMsg  = "invalid input"

	// 401
	UnauthorizedCode = "unauthorized"
	UnauthorizedMsg  = "authentication credentials were not provided"
	InvalidAuthMsg   = "incorrect authentication credentials"

	//	403
	ForbiddenCode = "forbidden"
	ForbiddenMsg  = "forbidden"

	//	404
	NotFoundCode = "not_found"
	NotFoundMsg  = "not found"

	//	500
	InternalServerErrorCode = "error"
	InternalServerErrorMsg  = "a server error occurred"

	// 504
	StatusGatewayTimeoutCode = "timeout"
	StatusGatewayTimeoutMsg  = "Request timeout"
)

const (
	UserWhiteListErrorCode = "not_in_whitelist"
	UserWhiteListErrorMsg  = "not in whitelist"

	AmCodeNotEnoughCode = "not_enough_am_code_to_be_used"
	AmCodeNotEnoughMsg  = "am code cannot be used with multiple products"
)

// Response
type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// 失败信息
func ErrorResponse(code, msg string) Response {
	return Response{
		Code:    code,
		Message: msg,
	}
}

func BaseException(c *gin.Context, statusCode int, obj interface{}) {
	c.JSON(statusCode, obj)
	c.Abort()
	return
}

func Exception(c *gin.Context, statusCode int, code, msg string) {
	c.JSON(statusCode, ErrorResponse(code, msg))
	c.Abort()
	return

}

func ForbiddenException(c *gin.Context, msg string) {
	statusCode := http.StatusForbidden
	code := ForbiddenCode

	if msg == "" {
		msg = ForbiddenMsg
	}

	Exception(c, statusCode, code, msg)
}

func TimeOutRequestException(c *gin.Context, msg string) {
	statusCode := http.StatusGatewayTimeout
	code := StatusGatewayTimeoutCode

	if msg == "" {
		msg = StatusGatewayTimeoutMsg
	}

	Exception(c, statusCode, code, msg)
}

func BadRequestException(c *gin.Context, msg string) {
	statusCode := http.StatusBadRequest
	code := BadRequestCode

	if msg == "" {
		msg = BadRequestMsg
	}

	Exception(c, statusCode, code, msg)

}

func UnauthorizedException(c *gin.Context, msg string) {
	statusCode := http.StatusUnauthorized
	code := UnauthorizedCode

	if msg == "" {
		msg = UnauthorizedMsg
	}

	Exception(c, statusCode, code, msg)

}

func NotFoundException(c *gin.Context, msg string) {
	statusCode := http.StatusNotFound
	code := NotFoundCode

	if msg == "" {
		msg = NotFoundMsg
	}

	Exception(c, statusCode, code, msg)

}

func InternalServerErrorException(c *gin.Context, msg string) {
	statusCode := http.StatusInternalServerError
	code := InternalServerErrorCode

	if msg == "" {
		msg = InternalServerErrorMsg
	}

	Exception(c, statusCode, code, msg)

}
