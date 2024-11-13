package dashboard

type MonthlyTransaction struct {
	Month string `json:"month"`
	Total int    `json:"total"`
}

type UserRanking struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Total int    `json:"total"`
}

type Dashboard struct {
	TotalUsers          int                  `json:"total_users"`
	TotalTransactions   int                  `json:"total_transactions"`
	TotalProducts       int                  `json:"total_products"`
	TotalRevenue        int                  `json:"total_revenue"`
	MonthlyTransactions []MonthlyTransaction `json:"monthly_transactions"`
	UserRanking         []UserRanking        `json:"user_ranking"`
}
