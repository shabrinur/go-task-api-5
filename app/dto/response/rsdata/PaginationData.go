package rsdata

import "idstar-idp/rest-api/app/dto"

type PaginationData struct {
	Content          interface{}  `json:"content"`
	TotalElements    int64        `json:"totalElements"`
	TotalPages       int          `json:"totalPages"`
	NumberOfElements int          `json:"numberOfElements"`
	Pageable         dto.Pageable `json:"pageable"`
	Sortable         dto.Sortable `json:"sortable"`
	FirstPage        bool         `json:"firstPage,omitempty"`
	LastPage         bool         `json:"lastPage,omitempty"`
	EmptyPage        bool         `json:"emptyPage,omitempty"`
}

func (c *PaginationData) GetLimit() int {
	if c.Pageable.Size == 0 {
		c.Pageable.Size = 10
	}
	return c.Pageable.Size
}

func (c *PaginationData) GetPage() int {
	if c.Pageable.Page == 0 {
		c.Pageable.Page = 1
	}
	return c.Pageable.Page
}

func (c *PaginationData) GetOffset() int {
	c.Pageable.Offset = (c.GetPage() - 1) * c.GetLimit()
	return c.Pageable.Offset
}

func (c *PaginationData) GetSort() string {
	if c.Sortable.Field == "" {
		c.Sortable.Field = "id"
	}
	if c.Sortable.Direction == "" {
		c.Sortable.Direction = "asc"
	}
	return c.Sortable.Field + " " + c.Sortable.Direction
}

func (c *PaginationData) SetValueBeforeReturn() {
	if c.NumberOfElements <= 0 {
		c.EmptyPage = true
	} else {
		if c.Pageable.Page == 1 {
			c.FirstPage = true
		}
		if c.Pageable.Page == (c.TotalPages) {
			c.LastPage = true
		}
	}
}
