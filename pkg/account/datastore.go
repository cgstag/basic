package account

import (
	"basic/pkg/dynamo"
)

func ListTable(db *dynamo.DB) ([]*dynamo.Table, error) {
	result, err := dynamo.ListTables(db)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CreateTable(db *dynamo.DB, table *dynamo.Table) error {
	if err := dynamo.CreateTable(db, table); err != nil {
		return err
	}
	return nil
}
