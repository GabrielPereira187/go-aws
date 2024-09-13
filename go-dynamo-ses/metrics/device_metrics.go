package metrics

import (
	"log"
	"sync"
	"time"

	"github.com/GabrielPereira187/go-dynamo/email"
	"github.com/GabrielPereira187/go-dynamo/handler"
	"github.com/GabrielPereira187/go-dynamo/utils"
)

func StartSendingMetrics(concurrency int, timeBetweenRequest time.Duration, cfg *utils.ApiConfig) {
	devices := make([]string, 0)
	devices = append(devices, "0001")
	devices = append(devices, "0002")
	devices = append(devices, "0003")

	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		wg := &sync.WaitGroup{}
		for _, device := range devices {
			wg.Add(1)
			time.Sleep(time.Second * 5)
			go sendDeviceMetric(device, wg, cfg)
		}

		wg.Wait()
	}

}

func sendDeviceMetric(deviceId string, wg *sync.WaitGroup, cfg *utils.ApiConfig) {
	defer wg.Done()

	deviceResponse := handler.InsertDevice(cfg, deviceId)
	if deviceResponse.Temperature > 32 {
		log.Println("Email enviado")
		email.SendEmail(cfg, deviceId, deviceResponse.CreatedAt)
	}

}
