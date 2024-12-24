package response

type ChatbotResponse struct {
	Answer string   `json:"answer"`
	Images []string `json:"images"`
}