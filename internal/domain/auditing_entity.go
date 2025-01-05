package domain

import "time"

// AuditingEntity provides common audit fields for entities
type AuditingEntity struct {
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"` // Custom column name
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
