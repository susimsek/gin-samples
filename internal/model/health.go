package model

// HealthStatus represents the health status of the application
// @Description Health status model
type HealthStatus struct {
	// Status indicates the health of the application
	Status string `json:"status" example:"UP" validate:"required"`
}
