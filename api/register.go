package api

import "github.com/gin-gonic/gin"

func RegisterRouter(e *gin.Engine) {
	e.GET("/send/:email", SendMail)
	e.GET("/verify/:id/:code", Verify)
}
