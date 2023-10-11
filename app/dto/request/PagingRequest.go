package request

import (
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/dto"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slices"
)

type PagingRequest struct {
	Pageable dto.Pageable
	Sortable dto.Sortable
}

func (c *PagingRequest) Validate(validFields []string) error {
	validate := validator.New()

	if err := validate.Struct(c); err != nil {
		return err
	}

	if c.Sortable.Field != "" {
		c.Sortable.Field = strings.TrimSpace(c.Sortable.Field)
		if !slices.Contains(validFields, strings.ToLower(c.Sortable.Field)) {
			return errors.New(fmt.Sprint("invalid sort field; valid field: ", strings.Join(validFields, ", ")))
		}
	}

	if c.Sortable.Direction != "" {
		directions := []string{"asc", "desc"}
		c.Sortable.Direction = strings.TrimSpace(c.Sortable.Direction)
		if !slices.Contains(directions, strings.ToLower(c.Sortable.Direction)) {
			return errors.New("invalid sort direction; use asc or desc")
		}
	}
	return nil
}
