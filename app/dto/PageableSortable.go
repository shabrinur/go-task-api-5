package dto

type Pageable struct {
	Page   int `json:"page" form:"page" validate:"numeric"`
	Size   int `json:"size" form:"size" validate:"numeric"`
	Offset int `json:"offset"`
}

type Sortable struct {
	Field     string `json:"field" form:"field"`
	Direction string `json:"direction" form:"direction"`
}
