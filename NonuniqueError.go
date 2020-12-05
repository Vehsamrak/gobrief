package gobrief

import "fmt"

// NonuniqueError defines situation when key is not unique
type NonuniqueError struct {
	keyName string
}

func (err NonuniqueError) Error() string {
	return fmt.Sprintf("Key %s is not unique", err.keyName)
}
