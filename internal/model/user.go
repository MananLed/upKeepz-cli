package model

type UserRole string

const(
	RoleAdmin UserRole = "admin"
	RoleOfficer UserRole = "officer"
	RoleResident UserRole = "resident"
)

type User struct{
	firstName string
	middleName string 
	lastName string 
	mobileNumber string
	email string 
	id string  
	password string
	role UserRole
}