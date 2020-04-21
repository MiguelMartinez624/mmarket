package profiles

import "fmt"

type MissingAccountIDError struct{}

func (e MissingAccountIDError) Error() string {
	return fmt.Sprintf("missing accoundID. Its a necesary field of the profile.")
}

type MissingFirstNameError struct {
}

func (e MissingFirstNameError) Error() string {
	return fmt.Sprintf("missing lastname. Its a necesary field of the profile.")
}

type ProfileDontFoundError struct {
}

func (e ProfileDontFoundError) Error() string {
	return fmt.Sprintf("profile dont found.")
}

type MissingParamError struct {
	Param string
}

func (e MissingParamError) Error() string {
	return fmt.Sprintf("missing [%v] paramerter", e.Param)
}
