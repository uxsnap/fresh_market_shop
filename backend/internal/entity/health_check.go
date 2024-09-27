package entity

import "time"

const (
	HealthStatusOK = "OK"
	HealthStatusNotAvailable = "NotAvailable"
)

type AuthServiceHealthCheck struct {
	Status    string
	GitTag    string
	GitBranch string
	UpTime    time.Time
	Message   string
}
