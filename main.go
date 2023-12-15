package main

import (
	"net/http"

	_ "gin-okane-no-kyouiku/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title		okane no kyouiku API
//	@version	1.0
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	r := gin.Default()

	r.GET("/ping", ping)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// GetPong godoc
// @Summary Pingのエンドポイント
// @Description Pingへのリクエストに対してJSON形式で{"message": "pong"}を返す
// @ID ping
// @Tags ping
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
