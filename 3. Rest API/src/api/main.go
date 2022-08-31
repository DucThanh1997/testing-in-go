package main

import (
	"backend/providers/locations_provider"
	"fmt"
)

func main() {
	country, err := locations_provider.GetCountry("AR")
	fmt.Println("err: ", err)
	fmt.Println("country: ", country)
}