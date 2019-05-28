package account

import (
	"basic/api"

	"go.uber.org/zap"

	"github.com/guregu/dynamo"
	"github.com/labstack/echo/v4"
)

type resource struct {
	db  *dynamo.DB
	log *zap.SugaredLogger
}

func ServeResources(env *api.Env, router *echo.Group) {
	r := &resource{env.Db, env.Log}

	rg := router.Group("/account")

	// META - meta.go
	rg.GET("/getTable/:tableName", r.getTable)
	rg.GET("/listTables", r.listTables)
	rg.POST("/createTable", r.createTable)
	rg.GET("/random", r.random)
	rg.POST("/populate", r.populate)

	// CRUD ACCOUNT - account.go
	rg.POST("", r.create)
	rg.GET("/:uuid", r.read)
	rg.PUT("/:uuid", r.update)
	rg.DELETE("/:uuid", r.delete)

	// CRUD BALANCE - balance.go
	rg.GET("/:uuid/balance", r.balanceRead)
	rg.PATCH("/:uuid/balance", r.balanceUpdate)
}
