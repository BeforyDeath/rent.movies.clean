package interfaces

import "github.com/BeforyDeath/rent.movies.clean/domain"

func NewPagination(p map[string]interface{}) (*domain.Pagination, error) {
	pages := new(domain.Pagination)
	err := NewValidator(p, pages)

	if pages.Limit < 1 {
		pages.Limit = 100 // fixme
	}
	if pages.Page > 0 {
		pages.Offset = pages.Limit * (pages.Page - 1)
	}
	return pages, err
}
