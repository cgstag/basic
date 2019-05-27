package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateTable(db *DB, table *Table) error {
	input := table.toCreateTableInput()
	_, err := db.CreateTable(input)
	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Created the table", table.Name)
	return nil
}

func ListTables(db *DB) ([]*Table, error) {
	input := &dynamodb.ListTablesInput{}
	output := make([]*Table, 0)
	for {
		// Get the list of tables
		result, err := db.DynamoDB.ListTables(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeInternalServerError:
					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return nil, nil
		}

		for _, n := range result.TableNames {
			output = append(output, &Table{Name: *n})
		}

		// assign the last read tablename as the start for our next call to the ListTables function
		// the maximum number of table names returned in a call is 100 (default), which requires us to make
		// multiple calls to the ListTables function to retrieve all table names
		input.ExclusiveStartTableName = result.LastEvaluatedTableName

		if result.LastEvaluatedTableName == nil {
			break
		}
	}

	return output, nil
}
