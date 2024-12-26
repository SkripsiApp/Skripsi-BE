package service

import (
	"context"
	"fmt"
	"skripsi/features/chatbot/entity"
	"skripsi/features/chatbot/interfaces"
	product "skripsi/features/product/interfaces"
	"skripsi/utils/helper"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

type ChatbotUseCase struct {
	openAIKey    string
	productRepo  product.ProductRepositoryInterface
	faqResponses map[string]string
	imageURLs    map[string]bool
}

func NewChatbotService(openAIKey string, productRepo product.ProductRepositoryInterface) interfaces.ChatbotServiceInterface {
	return &ChatbotUseCase{
		openAIKey:   openAIKey,
		productRepo: productRepo,
		faqResponses: map[string]string{
			"mandi":       "Ya, aksesoris titanium kami dapat dipakai saat mandi. Titanium tahan terhadap air dan tidak akan berkarat atau berubah warna.",
			"sehari-hari": "Tentu, aksesoris titanium kami dirancang untuk penggunaan sehari-hari. Titanium sangat tahan lama dan nyaman dipakai dalam jangka waktu yang lama.",
			"alergi":      "Titanium adalah logam yang sangat hipoalergenik, artinya sangat kecil kemungkinan menimbulkan reaksi alergi. Namun, jika Anda memiliki alergi logam yang parah, sebaiknya konsultasikan dengan dokter sebelum menggunakan.",
		},
		imageURLs: make(map[string]bool),
	}
}

func (uc *ChatbotUseCase) HandleCustomerQuery(query string) (string, []entity.ProductRecommendation, error) {
	query = strings.ToLower(query)

	// Cek apakah query ada di FAQ atau mengandung kata kunci FAQ
	for keyword, response := range uc.faqResponses {
		if strings.Contains(query, keyword) {
			return response, nil, nil
		}
	}

	complaintKeywords := []string{"komplain", "masalah", "bantuan", "error", "keluhan", "mengadu", "mengeluh", "mengadukan", "mengeluhkan", "cacat", "rusak", "retur", "refund", "pengembalian", "pengembalian dana", "pengembalian produk", "penggantian", "penggantian produk", "penggantian barang", "penggantian dana", "penggantian order", "penggantian pesanan", "penggantian item", "penggantian produk", "penggantian pesanan"}
	for _, keyword := range complaintKeywords {
		if strings.Contains(query, keyword) {
			whatsappNumber := "6282189638011"
			whatsappLink := fmt.Sprintf("https://wa.me/%s", whatsappNumber)
			response := fmt.Sprintf(
				"Kami sangat menghargai kesabaran Anda dalam menghadapi situasi ini.\n\n"+
					"Untuk bantuan lebih lanjut:\n"+
					"1. Anda dapat langsung menghubungi admin kami melalui WhatsApp di nomor berikut: %s.\n"+
					"2. Copy tautan berikut untuk membuka WhatsApp: %s.\n"+
					"3. Atau klik tombol di bawah untuk langsung menghubungi kami melalui WhatsApp.\n\n"+
					"Admin kami akan dengan senang hati membantu menyelesaikan permasalahan Anda.",
				whatsappNumber, whatsappLink,
			)
			return response, nil, nil
		}
	}

	keywords := []string{"cincin", "kalung", "gelang", "aksesoris", "anting", "liontin"}
	matchedKeywords := []string{}
	for _, keyword := range keywords {
		if strings.Contains(strings.ToLower(query), keyword) {
			matchedKeywords = append(matchedKeywords, keyword)
		}
	}

	uc.imageURLs = make(map[string]bool)
	var allRecommendations []entity.ProductRecommendation

	var productInfo string
	if len(matchedKeywords) > 0 {
		_, recommendations := uc.getProductInfo(matchedKeywords)
		allRecommendations = append(allRecommendations, recommendations...)
	}

	ctx := context.Background()
	client := openai.NewClient(uc.openAIKey)
	model := "ft:gpt-3.5-turbo-1106:exoream::Ahi25eiT" // Model fine-tuned

	systemPrompt := `Anda adalah Tia, asisten penjualan untuk toko aksesoris titanium. Berikan informasi yang akurat dan ramah tentang produk kami. 
	Jika ada pertanyaan di luar pengetahuan Anda, mohon maaf dan arahkan pelanggan untuk menghubungi customer service.`

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: query,
		},
	}

	if productInfo != "" {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Informasi produk terkait: " + productInfo,
		})
	}

	request := entity.ChatRequest{
		Context:  ctx,
		Client:   client,
		Model:    model,
		Messages: messages,
	}

	resp, err := uc.GetCompletionFromMessages(request)
	if err != nil {
		return "", nil, err
	}

	answer := resp.Choices[0].Message.Content

	mentionedKeywords := []string{}
	for _, keyword := range keywords {
		if strings.Contains(strings.ToLower(answer), keyword) {
			mentionedKeywords = append(mentionedKeywords, keyword)
		}
	}

	if len(mentionedKeywords) > 0 {
		_, additionalRecommendations := uc.getProductInfo(mentionedKeywords)
		if len(additionalRecommendations) > 0 {
			allRecommendations = append(allRecommendations, additionalRecommendations...)
		}
	}

	return answer, allRecommendations, nil
}

func (uc *ChatbotUseCase) getProductInfo(keywords []string) (string, []entity.ProductRecommendation) {
	var productInfo string
	var recommendations []entity.ProductRecommendation
	for _, keyword := range keywords {
		titleCaseKeyword := helper.ToTitleCase(keyword)
		products, _, _, err := uc.productRepo.GetAll(titleCaseKeyword, 1, 2) // Limit to 2 products per category
		if err != nil {
			continue
		}
		if len(products) > 0 {
			productInfo += fmt.Sprintf("Produk %s:\n", titleCaseKeyword)
			for _, product := range products {
				productInfo += fmt.Sprintf("- %s, Harga: Rp%d\n", product.Name, product.Price)
				if product.Image != "" && !uc.imageURLs[product.Image] {
					uc.imageURLs[product.Image] = true
					recommendations = append(recommendations, entity.ProductRecommendation{
						Name:  product.Name,
						Price: product.Price,
						Image: product.Image,
					})
				}
			}
		}
	}
	return productInfo, recommendations
}

func (uc *ChatbotUseCase) GetCompletionFromMessages(request entity.ChatRequest) (openai.ChatCompletionResponse, error) {
	if request.Model == "" {
		request.Model = "ft:gpt-3.5-turbo-1106:exoream::Ahi25eiT"
	}

	resp, err := request.Client.CreateChatCompletion(
		request.Context,
		openai.ChatCompletionRequest{
			Model:    request.Model,
			Messages: request.Messages,
		},
	)
	return resp, err
}