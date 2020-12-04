package utils

import "github.com/gin-gonic/gin"

var headers = map[string]string{
	"Access-Control-Allow-Origin":      "*",
	"Access-Control-Allow-Credentials": "true",
	"Access-Control-Allow-Headers":     "Origin,Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
	"Access-Control-Allow-Methods":     "POST, OPTIONS",
}

func BuildHeaders(c *gin.Context) {
	for key, element := range headers {
		c.Header(key, element)
	}
}
