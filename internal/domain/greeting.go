package domain

// Greeting represents a greeting domain in the database
type Greeting struct {
	ID             uint   `gorm:"primaryKey;autoIncrement;column:id"` // Primary key
	Message        string `gorm:"type:text;not null;column:message"`  // Message column
	AuditingEntity        // Embedded AuditingEntity for auditing fields
}

func (Greeting) TableName() string {
	return "greeting"
}

func (g Greeting) GetID() interface{} {
	return g.ID
}
