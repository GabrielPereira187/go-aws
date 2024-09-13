package handler

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	c "github.com/GabrielPereira187/go-dynamo/constants"
	"github.com/GabrielPereira187/go-dynamo/responses"
	"github.com/GabrielPereira187/go-dynamo/structs"
	"github.com/GabrielPereira187/go-dynamo/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	limit = 5
)

const (
	tableName = "DeviceMetrics"
)

func GetDeviceInformation(ctx *gin.Context, cfg *utils.ApiConfig) {
	var scanInput *dynamodb.ScanInput
	var startDay string
	var endDay string

	date := ctx.Query("date")
	if date != c.Empty {
		startDay, endDay = concatStartDate(date)
	}

	expAttrValues, query := buildQuery(startDay, endDay, ctx.Query("warning"), ctx.Query("id"))
	if query == c.Empty {
		scanInput = &dynamodb.ScanInput{
			TableName: aws.String(tableName),
			Limit:     aws.Int32(int32(limit)),
		}
	} else {
		scanInput = &dynamodb.ScanInput{
			TableName:                 aws.String(tableName),
			FilterExpression:          aws.String(query),
			ExpressionAttributeValues: expAttrValues,
			Limit:                     aws.Int32(int32(limit)),
		}
	}

	result, err := cfg.DB.Scan(context.TODO(), scanInput)
	if err != nil {
		responses.GetResponseQueryError(ctx, 404, "Error to get query:"+err.Error())
		return
	}

	responses.GetResponseWithJson(ctx, 200, result.Items)

}

func buildQuery(start, end, searchMode, deviceID string) (map[string]types.AttributeValue, string) {
	var builder strings.Builder
	attrValues := make(map[string]types.AttributeValue)

	if deviceID != c.Empty {
		builder.WriteString(" DeviceId = :id AND")
		attrValues[":id"] = &types.AttributeValueMemberS{Value: deviceID}
	}

	if start != c.Empty {
		builder.WriteString(" CreatedAt BETWEEN :start AND :end AND")
		attrValues[":start"] = &types.AttributeValueMemberS{Value: replace(start)}
		attrValues[":end"] = &types.AttributeValueMemberS{Value: replace(end)}
	}

	if searchMode != "" {
		builder.WriteString(" Warning = :warning AND")
		attrValues[":warning"] = &types.AttributeValueMemberS{Value: searchMode}
	}

	if builder.String() == c.Empty {
		return attrValues, builder.String()
	}

	return attrValues, builder.String()[:len(builder.String())-3]
}

func InsertDevice(cfg *utils.ApiConfig, deviceId string) structs.Device {
	timestamp := time.Now().Format(time.RFC3339)
	warning := "no"

	if temperature := generateTemperature(); temperature > 30 {
		warning = "yes"
	}

	device := structs.Device{
		Id:          uuid.NewString(),
		DeviceId:    deviceId,
		CreatedAt:   timestamp,
		Temperature: generateTemperature(),
		Warning:     warning,
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"MetricId":    &types.AttributeValueMemberS{Value: device.Id},
			"DeviceId":    &types.AttributeValueMemberS{Value: device.DeviceId},
			"CreatedAt":   &types.AttributeValueMemberS{Value: device.CreatedAt},
			"Temperature": &types.AttributeValueMemberN{Value: strconv.Itoa(device.Temperature)},
			"Warning":     &types.AttributeValueMemberS{Value: device.Warning},
		},
	}

	_, err := cfg.DB.PutItem(context.TODO(), input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	return device
}

func replace(date string) string {
	return strings.ReplaceAll(date, "/", "-")
}

func generateTemperature() int {
	return rand.Intn(c.MaxTemperatureValue-c.MinTemperatureValue+1) + c.MinTemperatureValue
}

func concatStartDate(date string) (string, string) {
	return date + c.HourStartDay, date + c.HourEndDay
}
