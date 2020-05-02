package core

import "fmt"

type ErrUnavailibleItems struct {
	Items []string `json:"items"`
}

func (e ErrUnavailibleItems) Error() string {
	return fmt.Sprintf("There are %v not longer availables", len(e.Items))
}
