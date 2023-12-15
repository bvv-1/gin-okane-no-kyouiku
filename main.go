package main

import (
	"net/http"

	_ "gin-okane-no-kyouiku/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title okane no kyouiku API
// @version 1.0
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()

	r.GET("/", helloWorld)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// helloWorld godoc
// @Summary Hello Worldのエンドポイント
// @Description GETリクエストに対して{"message": "Hello, World!"}を返す
// @ID helloWorld
// @Tags hello
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func helloWorld(c *gin.Context) {
	data := map[string]string{"message": "Hello, World!"}
	c.JSON(http.StatusOK, data)
}
