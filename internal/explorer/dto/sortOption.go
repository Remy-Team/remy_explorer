package dto

// this abstraction of the sort option is used to define the order in which the results should be sorted
// it helps to avoid using string literals in the SQL query and to avoid SQL injection
import "fmt"

// SortOrder represents the order in which the results should be sorted.
type SortOrder string

const (
	Ascending  SortOrder = "ASC"
	Descending SortOrder = "DESC"
)

// SortOption represents the sorting options for the results.
type SortOption struct {
	Field string
	Order SortOrder
}

// NewSortOption creates a new SortOption instance by validating the field and order.
func NewSortOption(field, order string) (*SortOption, error) {
	validFields := map[string]bool{
		"id": true, "name": true, "size": true, "created_at": true, "updated_at": true,
	}
	if _, ok := validFields[field]; !ok {
		return nil, fmt.Errorf("invalid sort field")
	}

	if order != "ASC" && order != "DESC" {
		return nil, fmt.Errorf("invalid sort order")
	}

	return &SortOption{Field: field, Order: SortOrder(order)}, nil
}
