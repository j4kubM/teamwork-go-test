// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// 1. read the csv file, check if it's there and if data is not missing, otherwise -> throw an error
// 2. create a method that takes the emails and extracts just the domain (read only symbols after @ and finish if the next one is .) append domain names as the keys of the map
// 3. create the comparing function that checks if the email contains a domain name, ??? Or maybe just when extracting domain name, append it to the map as the key, and increase the count by 1 if only the emails matter not the customer numbers.

func ReadCsvFile() {
	file, err := os.Open("../customers.csv")
	if err != nil {
		fmt.Println("failed to open csv file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		fmt.Println("failed to read csv header:", err)
		return
	}

	// var customerEmails []string

	columnIndex := make(map[string]int)
	for i, colName := range header {
		columnIndex[colName] = i
	}

	for {
		// Read a row
		row, err := reader.Read()
		// Check for end of file
		if err != nil {
			break
		}

		email := row[columnIndex["email"]]

		// customerEmails = append(customerEmails, email)
		domain, err := extractDomain(email)
		if err != nil {
			fmt.Println("failed to extract email domain:", err)
		}
		fmt.Println(domain)
	}
}

func extractDomain(email string) (string, error) {
	emailTail, found := strings.CutPrefix(email, "@")
	if !found {
		return "", fmt.Errorf("wrong email format, missing @ in: %s", email)
	}
	domain, found := strings.CutSuffix(emailTail, ".")
	if !found {
		return "", fmt.Errorf(`failed to find domain, wrong email format, missing ".": %s`, email)
	}
	return domain, nil
}
