package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa o Router utilizando as configurações padrões do Gin
	router := gin.Default()

	// Define as rotas
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Inicia o servidor
	router.Run(":8080")
}
