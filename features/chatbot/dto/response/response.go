package response

type ChatbotResponse struct {
	Answer         string           `json:"answer"`
	Recommendation []Recommendation `json:"recommendation"`
}

type Recommendation struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}
