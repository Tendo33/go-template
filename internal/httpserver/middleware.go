package httpserver

import "github.com/gin-gonic/gin"

func defaultMiddleware() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		gin.Recovery(),
	}
}
