package domain

// UserRoleMapping represents the many-to-many relationship between users and roles
type UserRoleMapping struct {
	UserID         string `gorm:"type:text;not null;primaryKey;column:user_id"`
	RoleID         string `gorm:"type:text;not null;primaryKey;column:role_id"`
	User           User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"` // Relation to User
	Role           Role   `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE;"`
	AuditingEntity        // Embedded AuditingEntity for auditing fields
}

// TableName specifies the table name for UserRoleMapping
func (UserRoleMapping) TableName() string {
	return "user_role_mapping"
}
