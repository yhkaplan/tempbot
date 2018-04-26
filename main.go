package main

import (
	"fmt"

	"github.com/d2r2/go-dht"
)

func main() {
	sensorType := dht.DHT11

	temp, humid, _, err := dht.ReadDHTxxWithRetry(sensorType, 4, false, 5); if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Tempurature: %v, Humidity: %v", temp, humid)
}
