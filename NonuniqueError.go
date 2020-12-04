package gobrief

import "fmt"

type NonuniqueError struct {
	keyName string
}

func (err NonuniqueError) Error() string {
	return fmt.Sprintf("Key %s is not unique", err.keyName)
}
