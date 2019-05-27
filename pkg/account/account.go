package account

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Generate a random Bank Account
func (r *resource) random(c echo.Context) error {
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

/**
 * Create a Bank Account
 * @param Account
 */
func (r *resource) create(c echo.Context) error {
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

/**
 * Read a Bank Account
 * @param :uuid string
 */
func (r *resource) read(c echo.Context) error {
	return c.JSON(http.StatusOK, "Account Detail")
}

/**
 * Update a Bank Account
 * @param Account
 */
func (r *resource) update(c echo.Context) error {
	return c.JSON(http.StatusOK, "Account Detail")
}

/**
 * Delete a Bank Account
 * @param :uuid string
 */
func (r *resource) delete(c echo.Context) error {
	return c.JSON(http.StatusOK, "Consolidated Position")
}
