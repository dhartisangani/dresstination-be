package models

type Product struct {
	ID       string  `json:"id"`
	Image    string  `json:"image"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
	Link     string  `json:"link"`
}

type OutfitElement struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageLink   string    `json:"image_link"`
	ImagePrompt string 	  `json:"image_prompt"`
	SearchQuery string    `json:"search_query"`
	Products    []Product `json:"products"`
}

type Outfit struct {
	ID             string          `json:"id"`
	Title          string          `json:"title"`
	OutfitElements []OutfitElement `json:"outfit_elements"`
}