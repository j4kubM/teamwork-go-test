# teamwork-go-test
teamwork go test is the program that uses internal package customerimporter and it's method DomainEmailCounter to read information from the .csv file with the following columns:
* first_name,
* last_name,
* email,
* gender,
* ip_address

Based on the file, it extracts all of the valid email domains, and for each domain counts how many customers are registered under that domain. The output is a sorted array of the DomainCount objects, that have following fields:
* Domain
* EmailCount

Valid domain in this project is defined as follows:
* It's last part is is at least 2 characters long (valid: example.co, invalid: example.d)
* It is does not include "." (valid: example.com, invalid:example.two.com)

## Running the code
In order to run the program, type `go run main.go` while being inside the teamwork-go-test folder