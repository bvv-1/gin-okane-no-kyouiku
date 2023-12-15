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

```
go mod init gin-okane-no-kyouiku
go get -u github.com/gin-gonic/gin
go mod tidy
```

### swaggo (swagger) の setup

```
go install github.com/swaggo/swag/cmd/swag@latest
export PATH=$(go env GOPATH)/bin:$PATH

go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
```

`main.go`の import に`_ "gin-okane-no-kyouiku/docs"` `swaggerFiles "github.com/swaggo/files"` `ginSwagger "github.com/swaggo/gin-swagger"`を追加

```
swag init
swag fmt
```

### 実行

```
go run main.go
```

<p align="right">(<a href="#top">トップへ</a>)</p>
