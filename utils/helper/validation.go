package helper

func ValidateCountLimitAndPage(page, limit int) (int, int) {
	if page <= 0 {
		page = 1
	}

	maxLimit := 10
	if limit <= 0 || limit > maxLimit {
		limit = maxLimit
	}

	return page, limit
}
