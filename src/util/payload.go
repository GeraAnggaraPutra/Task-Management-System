package util

import (
	"fmt"

	"task-management-system/src/constant"
)

type PaginationPayload struct {
	Search      string `form:"search"`
	Sort        string `form:"sort"`
	Direction   string `form:"direction"`
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
	SetSearch   bool
	Offset      int
	Order       string
	SetPaginate bool
}

// Initialize pagination payload.
func (p *PaginationPayload) Init() {
	if p.Search != "" {
		p.SetSearch = true
		p.Search = fmt.Sprintf("%%%s%%", p.Search)
	}

	if p.Sort == "" || p.Direction == "" {
		p.Order = constant.DefaultOrder
	} else {
		p.Order = fmt.Sprintf("%s %s", p.Sort, p.Direction)
	}

	if !(p.Page <= 0 && p.Limit <= 0) {
		if p.Page <= 0 {
			p.Page = constant.DefaultPage
		}

		if p.Limit <= 0 {
			p.Limit = constant.DefaultLimit
		}

		p.Offset = (p.Page * p.Limit) - p.Limit
		p.SetPaginate = true
	}
}
