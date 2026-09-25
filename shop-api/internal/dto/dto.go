package dto

type ErrorResponse struct {
	Error string `json:"error" example:"invalid request body"`
}

type AddressDTO struct {
	Country string `json:"country" validate:"required,min=1,max=100" example:"Russia"`
	City    string `json:"city"    validate:"required,min=1,max=100" example:"Moscow"`
	Street  string `json:"street"  validate:"required,min=1,max=200" example:"Tverskaya 1"`
}

// ---------- Client ----------

type CreateClientRequest struct {
	ClientName    string     `json:"client_name"    validate:"required,min=1,max=100"      example:"Ivan"`
	ClientSurname string     `json:"client_surname" validate:"required,min=1,max=100"      example:"Ivanov"`
	Birthday      string     `json:"birthday"       validate:"required,datetime=2006-01-02" example:"1990-05-15"`
	Gender        string     `json:"gender"         validate:"required,oneof=male female other" example:"male"`
	Address       AddressDTO `json:"address"        validate:"required"`
}

type ClientResponse struct {
	ID               string     `json:"id"                example:"550e8400-e29b-41d4-a716-446655440000"`
	ClientName       string     `json:"client_name"       example:"Ivan"`
	ClientSurname    string     `json:"client_surname"    example:"Ivanov"`
	Birthday         string     `json:"birthday"          example:"1990-05-15"`
	Gender           string     `json:"gender"            example:"male"`
	RegistrationDate string     `json:"registration_date" example:"2024-01-01T12:00:00Z"`
	Address          AddressDTO `json:"address"`
}

type UpdateAddressRequest struct {
	Country string `json:"country" validate:"required,min=1,max=100" example:"Russia"`
	City    string `json:"city"    validate:"required,min=1,max=100" example:"Moscow"`
	Street  string `json:"street"  validate:"required,min=1,max=200" example:"Tverskaya 1"`
}

// ---------- Product ----------

type CreateProductRequest struct {
	Name           string  `json:"name"            validate:"required,min=1,max=200" example:"Bosch KGN39"`
	Category       string  `json:"category"        validate:"required,min=1,max=100" example:"Fridges"`
	Price          float64 `json:"price"           validate:"required,gte=0"         example:"49999.99"`
	AvailableStock int     `json:"available_stock" validate:"gte=0"                  example:"10"`
	SupplierID     string  `json:"supplier_id"     validate:"required,uuid"          example:"550e8400-e29b-41d4-a716-446655440000"`
}

type ProductResponse struct {
	ID             string  `json:"id"               example:"550e8400-e29b-41d4-a716-446655440000"`
	Name           string  `json:"name"             example:"Bosch KGN39"`
	Category       string  `json:"category"         example:"Fridges"`
	Price          float64 `json:"price"            example:"49999.99"`
	AvailableStock int     `json:"available_stock"  example:"10"`
	LastUpdateDate string  `json:"last_update_date" example:"2024-01-01T12:00:00Z"`
	SupplierID     string  `json:"supplier_id"      example:"550e8400-e29b-41d4-a716-446655440000"`
	ImageID        *string `json:"image_id,omitempty"`
}

type DecreaseStockRequest struct {
	Amount int `json:"amount" validate:"required,gt=0" example:"1"`
}

// ---------- Supplier ----------

type CreateSupplierRequest struct {
	Name        string     `json:"name"         validate:"required,min=1,max=200" example:"Bosch"`
	Address     AddressDTO `json:"address"      validate:"required"`
	PhoneNumber string     `json:"phone_number" validate:"required,min=5,max=20" example:"+7-800-555-35-35"`
}

type SupplierResponse struct {
	ID          string     `json:"id"           example:"550e8400-e29b-41d4-a716-446655440000"`
	Name        string     `json:"name"         example:"Bosch"`
	Address     AddressDTO `json:"address"`
	PhoneNumber string     `json:"phone_number" example:"+7-800-555-35-35"`
}

// ---------- Image ----------

type ImageResponse struct {
	ID        string `json:"id"         example:"550e8400-e29b-41d4-a716-446655440000"`
	ProductID string `json:"product_id" example:"660e8400-e29b-41d4-a716-446655440000"`
}
