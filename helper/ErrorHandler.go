package helper

import (
	"github.com/gin-gonic/gin"
)

func SuccesResponse(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, gin.H{
		"code":   code,
		"msg":    msg,
		"reason": nil,
		"data":   data,
	})
}

func ErrorResponse(c *gin.Context, code int, reason interface{}) {
	var reasonStr string
	// convert type tp String
	switch v := reason.(type) {
	case string:
		reasonStr = v
	case error:
		reasonStr = v.Error()
	case interface{}:
		reasonStr = v.(string)
	}

	c.AbortWithStatusJSON(code, gin.H{
		"code":   code,
		"msg":    "Kesalahan Sistem (" + reasonStr + ")",
		"reason": reason,
		"data":   nil,
	})
}

func ErrorCustom(c *gin.Context, code int, msg string, reason interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"code":   code,
		"msg":    msg,
		"reason": reason,
		"data":   nil,
	})
}

func Unauthorized(c *gin.Context, code int, reason interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"code":   code,
		"msg":    "user login tidak ditemukan, silahkan login kembali",
		"reason": reason,
		"data":   nil,
	})
}

func ExpiredToken(c *gin.Context, code int, reason interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"code":   code,
		"msg":    "Masa loginmu telah habis silahkan lakukan login ulang",
		"reason": reason,
		"data":   nil,
	})
}

func Badrequest(c *gin.Context, code int, reason interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"code":   code,
		"msg":    "This Field is required. Please fill",
		"reason": reason,
		"data":   nil,
	})
}
