package db

import (
	"log/slog"
	"time"

	"go_sampler/filters"

	"gorm.io/gorm/schema"
)

type BaseModel struct {
	ID             uint      `json:"id" gorm:"column:id;not null;primaryKey;autoIncrement;"`
	CreatedTime    time.Time `json:"created_time" gorm:"column:created_time;autoCreateTime;"`
	LastModifyTime time.Time `json:"last_modify_time" gorm:"column:last_modify_time;autoUpdateTime;"`
}

type QueryWhere struct {
	where string
	args  []any
}

type QueryChain[T schema.Tabler] struct {
	wheres     []*QueryWhere
	orders     []string
	pagination *Pagination
	filter     filters.Filter
	limit      int
	offset     int
}

func Model[T schema.Tabler](_ T) *QueryChain[T] {
	return &QueryChain[T]{}
}

func (c *QueryChain[T]) Where(where string, args ...any) *QueryChain[T] {
	c.wheres = append(c.wheres, &QueryWhere{where: where, args: args})
	return c
}

func (c *QueryChain[T]) Order(orderBy string) *QueryChain[T] {
	c.orders = append(c.orders, orderBy)
	return c
}

func (c *QueryChain[T]) Limit(limit int) *QueryChain[T] {
	if c.pagination != nil {
		slog.Warn("limit and pagination are in conflict, pagination will take priority")
	}

	c.limit = limit
	return c
}

func (c *QueryChain[T]) Offset(offset int) *QueryChain[T] {
	if c.pagination != nil {
		slog.Warn("offset and pagination are in conflict, pagination will take priority")
	}

	c.offset = offset
	return c
}

func (c *QueryChain[T]) Pagination(pg *Pagination) *QueryChain[T] {
	if c.limit > 0 || c.offset > 0 {
		slog.Warn("offset, limit and pagination are in conflict, pagination will take priority")
	}

	c.pagination = pg
	return c
}

func (c *QueryChain[T]) Filter(f filters.Filter) *QueryChain[T] {
	c.filter = f
	return c
}

func (c *QueryChain[T]) First() (obj T, err error) {
	tx := DB.Model(&obj)
	for _, w := range c.wheres {
		tx.Where(w.where, w.args...)
	}

	err = tx.First(&obj).Error
	return obj, err
}

func (c *QueryChain[T]) Exist() bool {
	_, err := c.First()
	return err == nil
}

func (c *QueryChain[T]) Count() (cnt int64) {
	var obj T
	tx := DB.Model(&obj)
	for _, w := range c.wheres {
		tx.Where(w.where, w.args...)
	}

	tx.Count(&cnt)
	return cnt
}

func (c *QueryChain[T]) Delete() error {
	var obj T
	tx := DB.Model(&obj)
	for _, w := range c.wheres {
		tx.Where(w.where, w.args...)
	}

	return tx.Delete(&obj).Error
}

func (c *QueryChain[T]) List() ([]T, error) {
	var model T
	tx := DB.Model(model)

	for _, w := range c.wheres {
		tx.Where(w.where, w.args...)
	}

	for _, orderBy := range c.orders {
		tx.Order(orderBy)
	}

	if c.filter != nil {
		c.filter.Filter(tx)
	}

	if c.pagination != nil {
		if err := c.pagination.CountTotal(tx); err != nil {
			return nil, err
		}

		tx.Scopes(c.pagination.Scopes())
	} else {
		if c.offset > 0 {
			tx.Offset(c.offset)
		}

		if c.limit > 0 {
			tx.Limit(c.limit)
		}
	}

	var objs = make([]T, 0)
	err := tx.Find(&objs).Error
	return objs, err
}

func (c *QueryChain[T]) Walk(callback func(obj T) error) error {
	var model *T
	tx := DB.Model(&model)
	for _, w := range c.wheres {
		tx.Where(w.where, w.args...)
	}

	for _, orderBy := range c.orders {
		tx.Order(orderBy)
	}

	if c.filter != nil {
		c.filter.Filter(tx)
	}

	var pos = c.offset
	var step = 100
	var already = 0
	for {
		if c.limit > 0 && already >= c.limit {
			break
		}

		var objs []T
		err := tx.Offset(pos).Limit(step).Find(&objs).Error
		if err != nil {
			return err
		}

		if len(objs) == 0 {
			break
		}

		pos += step
		for _, obj := range objs {
			already++
			err = callback(obj)
			if err != nil {
				continue
			}
		}
	}

	return nil
}

func Save[T schema.Tabler](obj T) error {
	return DB.Save(obj).Error
}

func Create[T schema.Tabler](obj T) error {
	return DB.Create(obj).Error
}
