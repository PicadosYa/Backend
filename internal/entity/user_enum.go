package entity

type UserRole string

const (
	RoleCliente UserRole = "client"
	RoleCancha  UserRole = "field"
	RoleAdmin   UserRole = "admin"
)
