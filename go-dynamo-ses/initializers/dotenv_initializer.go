package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	FROM    string
	TO      []string
	SUBJECT string
	MESSAGE string
)

func LoadDotEnv() {
	err := godotenv.Load("/home/gabriel/Documentos/go-aws/go-dynamo/.env")
	if err != nil {
		panic(err)
	}

	log.Println(os.Getenv("EMAIL_FROM"))
	FROM = os.Getenv("EMAIL_FROM")
	TO = []string{os.Getenv("EMAIL_TO")}
	SUBJECT = os.Getenv("EMAIL_SUBJECT")
	MESSAGE = os.Getenv("EMAIL_CONTENT")
}
