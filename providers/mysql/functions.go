package mysql

import (
	"go_sampler/filters"

	"gorm.io/gorm/schema"
)

func Walk[T schema.Tabler](callback func(obj T) error) error {
	var model *T
	queryset := DB.Model(&model).Order("id asc")
	var pos = 0
	var step = 100
	for {
		var objs []T
		err := queryset.Offset(pos).Limit(step).Find(&objs).Error
		if err != nil {
			return err
		}

		if len(objs) == 0 {
			break
		}

		pos += step
		for _, obj := range objs {
			err = callback(obj)
			if err != nil {
				continue
			}
		}
	}

	return nil
}

func List[T schema.Tabler](pagination *Pagination, filter filters.Filter) ([]T, error) {
	var model T
	queryset := DB.Model(model)
	filter.Filter(queryset)

	if err := pagination.CountTotal(queryset); err != nil {
		return nil, err
	}

	var objs = make([]T, 0)
	err := queryset.Scopes(pagination.Scopes()).Find(&objs).Error
	return objs, err
}

func ListAll[T schema.Tabler](where, orderBy string, args ...any) ([]T, error) {
	var model T
	var objs = make([]T, 0)
	query := DB.Model(model)
	if where != "" {
		query.Where(where, args...)
	}

	if orderBy != "" {
		query.Order(orderBy)
	}

	err := query.Find(&objs).Error
	return objs, err
}

func First[T schema.Tabler](where string, args ...any) (obj T, err error) {
	var p = []any{where}
	p = append(p, args...)
	err = DB.Model(&obj).First(&obj, p...).Error
	return obj, err
}

func Exist[T schema.Tabler](where string, args ...any) bool {
	_, err := First[T](where, args...)
	return err == nil
}

func Count[T schema.Tabler](where string, args ...any) (cnt int64) {
	var obj T
	DB.Model(&obj).Where(where, args...).Count(&cnt)
	return cnt
}

func Delete[T schema.Tabler](where string, args ...any) error {
	var model T
	var p = []any{where}
	p = append(p, args...)
	return DB.Delete(&model, p...).Error
}

func Save[T schema.Tabler](obj T) error {
	return DB.Save(obj).Error
}

func Create[T schema.Tabler](obj T) error {
	return DB.Create(obj).Error
}
