package models

import (
	"github.com/goravel/framework/database/orm"
	"time"
)

type Category struct {
	orm.Model
	Id        int
	Name      string `form:"name" binding:"required"`
	CreatedAt time.Time
	orm.SoftDeletes
}

func (r *Category) TableName() string {
	return "category"
}
