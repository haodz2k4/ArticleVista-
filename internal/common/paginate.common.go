package common

import "math"

type Pagination struct {
	CurrentPage int
	Limit       int
	Skip        int
	CountPage   int
}

func NewPagination(currentPage int, limit int, total int64) (p *Pagination) {

	skip := (currentPage - 1) * limit
	countPage := int(math.Ceil(float64(total) / float64(limit)))
	return &Pagination{CurrentPage: currentPage, Limit: limit, Skip: skip, CountPage: countPage}
}
