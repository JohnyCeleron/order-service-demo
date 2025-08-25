package domain

type OrderItem struct {
	ChrtID      int    `json:"chrt_id" validate:"required" example:"9934930"`
	TrackNumber string `json:"track_number" validate:"required" example:"WBILMTESTTRACK"`
	Price       int    `json:"price" validate:"required" example:"453"`
	Rid         string `json:"rid" validate:"required" example:"ab4219087a764ae0btest"`
	Name        string `json:"name" validate:"required" example:"Mascaras"`
	Sale        int    `json:"sale" validate:"required" example:"30"`
	Size        string `json:"size" validate:"required" example:"0"`
	TotalPrice  int    `json:"total_price" validate:"required" example:"317"`
	NmID        int    `json:"nm_id" validate:"required" example:"2389212"`
	Brand       string `json:"brand" validate:"required" example:"Vivienne Sabo"`
	Status      int    `json:"status" validate:"required" example: "202"`
}
