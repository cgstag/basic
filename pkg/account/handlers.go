package account

import (
	"basic/api"
	"basic/pkg/dynamo"
	"net/http"

	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
)

type resource struct {
	db  *dynamo.DB
	log *zap.SugaredLogger
}

func ServeResources(env *api.Env, router *echo.Group) {
	r := &resource{env.Db, env.Log}

	rg := router.Group("/account")

	rg.GET("/list", r.listTables)
	rg.GET("/random", r.getRandom)
	rg.GET("/:uuid", r.getAccount)
	rg.GET("/:uuid/balance", r.getConsolidatedPosition)
	rg.POST("", r.createAccount)
}

func (r *resource) listTables(c echo.Context) error {
	return ListTable(r.db)
}

func (r *resource) getRandom(c echo.Context) error {
	account, err := NewRandomAccount()
	if err != nil {
		r.log.Error("Error generating random account")
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	if account != nil && account.isBlacklisted() {
		return c.JSON(http.StatusForbidden, account)
	}
	return c.JSON(http.StatusOK, account)

}

func (r *resource) getAccount(c echo.Context) error {
	return c.JSON(http.StatusOK, "Account Detail")
}

func (r *resource) getConsolidatedPosition(c echo.Context) error {
	return c.JSON(http.StatusOK, "Consolidated Position")
}

func (r *resource) createAccount(c echo.Context) error {
	// Bind Account Payload
	acc := new(Account)
	if err := c.Bind(acc); err != nil {
		return c.JSON(http.StatusBadRequest, "Couldn't bind JSON payload to account Struct")
	}

	// Verify Data
	if !acc.isValid() {
		return c.JSON(http.StatusBadRequest, "This account is invalid")
	}
	if acc.isBlacklisted() {
		return c.JSON(http.StatusForbidden, "This account is blacklisted")
	}

	// Create Account

	return c.JSON(http.StatusOK, "Account created")
}
