package dto

type CreateRoleRequestDTO struct {
	Name        string `json:"name" validate:"required,min=2,max=50"`
	Description string `json:"description" validate:"required,min=5,max=200"`
}

type UpdateRoleRequestDTO struct {
	Name        string `json:"name" validate:"required,min=2,max=50"`
	Description string `json:"description" validate:"required,min=5,max=200"`
}

type AssignRoleRequestDTO struct {
	RoleID int64 `json:"role_id" validate:"required"`
}

type RemoveRoleRequestDTO struct {
	RoleID int64 `json:"role_id" validate:"required"`
}
