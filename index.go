package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doInterval(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func measuring(t time.Time) {

	min := 1
	max := 15

	rand.Seed(time.Now().Unix())

	numbrandwater := rand.Intn(max - min)
	numbrandwind := rand.Intn(max - min)

	switch numbrandwater > 0 {
	case numbrandwater < 6:
		result := "aman"
		fmt.Printf("status water: %s \n", result)
	case numbrandwater > 5 && numbrandwater < 9:
		result := "siaga"
		fmt.Printf("status water: %s \n", result)
	case numbrandwater > 8:
		result := "bahaya"
		fmt.Printf("status water: %s \n", result)
	default:
		result := "water measurable error"
		fmt.Println(result)
	}

	switch numbrandwind > 0 {
	case numbrandwind < 6:
		result := "aman"
		fmt.Printf("status wind: %s \n", result)
	case numbrandwind > 5 && numbrandwind < 9:
		result := "siaga"
		fmt.Printf("status wind: %s \n", result)
	case numbrandwind > 8:
		result := "bahaya"
		fmt.Printf("status wind: %s \n", result)
	default:
		result := "wind measurable error"
		fmt.Println(result)
	}

	fmt.Println("Random value Water : ", numbrandwater)
	fmt.Println("Random value Wind : ", numbrandwind)
}

func main() {
	doInterval(5000*time.Millisecond, measuring)
}
