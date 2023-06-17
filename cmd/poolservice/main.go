package main

import (
	"github.com/gin-gonic/gin"
	docs "github.com/khoindq/tcbHomework/docs/poolservice"
	poolgin "github.com/khoindq/tcbHomework/module/pool/transport/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Pool service
// @version         1.0 d
// @description     A tcp homework backend server

// @contact.name   Khoi Nguyen
// @contact.email  khoindq@gmail.com

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	c := poolgin.NewPoolController()

	v1 := r.Group("/api/v1")
	{
		pool := v1.Group("/pool")
		{

			pool.POST("/insertorappend", c.InsertOrAppendPoolHandler())
			quantile := pool.Group("/quantile")
			{
				quantile.POST("/get", c.GetQuantileHandler())
			}
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")

}
