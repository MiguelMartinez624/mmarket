package errors

import "fmt"

// MissingField error when a object dont have all required fields
type MissingField struct {
	Field string
}

func (e MissingField) Error() string {
	return fmt.Sprintf("Missing object field [%v] ", e.Field)
}
