package documentModel

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	UserRefer uint
}
