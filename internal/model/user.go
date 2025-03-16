package model

type UserStatus string

const (
	UserStatusPending     UserStatus = "PENDING"
	UserStatusActive      UserStatus = "ACTIVE"
	UserStatusSuspended   UserStatus = "SUSPENDED"
	UserStatusDeactivated UserStatus = "DEACTIVATED"
)

type UserRole string

const (
	UserRoleAdmin  UserRole = "ADMIN"
	UserRoleStaff  UserRole = "STAFF"
	UserRoleDriver UserRole = "DRIVER"
)
