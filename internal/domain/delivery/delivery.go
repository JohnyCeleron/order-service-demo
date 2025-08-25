package domain

type Delivery struct {
	Name    string `json:"name" validate:"required" example:"Test Testov"`
	Phone   string `json:"phone" validate:"required" example:"+9720000000"`
	Zip     string `json:"zip" validate:"required" example:"2639809"`
	City    string `json:"city" validate:"required" example:"Kiryat Mozkin"`
	Address string `json:"address" validate:"required" example:"Ploshad Mira 15"`
	Region  string `json:"region" validate:"required" example:"Kraiot"`
	Email   string `json:"email" validate:"required" example:"test@gmail.com"`
}
