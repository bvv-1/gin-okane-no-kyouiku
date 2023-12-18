package main

import (
	"net/http"

	"time"

	_ "gin-okane-no-kyouiku/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gin-okane-no-kyouiku/controllers"
	"gin-okane-no-kyouiku/db"
	"gin-okane-no-kyouiku/utils"
)

// @title okane no kyouiku API
// @version 1.0
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	utils.LoadEnv()
	db.InitDB()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173", "https://okane-no-kyouiku.onrender.com"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// - GET /api/v1/goals: 設定したゴールを返す
	// - POST /api/v1/goals: ゴールとタスクの情報を受け取って、DBにセットする
	// - GET /api/v1/plans/suggested: DBのゴールとタスクの情報から、お手伝いプランを提案する
	// - PUT /api/v1/plans/suggested: 提案をacceptするならプランIDとタスクIDの情報を受け取って、お手伝いプランをin progress状態に設定する
	// - GET /api/v1/goals/progress: 設定したゴールと、溜まったポイントと、on trackかどうかを返す
	// - GET /api/v1/plans: 設定したプランを返す
	// - GET /api/v1/plans/today: 設定したプランのうち、本日のプランを返す
	// - POST ??: 本日のプランの達成状況をDBにセットする

	r.GET("/", helloWorld)
	r.GET("/api/v2/goals", controllers.GetGoal)
	r.POST("/api/v1/goals", controllers.SetGoalAndTasks)
	r.GET("/api/v1/plans/suggested", controllers.GetSuggestedPlans)
	r.PUT("/api/v1/plans/suggested", controllers.AcceptSuggestedPlans)
	r.GET("/api/v1/goals/progress", controllers.CheckProgress)
	r.GET("/api/v1/plans", controllers.GetPlans)
	r.GET("/api/v2/plans/today", controllers.GetTodayPlan)
	r.POST("/api/v2/plans/today", controllers.SubmitTodayProgress)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

// helloWorld godoc
// @Summary Hello Worldのエンドポイント
// @Description GETリクエストに対して{"message": "Hello, World!"}を返す
// @ID helloWorld
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessResponse
// @Router / [get]
func helloWorld(c *gin.Context) {
	response := utils.SuccessResponse{Message: "Hello, World!"}
	c.JSON(http.StatusOK, response)
}
