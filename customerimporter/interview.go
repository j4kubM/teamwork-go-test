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

// DomainCount is the struct that stores domain name and the amount of emails that belong to it
type DomainCount struct {
	Domain     string
	EmailCount int
}

// DomainEmailsCounter returns sorted array of domain and count of emails belonging to it, that are read from the given file
func DomainEmailsCounter(filename string) ([]DomainCount, error) {
	if filename == "" {
		return nil, fmt.Errorf("file name is empty")
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open csv file: %s", err)
	}

	defer file.Close()

	domains, domainEmailsCount, err := readDomainsAndCountEmails(file)

	return sortDomainCount(domains, domainEmailsCount), nil
}

// readDomainsAndCountEmails reads given csv file and returns an array of domains, and map with the count of emails existing for each domain
func readDomainsAndCountEmails(file *os.File) ([]string, map[string]int, error) {
	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read csv header: %s", err)
	}

	var domains []string
	domainEmailsCount := make(map[string]int)

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

		domainParts, err := extractDomainParts(email)
		if err != nil {
			fmt.Printf("failed to extract email domain: %s\n", err)
			continue
		}
		// only count valid domains
		if !isDomainValid(domainParts) {
			continue
		}

		// only append domains that are new
		if domainEmailsCount[domainParts[0]] == 0 {
			domains = append(domains, domainParts[0])
		}

		domainEmailsCount[domainParts[0]] += 1
	}
	return domains, domainEmailsCount, nil
}

// extractDomainParts extracts domain parts from the given email, and returns it. In case of the invalid email format it returns an error
func extractDomainParts(email string) ([]string, error) {
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

// isDomainValid checks if the elements of array with the period-separated domain parts are valid
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

// sortDomainCount returns sorted array of DomainCount structs based on the array of domains and the map of domainEmailCount
func sortDomainCount(domains []string, domainEmailsCount map[string]int) []DomainCount {
	if len(domains) == 0 || domainEmailsCount == nil {
		return nil
	}
	sort.Strings(domains)
	res := []DomainCount{}
	for _, d := range domains {
		res = append(res, DomainCount{
			Domain:     d,
			EmailCount: domainEmailsCount[d],
		})
	}
	return res
}
