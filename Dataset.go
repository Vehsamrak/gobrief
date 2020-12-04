package gobrief

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

// Get values of all keys that have name started with provided argument
func (dataset *Dataset) Get(key string) interface{} {
	return dataset.dataset[key]
}
