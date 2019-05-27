package dynamo

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

/**
* Struct for DynamoDB Tables - Following Example
{
   "name":"Movies",
   "attributes":[
      {
         "attributeName":"Year",
         "attributeType":"N"
      },
      {
         "attributeName":"Titles",
         "attributeType":"S"
      }
   ],
   "keys":[
      {
         "keyName":"Year",
         "keyType":"Hash"
      },
      {
         "keyName":"Year",
         "keyType":"RANGE"
      }
   ],
   "throughput":{
      "read":10,
      "write":10
   }
}
*/ /*
type Table struct {
	Name       string `json:"name"`
	Attributes []struct {
		AttributeName string `json:"attributeName"`
		AttributeType string `json:"attributeType"`
	} `json:"attributes"`
	Keys []struct {
		KeyName string `json:"keyName"`
		KeyType string `json:"keyType"`
	} `json:"keys"`
	Throughput struct {
		Read  int `json:"read"`
		Write int `json:"write"`
	} `json:"throughput"`
}
*/
type Table struct {
	Name       string       `json:"name"`
	Attributes []*Attribute `json:"attributes"`
	Keys       []*Key       `json:"keys"`
	Throughput *Throughput  `json:"throughput"`
}

type Attribute struct {
	Name string `json:"attributeName"`
	Type string `json:"attributeType"`
}

type Key struct {
	Name string `json:"keyName"`
	Type string `json:"keyType"`
}

type Throughput struct {
	Read  int64 `json:"read"`
	Write int64 `json:"write"`
}

func (table Table) toCreateTableInput() *dynamodb.CreateTableInput {
	attributeDefinitions := make([]*dynamodb.AttributeDefinition, 0)
	for _, attribute := range table.Attributes {
		attdef := dynamodb.AttributeDefinition{
			AttributeName: &attribute.Name,
			AttributeType: &attribute.Type,
		}
		attributeDefinitions = append(attributeDefinitions, &attdef)
	}
	keySchemas := make([]*dynamodb.KeySchemaElement, 0)
	for _, key := range table.Keys {
		keyelem := dynamodb.KeySchemaElement{
			AttributeName: &key.Name,
			KeyType:       aws.String(strings.ToUpper(key.Type)),
		}
		keySchemas = append(keySchemas, &keyelem)
	}
	provisionedThroughput := &dynamodb.ProvisionedThroughput{
		ReadCapacityUnits:  &table.Throughput.Read,
		WriteCapacityUnits: &table.Throughput.Write,
	}
	return &dynamodb.CreateTableInput{
		AttributeDefinitions:  attributeDefinitions,
		KeySchema:             keySchemas,
		ProvisionedThroughput: provisionedThroughput,
		TableName:             &table.Name,
	}
}
