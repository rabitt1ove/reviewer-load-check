package server

import (
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "reviewer-load-check/controllers"
)

func Serve() {
    // デフォルトのミドルウェアでginのルーターを作成
    router := gin.Default()

    // 静的ファイルをインポート
    router.Static("/assets", "./assets")

    // URLへのアクセスに対して静的ページを返す
    router.StaticFS("/views", http.Dir("./views"))

    // ルーティング設定
    setupRoutes(router)

    // サーバー起動
    if err := router.Run(":8081"); err != nil {
        log.Fatal("Server Run Failed.: ", err)
    }
}

// ルーティング設定
func setupRoutes(r *gin.Engine) {
    r.GET("/", controllers.Index)

    v1 := r.Group("/v1")
    {
        // チームメンバーのJSONを返す
        v1.GET("/member", controllers.GetTeamMembers)

        // PR情報のJSONを返す
        v1.POST("/search", controllers.GetPullRequests)
    }
}
