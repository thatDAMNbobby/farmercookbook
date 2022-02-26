package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Config struct {
	Name      string
	AWSConfig *aws.Config
}

type Database interface {
	GetName() string
}

type Impl struct {
	name string
	db   *dynamodb.DynamoDB
}

func New(config Config) Database {
	sess := session.Must(
		session.NewSessionWithOptions(
			session.Options{
				Config: *config.AWSConfig,
			},
		))

	client := dynamodb.New(sess)

	return &Impl{
		db:   client,
		name: config.Name,
	}
}

func (impl *Impl) GetName() string {
	return impl.name
}
