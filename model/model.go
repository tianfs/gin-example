package model

import "time"

type Model struct {
    Id        int `gorm:"primary_key"`//primary_key:设置主键
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `gorm:"index:idx_del_at"`
}