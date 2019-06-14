package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Defines a table with the params used in the tests.
// Then initializes as values as tests you want to execute.
// Finally iterates over the table, calls the function as many
// iterations as values you have defined and asserts the expected
// results.
func TestGetCompanyData(t *testing.T) {

	tests := map[string]struct {
		input         string
		response      string
		expectedError bool
	}{
		"get nickname": {
			input:         "rayo",
			response:      "walter",
			expectedError: false,
		},
		"get name": {
			input:         "unknown",
			response:      "",
			expectedError: true,
		},
	}

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		response, err := getCatNickname(test.input)
		assert.Equal(t, err != nil, test.expectedError)
		assert.Equal(t, response, test.response)
	}
}

func getCatNickname(name string) (string, error) {
	if name == "rayo" {
		return "walter", nil
	}
	return "", errors.New("Fail")

}
