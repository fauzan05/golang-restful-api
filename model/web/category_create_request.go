package web

type CategoryCreateRequest struct {
	Name string `validate:"required,max=200,min=3" json:"name"`
}