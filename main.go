package main

import (
	"fmt"

	"github.com/j4kubM/teamwork-go-test/customerimporter"
)

func main() {
	file := "customers.csv"
	domainEmails, err := customerimporter.DomainEmailsCounter(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(domainEmails)
}
