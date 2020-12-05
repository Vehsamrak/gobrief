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
	name           string
	key            string
	existingKeys   map[string]string
	expectedResult interface{}
}{
	{"key not exists", "key", map[string]string{}, nil},
	{"one key exists", "key", map[string]string{"key": "value"}, "value"},
	{"multiple keys started with different letters exists. Must return one value",
		"first key",
		map[string]string{
			"first key":  "first value",
			"second key": "second value",
		},
		"first value",
	},
	{"multiple keys started with same letters exists. Must return one value of key with exact name",
		"key",
		map[string]string{
			"key":        "first value",
			"key second": "second value",
			"third key":  "third value",
		},
		"first value",
	},
}

var getStartedWithTestTable = []struct {
	name           string
	key            string
	existingKeys   map[string]string
	expectedResult map[string]interface{}
}{
	{"key not exists", "key", map[string]string{}, map[string]interface{}{}},
	{"one key exists", "key", map[string]string{"key": "value"}, map[string]interface{}{"key": "value"}},
	{"multiple keys started with different letters exists. Must return one key/value",
		"first key",
		map[string]string{
			"first key":  "first value",
			"second key": "second value",
		},
		map[string]interface{}{"first key": "first value"},
	},
	{"multiple keys started with same letters exists. Must return two keys/values",
		"key",
		map[string]string{
			"key":        "first value",
			"key second": "second value",
			"third key":  "third value",
		},
		map[string]interface{}{
			"key":        "first value",
			"key second": "second value",
		},
	},
}

var getFirstTestTable = []struct {
	name           string
	key            string
	existingKeys   map[string]string
	expectedResult interface{}
}{
	{"key not exists", "key", map[string]string{}, nil},
	{"one key exists", "key", map[string]string{"key": "value"}, "value"},
	{"multiple keys started with different letters exists. Must return one value",
		"first key",
		map[string]string{
			"first key":  "first value",
			"second key": "second value",
		},
		"first value",
	},
	{"multiple keys started with same letters exists. Must return one value",
		"key",
		map[string]string{
			"key first":  "first value",
			"key second": "second value",
			"third key":  "third value",
		},
		"first value",
	},
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
		if testCase.expectedError != nil {
			assert.NotEmpty(suite.T(), err.Error(), fmt.Sprintf("Test case: \"%s\"", testCase.name))
		}
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

func (suite *DatasetTestSuite) Test_Get_givenDataset_mustReturnIfKeysExists() {
	for _, testCase := range getTestTable {
		datum := Dataset{}.Create()
		for existingKey, existingValue := range testCase.existingKeys {
			datum.Add(existingKey, existingValue)
		}

		result := datum.Get(testCase.key)

		assert.Equal(suite.T(), testCase.expectedResult, result, fmt.Sprintf("Test case: \"%s\"", testCase.name))
	}
}

func (suite *DatasetTestSuite) Test_GetStartedWith_givenDataset_mustReturnResultMap() {
	for _, testCase := range getStartedWithTestTable {
		datum := Dataset{}.Create()
		for existingKey, existingValue := range testCase.existingKeys {
			datum.Add(existingKey, existingValue)
		}

		result := datum.GetStartedWith(testCase.key)

		assert.Equal(suite.T(), testCase.expectedResult, result, fmt.Sprintf("Test case: \"%s\"", testCase.name))
	}
}

func (suite *DatasetTestSuite) Test_First_givenDataset_mustReturnResultMap() {
	for _, testCase := range getFirstTestTable {
		datum := Dataset{}.Create()
		for existingKey, existingValue := range testCase.existingKeys {
			datum.Add(existingKey, existingValue)
		}

		result := datum.First(testCase.key)

		assert.Equal(suite.T(), testCase.expectedResult, result, fmt.Sprintf("Test case: \"%s\"", testCase.name))
	}
}
