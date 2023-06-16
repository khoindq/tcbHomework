package main

import (
	"github.com/gin-gonic/gin"
	docs "github.com/khoindq/tcbHomework/docs/poolservice"
	poolgin "github.com/khoindq/tcbHomework/module/pool/transport/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		pool := v1.Group("/pool")
		{
			pool.POST("/insert", poolgin.InsertAppendPoolHandler())
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")

}
