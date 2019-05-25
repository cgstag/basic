package account

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
)

type accountResource struct {
	log *zap.SugaredLogger
}

func ServeResources(router *echo.Group, log *zap.SugaredLogger) {
	r := &accountResource{log}

	rg := router.Group("/account")

	rg.GET("/random", r.getRandom)
}

func (r *accountResource) getRandom(c echo.Context) error {
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

func (r *accountResource) getAccount(c echo.Context) error {
	return nil
}

func (r *accountResource) getConsolidatedPosition(c echo.Context) error {
	return nil
}
