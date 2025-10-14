package dto

type CreateProductRequest struct {
	Active       bool    `json:"active" validate:"required"`
	Category     string  `json:"category" validate:"required"`
	ImageUrl     string  `json:"image_url"`
	Name         string  `json:"name" validate:"required"`
	NameProvider string  `json:"name_provider" validate:"required"`
	Price        float64 `json:"price" validate:"required"`
	Sku          string  `json:"sku" validate:"required"`
}

type ChangeStockProductRequest struct {
	Sku   string `json:"sku" validate:"required"`
	Stock int    `json:"stock" validate:"required"`
	Type  string `json:"type" validate:"required"`
}

type UpdateProductRequest struct {
	Sku      string  `json:"sku" validate:"required"`
	ImageUrl string  `json:"image_url"`
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	Active   bool    `json:"active" validate:"required"`
}

type DeleteProductRequest struct {
	Sku string `json:"sku" validate:"required"`
}

// ErrorResponse representa una respuesta de error
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// SuccessResponse representa una respuesta exitosa genérica
type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
