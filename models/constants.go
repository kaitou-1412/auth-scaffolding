package models

type RoleType string

const (
	USER    RoleType = "user"
	MANAGER RoleType = "manager"
	ADMIN   RoleType = "admin"
)
