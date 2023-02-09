package util

import (
	"math"

	"gorm.io/gorm"
)

func FilterPaginate(sqlSess *gorm.DB, page, contentPerPage int64) *gorm.DB {
	return sqlSess.
		Offset(int((page - 1) * contentPerPage)).
		Limit(int(contentPerPage))
}

func CountMaxPage(totalData, contentPerPage int64) int64 {
	if totalData == 0 || contentPerPage == 0 {
		return 1
	}

	ceil := math.Ceil(float64(totalData) / float64(contentPerPage))
	return int64(ceil)
}
