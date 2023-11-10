package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"example.com/hello/utils"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	// message, _ := greetings.Hello("")
	// If an error was returned, print it to the console and
	// exit the program.
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println("Posting data to prometheus")
	// Create a new random number generator with the current time as the seed
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	counter := random.Intn(100)

	prometheusPushgatewayURL := os.Getenv("DAPR_PERF_METRICS_PROMETHEUS_PUSHGATEWAY_URL")
	if prometheusPushgatewayURL == "" {
		log.Println("DAPR_PERF_METRICS_PROMETHEUS_PUSHGATEWAY_URL is not set, skipping pushing perf test metrics to Prometheus Pushgateway")
		return
	} else {
		fmt.Println("DAPR_PERF_METRICS_PROMETHEUS_URL is set to " + prometheusPushgatewayURL)
	}

	for i := 0; i < counter; i++ {

		// Generate a random number between 1 and 20
		randomNumber := random.Intn(15000) + 1

		daprSIMetrics := utils.DaprMetrics{
			BaselineLatency:       float64(i + randomNumber + 5),
			DaprLatency:           float64(i + randomNumber + 5),
			AddedLatency:          float64(i + randomNumber + 5),
			SidecarCPU:            int64(i + randomNumber + 5),
			AppCPU:                int64(i + randomNumber + 5),
			SidecarMemory:         float64(i + randomNumber + 5),
			AppMemory:             float64(i + randomNumber + 5),
			ApplicationThroughput: float64(i + randomNumber + 5),
		}
		utils.PushPrometheusMetrics(daprSIMetrics, "service-invocation-http", "")

		randomNumber = random.Intn(6000) + 1

		daprStateMetrics := utils.DaprMetrics{
			BaselineLatency:       float64(i + randomNumber + 5),
			DaprLatency:           float64(i + randomNumber + 5),
			AddedLatency:          float64(i + randomNumber + 5),
			SidecarCPU:            int64(i + randomNumber + 5),
			AppCPU:                int64(i + randomNumber + 5),
			SidecarMemory:         float64(i + randomNumber + 5),
			AppMemory:             float64(i + randomNumber + 5),
			ApplicationThroughput: float64(i + randomNumber + 5),
		}

		utils.PushPrometheusMetrics(daprStateMetrics, "state_get_http", "inmemory")
	}
}
