package router

import (
	"time"

	"github.com/GabrielPereira187/go-dynamo/configuration"
	"github.com/GabrielPereira187/go-dynamo/handler"
	"github.com/GabrielPereira187/go-dynamo/metrics"
	"github.com/GabrielPereira187/go-dynamo/utils"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func initializeRoutes(router *gin.Engine) {

	logger, _ := zap.NewProduction()

	cfg, err := configuration.GetAwsConfig()

	if err != nil {
		panic(err)
	}

	apiConfig := utils.ApiConfig{
		DB:        dynamodb.NewFromConfig(cfg),
		Logger:    logger,
		SesClient: sesv2.NewFromConfig(cfg),
	}

	v1 := router.Group("/api/v1")
	{
		v1.GET("/info", func(ctx *gin.Context) {
			handler.GetDeviceInformation(ctx, &apiConfig)
		})
	}

	const collectionConcurrency = 10
	const collectionInterval = time.Second * 20

	go metrics.StartSendingMetrics(collectionConcurrency, collectionInterval, &apiConfig)

}
