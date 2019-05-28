package account

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type MovieTable struct {
	UserID string    `dynamo:"ID,hash" index:"Seq-ID-index,range"`
	Time   time.Time `dynamo:",range"`
	Seq    int64     `localIndex:"ID-Seq-index,range" index:"Seq-ID-index,hash"`
	UUID   string    `index:"UUID-index,hash"`
}

type DeleteTable struct {
	TableName string
}

/**
 * List DynamoDB Tables
 */
func (r *resource) listTables(c echo.Context) error {
	tables, err := r.db.ListTables().All()
	if err != nil {
		r.log.Error("Error retrieving DynamoDB Tables")
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tables)
}

// Create DynamoDB Table
func (r *resource) createTable(c echo.Context) error {
	err := r.db.CreateTable("Movie", MovieTable{}).Run()
	if err != nil {
		r.log.Error("Error creating DynamoDB Table")
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "Successfully Created DynamoDB Table")
}

// Populate DynamoDB Tables
func (r *resource) populate(c echo.Context) error {
	return c.JSON(http.StatusCreated, "")
}
