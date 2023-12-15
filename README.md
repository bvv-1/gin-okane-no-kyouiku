<div id="top"></div>

## 使用技術一覧

<!-- シールド一覧 -->
<!-- 該当するプロジェクトの中から任意のものを選ぶ-->
<p style="display: inline">
  <!-- バックエンドのフレームワーク一覧 -->
  <img src="https://img.shields.io/badge/-Gin-000000.svg?logo=gin&style=for-the-badge">
  <!-- バックエンドの言語一覧 -->
  <img src="https://img.shields.io/badge/-Go-9DCCE0.svg?logo=go&style=for-the-badge">
  <!-- ミドルウェア一覧 -->
  <!-- インフラ一覧 -->
</p>

## 環境構築メモ

### Gin の setup

インストール

```
go mod init gin-okane-no-kyouiku
go get -u github.com/gin-gonic/gin
go mod tidy
```

### swaggo (swagger) の setup

インストール

```
go install github.com/swaggo/swag/cmd/swag@latest
export PATH=$(go env GOPATH)/bin:$PATH

go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
```

`main.go`の import に`_ "gin-okane-no-kyouiku/docs"` `swaggerFiles "github.com/swaggo/files"` `ginSwagger "github.com/swaggo/gin-swagger"`を追加

以下を`main.go`に追加

```{go}
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
```

swagger を生成

```
swag init
swag fmt
```

### 実行

```
go run main.go
```

<p align="right">(<a href="#top">トップへ</a>)</p>
