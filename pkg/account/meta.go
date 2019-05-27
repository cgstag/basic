package account

import (
	"basic/pkg/dynamo"
	"net/http"

	"github.com/labstack/echo/v4"
)

// List DynamoDB Tables
func (r *resource) listTables(c echo.Context) error {
	tables, err := ListTable(r.db)
	if err != nil {
		r.log.Error("Error retrieving DynamoDB Tables")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tables)
}

// Create DynamoDB Table
func (r *resource) createTable(c echo.Context) error {
	table := new(dynamo.Table)
	if err := c.Bind(table); err != nil {
		r.log.Error("Error binding DynamoDB Table struct", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	err := CreateTable(r.db, table)
	if err != nil {
		r.log.Error("Error populating DynamoDB Tables", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "")
}

// Populate DynamoDB Tables
func (r *resource) populate(c echo.Context) error {
	table := new(dynamo.Table)
	if err := c.Bind(table); err != nil {
		r.log.Error("Error binding DynamoDB Table struct", err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}
	err := CreateTable(r.db, table)
	if err != nil {
		r.log.Error("Error populating DynamoDB Tables", err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, "")
}
