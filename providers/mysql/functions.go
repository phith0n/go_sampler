package mysql

import (
	"go_sampler/filters"
	"gorm.io/gorm/schema"
)

func Walk[T schema.Tabler](callback func(vuln T) error) error {
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

	var objs []T
	err := queryset.Scopes(pagination.Scopes()).Find(&objs).Error
	return objs, err
}

func First[T schema.Tabler](where string, args ...any) (obj T, err error) {
	var p = []any{where}
	p = append(p, args...)
	err = DB.Model(&obj).First(&obj, p...).Error
	return obj, err
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
