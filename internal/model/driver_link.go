package model

type DriverLinkStatus string

const (
	DriverLinkStatusPending  DriverLinkStatus = "PENDING"
	DriverLinkStatusAccepted DriverLinkStatus = "ACCEPTED"
	DriverLinkStatusRejected DriverLinkStatus = "REJECTED"
)
