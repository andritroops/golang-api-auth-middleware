package request

type CategoryCreateRequest struct {
	Name string   `json:"name" form:"name" validate:"required"`
	File []string `json:"files" form:"files"`
}
