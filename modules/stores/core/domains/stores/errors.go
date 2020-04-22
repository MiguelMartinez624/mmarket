package stores

import "fmt"

type MissinField struct {
	Field string
}

func (e MissinField) Error() string {
	return fmt.Sprintf("Missing field %v", e.Field)
}
