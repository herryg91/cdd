// Code generated by protoc-gen-cdd. DO NOT EDIT.
// source: province.proto

package entity

import (
	"time"
)

// Mysql Table: tbl_province
type Province struct {
	Id   int    `gorm:"primary_key;column:id"`
	Name string `gorm:"column:name"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}