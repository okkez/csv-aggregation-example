package main

import(
	"fmt"
	"math/rand"
	"strings"

	"github.com/bxcodec/faker/v3"
)

const lines = 3000000

var serviceNames []string


func randomCost() float32 {
	return rand.Float32() * 100
}

func init() {
	for i := 0; i < 20; i++ {
		serviceNames = append(serviceNames, faker.Username())
	}
}

func main() {
	NumberOfServiceNames := len(serviceNames)
	headerNames := []string{
		"id",
		"name",
		"description",
		"cost",
	}
	headerLine := strings.Join(headerNames, ",")
	fmt.Println(headerLine)
	initialID := rand.Uint32()
	for i := 0; i < lines; i++ {
		line := fmt.Sprintf("%d,%s,%s,%.3f", initialID + uint32(i), serviceNames[rand.Int() % NumberOfServiceNames], faker.Sentence(), randomCost())
		fmt.Println(line)
	}
}
