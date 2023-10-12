package middleware

import "github.com/gin-gonic/gin"

type HandlerMiddleware func(engine *gin.Engine)
