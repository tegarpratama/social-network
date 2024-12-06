package helper

import "github.com/gin-gonic/gin"

type (
	Response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Data    *any   `json:"data"`
	}
)

func SetResponse(c *gin.Context, statusCode int, res *Response) {
	if statusCode >= 200 && statusCode < 300 {
		res.Status = "success"
	} else {
		res.Status = "failed"
	}

	c.JSON(statusCode, &res)
}
