package domain

// Role represents a user role in the system
type Role struct {
	ID             string `gorm:"primaryKey;type:text;column:id"`        // Unique identifier
	Name           string `gorm:"type:text;not null;unique;column:name"` // Role name
	Description    string `gorm:"type:text;column:description"`          // Role description
	AuditingEntity        // Embedded AuditingEntity for auditing fields
}

// TableName specifies the table name for Role
func (Role) TableName() string {
	return "role"
}
