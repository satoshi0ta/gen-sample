package entity

import (
	"time"
)

type Task struct {
	ID        int64     `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	AdminID   int64     `gorm:"column:admin_id;type:bigint;"`
	Name      string    `gorm:"column:name;type:varchar;size:255;"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;"`
}
