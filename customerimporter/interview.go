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
	"sort"
	"strings"
)

// 1. read the csv file, check if it's there and if data is not missing, otherwise -> throw an error
// 2. create a method that takes the emails and extracts just the domain (read only symbols after @ and finish if the next one is .) append domain names as the keys of the map
// 3. create the comparing function that checks if the email contains a domain name, ??? Or maybe just when extracting domain name, append it to the map as the key, and increase the count by 1 if only the emails matter not the customer numbers.

func DomainCustomersCounter(filename string) error {
	if filename == "" {
		return fmt.Errorf("file name is empty")
	}
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open csv file: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read csv header: %s", err)
	}

	var domains []string

	columnIndex := make(map[string]int)
	for i, colName := range header {
		columnIndex[colName] = i
	}

	for {
		// read a row
		row, err := reader.Read()
		// check for end of file
		if err != nil {
			break
		}

		email := row[columnIndex["email"]]

		domainParts, err := extractDomain(email)
		if err != nil {
			return fmt.Errorf("failed to extract email domain: %s", err)
		}
		// only append valid domains
		if isDomainValid(domainParts) {
			domains = append(domains, domainParts[0])
		}
	}

	sort.Strings(domains)

	return nil
}

// extractDomain extracts domain from the given email, and returns it. In case of the invalid email or domain format it returns an error
func extractDomain(email string) ([]string, error) {
	emailParts := strings.SplitAfter(email, "@")
	if len(emailParts) == 1 {
		return []string{}, fmt.Errorf(`wrong email format, missing "@" in: %s`, email)
	}

	emailParts = strings.Split(emailParts[1], ".")
	if len(emailParts) == 1 {
		return []string{}, fmt.Errorf(`wrong email format, missing "." in: %s`, email)
	}

	return emailParts, nil
}

// isDomainValid checks if the elements of array with the period-separated domain parts is valid
func isDomainValid(domainParts []string) bool {
	// last part of the email must consist of two period-separated parts
	if len(domainParts) != 2 {
		return false
	}
	// last portion of the domain must be at least two characters
	if len(domainParts[1]) <= 1 {
		return false
	}
	return true
}
