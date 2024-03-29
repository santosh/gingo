package main

import (
	"sync"
	"time"

	"github.com/santosh/gingo/db"
	_ "github.com/santosh/gingo/docs"
	"github.com/santosh/gingo/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Gin Book Service
// @version         1.0
// @description     A book management service API in Go using Gin framework.
// @termsOfService  https://tos.santoshk.dev

// @contact.name   Santosh Kumar
// @contact.url    https://twitter.com/sntshk
// @contact.email  sntshkmr60@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go startServer(wg)

	// wait for server to start properly before seeding db
	time.Sleep(time.Second * 2)
	db.Seed()
	wg.Wait()
}

func startServer(wg *sync.WaitGroup) {
	router := routes.SetupRouter()

	db.ConnectDatabase()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8090")
}
