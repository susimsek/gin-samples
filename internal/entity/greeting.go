package entity

// Greeting represents a greeting entity in the database
type Greeting struct {
	ID             uint   `gorm:"primaryKey;autoIncrement;column:id"` // Primary key
	Message        string `gorm:"type:text;not null;column:message"`  // Message column
	AuditingEntity        // Embedded AuditingEntity for auditing fields
}
