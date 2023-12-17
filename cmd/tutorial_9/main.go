package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAX_CHICKEN_PRICE float32 = 5
	MAX_TOFU_PRICE    float32 = 3
)

func main() {
	chickenChannel := make(chan string)
	tofuChannel := make(chan string)
	websites := []string{"walmart.com", "cosco.com", "wholefood.com"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}

	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		chickenPrice := rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		tofuPrice := rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nText sent: Found a deal on chicken at %s", website)
	case website := <-tofuChannel:
		fmt.Printf("\nEmail sent: Found a deal on tofu at %s", website)
	}
}
