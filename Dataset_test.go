package gobrief

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

var addTestTable = []struct {
	name          string
	key           string
	existingKeys  []string
	expectedKeys  []string
	expectedError error
}{
	{"key not exists", "key", []string{}, []string{"key"}, nil},
	{"another key exists", "key", []string{"another key"}, []string{"key", "another key"}, nil},
	{"key already exists", "key", []string{"key"}, []string{"key"}, NonuniqueError{keyName: "key"}},
}

var getTestTable = []struct {
	name          string
	key           string
	existingKeys  map[string]string
	expectedValue interface{}
}{
	{"key not exists", "key", map[string]string{}, nil},
	{"key exists", "key", map[string]string{"key": "value"}, "value"},
}

func TestDataset(test *testing.T) {
	suite.Run(test, new(DatasetTestSuite))
}

type DatasetTestSuite struct {
	suite.Suite
}

func (suite *DatasetTestSuite) Test_Add_givenKeyAndDataset_keyMustBeAddedIfUnique() {
	for _, testCase := range addTestTable {
		value := ""
		datum := Dataset{}.Create()
		for _, existingKey := range testCase.existingKeys {
			datum.Add(existingKey, value)
		}

		err := datum.Add(testCase.key, value)

		for _, expectedKey := range testCase.expectedKeys {
			assert.True(suite.T(), datum.Exists(expectedKey), err, fmt.Sprintf("Test case: \"%s\"", testCase.name))
		}
		assert.Equal(suite.T(), testCase.expectedError, err, fmt.Sprintf("Test case: \"%s\"", testCase.name))
	}
}

func (suite *DatasetTestSuite) Test_Set_givenKeyAndValue_keyMustBeSetInDataset() {
	key := "key"
	keyExisting := "key existing"
	value := "value"
	valueExisting := "value existing"
	datum := Dataset{}.Create()
	datum.Set(keyExisting, valueExisting)

	datum.Set(key, value)
	datum.Set(keyExisting, value)

	assert.Equal(suite.T(), value, datum.Get(key))
	assert.Equal(suite.T(), value, datum.Get(keyExisting))
}

func (suite *DatasetTestSuite) Test_Get_givenDataset_mustReturnIfKeyExists() {
	for _, testCase := range getTestTable {
		datum := Dataset{}.Create()
		for existingKey, existingValue := range testCase.existingKeys {
			datum.Add(existingKey, existingValue)
		}

		result := datum.Get(testCase.key)

		assert.Equal(suite.T(), testCase.expectedValue, result, fmt.Sprintf("Test case: \"%s\"", testCase.name))
	}
}
