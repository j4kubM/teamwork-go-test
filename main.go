package main

import (
	"fmt"

	"github.com/j4kubM/teamwork-go-test/customerimporter"
)

func main() {
	file := "customers.csv"
	err := customerimporter.DomainCustomersCounter(file)
	if err != nil {
		fmt.Println(err)
	}
}
