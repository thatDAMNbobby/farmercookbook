package clients

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/thatDAMNbobby/farmercookbook/clients/database"
	"github.com/thatDAMNbobby/farmercookbook/clients/elasticsearch"
	"github.com/thatDAMNbobby/farmercookbook/clients/render"
	"github.com/thatDAMNbobby/farmercookbook/config"
	renderLib "github.com/unrolled/render"
)

type Clients struct {
	Elasticsearch elasticsearch.Elasticsearch
	DynamoDB      database.Database
	Render        render.Render
}

func New(config *config.Config) *Clients {
	return &Clients{
		Elasticsearch: elasticsearch.New(
			&elasticsearch.Config{
				Address:          config.Elasticsearch.Address,
				Sniff:            config.Elasticsearch.Sniff,
				IndexRetryMinS:   config.Elasticsearch.IndexRetryMinS,
				IndexRetryMaxS:   config.Elasticsearch.IndexRetryMaxS,
				SearchRetryMinMS: config.Elasticsearch.SearchRetryMinMS,
				SearchRetryMaxMS: config.Elasticsearch.SearchRetryMaxMS,
			}),
		DynamoDB: database.New(buildDynamoDBConfig(config.Database, "something")),
		Render:   render.New(&render.Deps{Render: renderLib.New()}),
	}
}

func buildDynamoDBConfig(config *config.DatabaseConfig, name string) database.Config {
	return database.Config{
		Name: name,
		AWSConfig: &aws.Config{
			Region: &config.Region,
		},
	}
}
