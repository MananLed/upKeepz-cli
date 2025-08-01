package model

type Status string

const(
	StatusPending Status = "pending"
	StatusApproved Status = "approved"
	StatusInProgress Status = "inProgress"
	StatusCompleted Status = "completed"
	StatusCancelled Status = "cancelled"
)

type ServiceType string

const(
	Electrician ServiceType = "electrician"
	Plumber ServiceType = "plumber"
)

type ServiceRequest struct {
	requestID   string
	residentID  string
	status      Status
	timeSlot    string
	serviceType ServiceType
}
