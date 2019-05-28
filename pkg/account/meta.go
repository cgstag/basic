package account

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/labstack/echo/v4"
)

type DescribeTable struct {
	Tablename string
}

type DeleteTable struct {
	TableName string
}

func (r *resource) getTable(c echo.Context) error {
	table := new(dynamodb.DescribeTableInput).SetTableName(c.Param("tableName"))
	describe, _ := r.db.Client().DescribeTable(table)
	return c.JSON(http.StatusOK, describe)
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
	err := r.db.CreateTable("Account", Account{}).Run()
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
