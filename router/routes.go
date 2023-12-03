package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/leobelini-studies/arthur-404-dev_api_moderna_golang/docs"
	"github.com/leobelini-studies/arthur-404-dev_api_moderna_golang/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(r *gin.Engine) {
	// Initialize Handler
	handler.Initialize()
	basePath := "/api/v1"

	v1 := r.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

	// Initialize Swagger
	docs.SwaggerInfo.BasePath = basePath
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
