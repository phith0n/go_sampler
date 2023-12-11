package filters

import "gorm.io/gorm"

type Filter interface {
	Filter(tx *gorm.DB)
}
