
package web

type CategoryCreateRequest struct{
	Name	string	`json:"name" validate:"required,max=10,min=1"`
}