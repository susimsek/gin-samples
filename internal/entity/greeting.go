package entity

// Greeting represents a greeting entity in the database
type Greeting struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Message string `gorm:"type:text;not null" json:"message"`
}
