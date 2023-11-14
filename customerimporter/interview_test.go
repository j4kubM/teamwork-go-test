package customerimporter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainEmailsCounter(t *testing.T) {
	for _, tc := range []struct {
		name     string
		filename string
		exp      []DomainCount
		expErr   error
	}{
		{
			name:     "missing filename",
			filename: "",
			exp:      nil,
			expErr:   fmt.Errorf("file name is empty"),
		},
		{
			name:     "successful file read",
			filename: "../testCustomers.csv",
			exp: []DomainCount{
				{
					Domain:     "360",
					EmailCount: 1,
				},
				{
					Domain:     "alternative",
					EmailCount: 1,
				},
				{
					Domain:     "github",
					EmailCount: 2,
				},
				{
					Domain:     "statcounter",
					EmailCount: 1,
				},
			},
		},
	} {
		actual, err := DomainEmailsCounter(tc.filename)
		if tc.expErr != nil {
			assert.Equal(t, tc.expErr, err)
		} else {
			assert.Equal(t, tc.exp, actual)
			assert.NoError(t, err)
		}
	}
}

func TestExtractDomain(t *testing.T) {
	for _, tc := range []struct {
		name   string
		email  string
		exp    []string
		expErr error
	}{
		{
			name:   "email missing @",
			email:  "notAValidEmail",
			expErr: fmt.Errorf(`wrong email format, missing "@" in: notAValidEmail`),
		},
		{
			name:   "email missing .",
			email:  "john-smith@invalidDomain",
			expErr: fmt.Errorf(`wrong email format, missing "." in: john-smith@invalidDomain`),
		},
		{
			name:  "successful domain extraction",
			email: "john-smith@validmail.com",
			exp:   []string{"validmail", "com"},
		},
	} {
		actual, err := extractDomainParts(tc.email)
		if tc.expErr != nil {
			assert.Equal(t, tc.expErr, err)
		} else {
			assert.Equal(t, tc.exp, actual)
			assert.NoError(t, err)
		}
	}
}

func TestIsDomainValid(t *testing.T) {
	for _, tc := range []struct {
		name        string
		domainParts []string
		exp         bool
	}{
		{
			name:        "not enough domain parts",
			domainParts: []string{"validmail"},
			exp:         false,
		},
		{
			name:        "too many domain parts",
			domainParts: []string{"validmail", "but-too-complicated", "com"},
			exp:         false,
		},
		{
			name:        "invalid last part",
			domainParts: []string{"validmail", "c"},
			exp:         false,
		},
		{
			name:        "valid domain",
			domainParts: []string{"validmail", "com"},
			exp:         true,
		},
	} {
		actual := isDomainValid(tc.domainParts)
		assert.Equal(t, tc.exp, actual)
	}

}

func TestSortDomainCount(t *testing.T) {
	for _, tc := range []struct {
		name              string
		domains           []string
		domainEmailsCount map[string]int
		exp               []DomainCount
	}{
		{
			name: "empty array of domains",
			domainEmailsCount: map[string]int{
				"validDomain": 1,
				"something":   2,
				"mymail":      3,
			},
			exp: nil,
		},
		{
			name: "empty map of domain emails count",
			domains: []string{
				"validDomain",
				"something",
				"mymail",
			},
			exp: nil,
		},
		{
			name: "empty array of domains",
			domainEmailsCount: map[string]int{
				"validDomain": 1,
				"something":   2,
				"mymail":      3,
			},
			domains: []string{
				"validDomain",
				"something",
				"mymail",
			},
			exp: []DomainCount{
				{
					Domain:     "mymail",
					EmailCount: 3,
				},
				{
					Domain:     "something",
					EmailCount: 2,
				},
				{
					Domain:     "validDomain",
					EmailCount: 1,
				},
			},
		},
	} {
		actual := sortDomainCount(tc.domains, tc.domainEmailsCount)
		assert.Equal(t, tc.exp, actual)
	}

}
