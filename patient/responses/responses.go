package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// return success
func Success(c *gin.Context, code int, msg interface{}, data interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": code, // code
		"msg":  msg,  // message
		"data": data, // data
	})
	return
}

// return error
func Error(c *gin.Context, httpCode int, code int, msg string, jsonStr interface{}) {
	c.JSON(httpCode, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": jsonStr,
	})
	return
}
