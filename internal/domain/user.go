package domain

// User represents a user in the system
type User struct {
	ID             string            `gorm:"primaryKey;type:text;column:id"`            // Unique identifier
	Username       string            `gorm:"type:text;not null;unique;column:username"` // Username
	Password       string            `gorm:"type:text;not null;column:password"`        // Password (hashed)
	Email          string            `gorm:"type:text;not null;unique;column:email"`    // Email
	FirstName      string            `gorm:"type:text;column:first_name"`               // First name
	LastName       string            `gorm:"type:text;column:last_name"`                // Last name
	Enabled        bool              `gorm:"type:boolean;not null;column:enabled"`      // Is the user enabled?
	Roles          []UserRoleMapping `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	AuditingEntity                   // Embedded AuditingEntity for auditing fields
}

// TableName specifies the table name for User
func (User) TableName() string {
	return "user_identity"
}

func (u User) GetID() interface{} {
	return u.ID
}
