package models

import (
	"time"
)

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time `json:"created_at"`           // 创建时间
	UpdatedAt time.Time `json:"-"`                    // 更新时间
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}

// swagger:model PageInfo
type PageInfo struct {
	// Page number for pagination
	// Required: true
	// Minimum: 1
	Page int `form:"page" json:"page"`

	// Search keyword
	Key string `form:"key" json:"key"`

	// Number of items per page
	// Required: true
	// Minimum: 1
	// Maximum: 100
	Limit int `form:"limit" json:"limit"`

	// Sort order (e.g., "asc" or "desc")
	Sort string `form:"sort" json:"sort"`
}
