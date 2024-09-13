package responses

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string
	Timestamp time.Time
}


func GetResponseWithJson(ctx *gin.Context, code int, payload interface{}) {
	ctx.JSON(code, gin.H{
		"data" : payload,
	})
}

func GetResponseUnmarshalError(ctx *gin.Context, code int, entity string) {
	GetResponseWithJson(ctx, code, ErrorResponse{
		Message: fmt.Sprintf("Error to unmarshal entity: %v\n", entity),
		Timestamp: time.Now().UTC(),
	})
}

func GetResponseMarshalError(ctx *gin.Context, code int, entity string) {
	GetResponseWithJson(ctx, code, ErrorResponse{
		Message: fmt.Sprintf("Error to marshal entity: %v\n", entity),
		Timestamp: time.Now().UTC(),
	})
}

func GetResponseQueryError(ctx *gin.Context, code int, message string) {
	GetResponseWithJson(
		ctx, 
		code,
		ErrorResponse{
			Message: message,
			Timestamp: time.Now().UTC(),
		},
	)
}