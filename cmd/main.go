package main

import (
	"time"

	routeV1 "github.com/Tonie-NG/go-api-achitecture/api/route/v1"
	"github.com/Tonie-NG/go-api-achitecture/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	routerV1 := gin.Group("v1")

	routeV1.Setup(env, timeout, db, routerV1)

	gin.Run(env.ServerAddress)
}
