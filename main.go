package main

import (
	_ "gin-okane-no-kyouiku/docs"
	"net/http"

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
		// MaxAge:           12 * time.Hour,
	}))

	r.GET("/", helloWorld)
	r.GET("/api/v2/goals", controllers.GetGoal)                        // 設定したゴールを返す
	r.POST("/api/v1/goals", controllers.SetGoalAndTasks)               // ゴールとタスクの情報を受け取って、DBにセットする
	r.GET("/api/v1/plans/suggested", controllers.GetSuggestedPlans)    // DBのゴールとタスクの情報から、お手伝いプランを提案する
	r.PUT("/api/v1/plans/suggested", controllers.AcceptSuggestedPlans) // 提案をacceptするならプランIDとタスクIDの情報を受け取って、お手伝いプランをin progress状態に設定する
	r.GET("/api/v1/goals/progress", controllers.CheckProgress)         // 設定したゴールと、溜まったポイントと、on trackかどうかを返す
	r.GET("/api/v1/plans", controllers.GetPlans)                       // 設定したプランを返す
	r.GET("/api/v2/plans/today", controllers.GetTodayPlan)             // 設定したプランのうち、本日のプランを返す
	r.POST("/api/v2/plans/today", controllers.SubmitTodayProgress)     // 本日のプランの達成状況をDBにセットする

	r.POST("api/v2/register", controllers.Register)
	r.POST("api/v2/login", controllers.Login)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

// helloWorld godoc
// @Summary Hello Worldのエンドポイント
// @Description GETリクエストに対して{"message": "Hello, World!"}を返す
// @ID helloWorld
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.HTTPError
// @Router / [get]
func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, utils.SuccessResponse{Message: "Hello, World!"})
}
