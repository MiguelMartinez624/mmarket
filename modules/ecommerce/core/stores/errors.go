package stores

import "fmt"

type MissinField struct {
	Field string
}

func (e MissinField) Error() string {
	return fmt.Sprintf("Missing field %v", e.Field)
}

type ErrStoreNotFound struct {
}

func (e ErrStoreNotFound) Error() string {
	return fmt.Sprintf("Store not found ")
}

type ErrMissingParam struct {
	Param string
}

func (e ErrMissingParam) Error() string {
	return fmt.Sprintf("missing parameter %v ", e.Param)
}
