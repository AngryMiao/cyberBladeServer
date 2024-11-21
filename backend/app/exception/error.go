package exception

import (
	"fmt"
)

type Error struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e Error) Error() string {
	return fmt.Sprintf("error code: %v,error message: %s", e.Code, e.Message)
}
