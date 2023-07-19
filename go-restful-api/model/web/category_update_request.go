
package web

type CategoryUpdateRequest struct{
	Id		int		`validate:"required"`
	Name	string	`json:"name" validate:"required,max=10,min=1"`
}