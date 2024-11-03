package pagination

import "math"

type PageInfo struct {
	Limit       int `json:"limit"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
}

func CalculateData(totalCount, limitInt, pageInt int) PageInfo {
	lastPage := int(math.Ceil(float64(totalCount) / float64(limitInt)))

	paginationInfo := PageInfo{
		Limit:       limitInt,
		CurrentPage: pageInt,
		LastPage:    lastPage,
	}
	return paginationInfo
}
