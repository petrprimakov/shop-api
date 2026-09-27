package router

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"shop-api/internal/handlers"
)

type Handlers struct {
	Client   *handlers.ClientHandler
	Product  *handlers.ProductHandler
	Supplier *handlers.SupplierHandler
	Image    *handlers.ImageHandler
}

func New(h Handlers) http.Handler {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	// clients
	api.HandleFunc("/clients", h.Client.CreateClient).Methods(http.MethodPost)
	api.HandleFunc("/clients", h.Client.ListClients).Methods(http.MethodGet)
	api.HandleFunc("/clients/{id}", h.Client.DeleteClient).Methods(http.MethodDelete)
	api.HandleFunc("/clients/{id}/address", h.Client.UpdateClientAddress).Methods(http.MethodPatch)

	// products
	api.HandleFunc("/products", h.Product.CreateProduct).Methods(http.MethodPost)
	api.HandleFunc("/products", h.Product.ListProducts).Methods(http.MethodGet)
	api.HandleFunc("/products/{id}", h.Product.GetProduct).Methods(http.MethodGet)
	api.HandleFunc("/products/{id}", h.Product.DeleteProduct).Methods(http.MethodDelete)
	api.HandleFunc("/products/{id}/stock", h.Product.DecreaseStock).Methods(http.MethodPatch)

	// suppliers
	api.HandleFunc("/suppliers", h.Supplier.CreateSupplier).Methods(http.MethodPost)
	api.HandleFunc("/suppliers", h.Supplier.ListSuppliers).Methods(http.MethodGet)
	api.HandleFunc("/suppliers/{id}", h.Supplier.GetSupplier).Methods(http.MethodGet)
	api.HandleFunc("/suppliers/{id}", h.Supplier.DeleteSupplier).Methods(http.MethodDelete)
	api.HandleFunc("/suppliers/{id}/address", h.Supplier.UpdateSupplierAddress).Methods(http.MethodPatch)

	// images
	api.HandleFunc("/images", h.Image.UploadImage).Methods(http.MethodPost)
	api.HandleFunc("/images/{id}", h.Image.GetImage).Methods(http.MethodGet)
	api.HandleFunc("/images/{id}", h.Image.UpdateImage).Methods(http.MethodPut)
	api.HandleFunc("/images/{id}", h.Image.DeleteImage).Methods(http.MethodDelete)
	api.HandleFunc("/products/{id}/image", h.Image.GetProductImage).Methods(http.MethodGet)

	// swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r
}
