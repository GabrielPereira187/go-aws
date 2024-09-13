package email

import (
	"context"
	"fmt"
	"log"

	"github.com/GabrielPereira187/go-dynamo/initializers"
	"github.com/GabrielPereira187/go-dynamo/utils"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

func SendEmail(cfg *utils.ApiConfig, deviceId, timestamp string) {
	from := initializers.FROM
	recipients := initializers.TO
	subject := fmt.Sprintf(initializers.SUBJECT, deviceId)

	message := fmt.Sprintf(initializers.MESSAGE, deviceId, timestamp)
	input := &sesv2.SendEmailInput{
		FromEmailAddress: &from,
		Destination: &types.Destination{
			ToAddresses: recipients,
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Text: &types.Content{
						Data: &message,
					},
				},
				Subject: &types.Content{
					Data: &subject,
				},
			},
		},
	}
	res, err := cfg.SesClient.SendEmail(context.TODO(), input)

	if err != nil || res.MessageId == nil {
		fmt.Println(err)
		return
	}

	log.Println("E-mail enviado com sucesso")

}
