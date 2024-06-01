package err

import "fmt"

// NotFound описывает ошибку, возникающую, когда элемент не найден.
type NotFound struct {
	ID string
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("Resource with ID %s not found", e.ID)
}
