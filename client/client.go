package main

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	rate := vegeta.Rate{Freq: 100, Per: time.Second}
	duration := 60 * 60 * time.Second
	targeter := vegeta.NewStaticTargeter(
		vegeta.Target{
			Method: "GET",
			URL:    "http://localhost:3001/",
		},
		vegeta.Target{
			Method: "GET",
			URL:    "http://localhost:3002/",
		},
	)
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
	fmt.Println("errors: %v", metrics.Errors)
}
