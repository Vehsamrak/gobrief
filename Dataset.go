package gobrief

import "strings"

type Dataset struct {
	dataset map[string]interface{}
}

// Create new dataset
func (dataset Dataset) Create() *Dataset {
	return &Dataset{dataset: map[string]interface{}{}}
}

// Add key/value pair to dataset. Key must be unique.
// If it is not unique, error would be returned
func (dataset *Dataset) Add(key string, value string) error {
	if dataset.Exists(key) {
		return NonuniqueError{keyName: key}
	}

	dataset.dataset[key] = value

	return nil
}

// Set value to key. If key exists it's value would be overridden.
// If key not exists, it would be created
func (dataset *Dataset) Set(key string, value string) {
	dataset.dataset[key] = value
}

// Exists checks if key exists in dataset
func (dataset *Dataset) Exists(key string) bool {
	_, exists := dataset.dataset[key]

	return exists
}

// Get value of key if key name exactly matched
// Will return nil if no matches were found
func (dataset *Dataset) Get(key string) interface{} {
	return dataset.dataset[key]
}

// GetStartedWith returns map of keys/values of all keys that have name started with provided argument
// Will return empty map if no matches were found
func (dataset *Dataset) GetStartedWith(key string) map[string]interface{} {
	resultMap := make(map[string]interface{})
	for dataKey, value := range dataset.dataset {
		if strings.HasPrefix(dataKey, key) {
			resultMap[dataKey] = value
		}
	}

	return resultMap
}

// First returns first match of key that have name started with provided argument
// Will return nil if no matches were found
func (dataset *Dataset) First(key string) interface{} {
	for dataKey, value := range dataset.dataset {
		if strings.HasPrefix(dataKey, key) {
			return value
		}
	}

	return nil
}
