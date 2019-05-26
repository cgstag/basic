package dynamo

import (
	"basic/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DB struct {
	*dynamodb.DynamoDB
}

func NewDB(config config.DBConfig) (*DB, error) {
	var cfg aws.Config

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            *cfg.WithEndpoint(config.Endpoint).WithLogLevel(aws.LogDebugWithHTTPBody),
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	return &DB{svc}, nil
}
