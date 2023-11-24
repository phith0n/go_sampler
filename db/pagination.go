package db

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Total       int         `json:"total"`
	TotalPage   int         `json:"total_page"`
	PageSize    int         `json:"page_size"`
	CurrentPage int         `json:"current_page"`
	Data        interface{} `json:"data"`
}

func (p *Pagination) CountTotal(db *gorm.DB) error {
	var count int64
	if err := db.Count(&count).Error; err != nil {
		return err
	}
	p.Total = int(count)
	p.TotalPage = (p.Total-1)/p.PageSize + 1
	return nil
}

func (p *Pagination) Scopes() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (p.CurrentPage - 1) * p.PageSize
		return db.Offset(offset).Limit(p.PageSize)
	}
}

func NewPagination(c *gin.Context, defaultPageSize int, maxPageSize int) *Pagination {
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.Query("size"))
	switch {
	case pageSize > maxPageSize:
		pageSize = maxPageSize
	case pageSize <= 0:
		pageSize = defaultPageSize
	}

	return &Pagination{
		PageSize:    pageSize,
		CurrentPage: page,
	}
}
