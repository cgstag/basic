package account

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/**
 * Read an account's balance
 * @param :uuid string
 */
func (r *resource) balanceRead(c echo.Context) error {
	return c.JSON(http.StatusOK, "Consolidated Position")
}

/**
 * Update an account's balance
 * @param :uuid string
 */
func (r *resource) balanceUpdate(c echo.Context) error {
	return c.JSON(http.StatusOK, "Consolidated Position")
}