package types

import "gorm.io/gorm"

type DeletedAt struct {
	gorm.DeletedAt
}