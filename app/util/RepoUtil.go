package util

import (
	"idstar-idp/rest-api/app/dto/response/rsdata"
	"math"

	"gorm.io/gorm"
)

func CheckRecordExists(anyModel interface{}, id uint, db *gorm.DB) bool {
	var count int64
	db.Model(anyModel).Where("id = ?", id).Count(&count)
	return count > 0
}

func CountRowsAndPages(anyModel interface{}, pageable *rsdata.PaginationData, db *gorm.DB) {
	var totalRows int64
	db.Model(anyModel).Count(&totalRows)
	pageable.TotalElements = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pageable.GetLimit())))
	pageable.TotalPages = totalPages
}

func Paginate(pageable *rsdata.PaginationData, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pageable.GetOffset()).Limit(pageable.GetLimit()).Order(pageable.GetSort())
	}
}
