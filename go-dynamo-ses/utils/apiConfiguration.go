package utils

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"go.uber.org/zap"
)

type ApiConfig struct {
	DB        *dynamodb.Client
	Logger    *zap.Logger
	SesClient *sesv2.Client
}
