package common

import "math"

type Pagination struct {
	currentPage int
	Limit       int
	Skip        int
	CountPage   int
}

func NewPagination(currentPage int, limit int, total int) (p *Pagination) {

	skip := (currentPage - 1) * limit
	countPage := int(math.Ceil(float64(total) / float64(limit)))
	return &Pagination{currentPage: currentPage, Limit: limit, Skip: skip, CountPage: countPage}
}
