package dto

type Role struct{
	Id int
	Name string
	Description string
	CreatedAt string
	UpdatedAt string
}
type Permission struct{
	Id int 
	Name string
	Description string
	Resource string
	Action string
	CreatedAt string
	UpdatedAt string
}
