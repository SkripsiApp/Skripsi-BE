package response

type ChatbotResponse struct {
	Answer         string           `json:"answer"`
	Recommendation []Recommendation `json:"recommendation"`
}

type Recommendation struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Sold  int    `json:"sold"`
	Image string `json:"image"`
}
