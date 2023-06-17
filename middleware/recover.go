package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/khoindq/tcbHomework/common"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Recover is a middleware function for Gin framework that recovers from panics.

		// Defer a function to handle any panic that occurs within the handler.
		defer func() {
			if err := recover(); err != nil {
				// If a panic occurs, set the Content-Type header to "application/json".
				c.Header("Content-Type", "application/json")

				// Check if the panic value is of type *common.AppError.
				if appErr, ok := err.(*common.AppError); ok {
					// If it is, abort the request and respond with the AppError as JSON.
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					// In case you want to analyze the error with the Gin framework, you can uncomment the line below.
					// panic(err)
					return
				}

				// If the panic value is not an *common.AppError, create an internal server error AppError.
				appErr := common.ErrInternal(err.(error))

				// Abort the request and respond with the internal server error as JSON.
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				// In case you want to analyze the error with the Gin framework, you can uncomment the line below.
				// panic(err)
				return
			}
		}()

		// Call the next middleware or handler in the chain.
		c.Next()
	}
}
