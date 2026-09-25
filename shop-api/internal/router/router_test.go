package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"shop-api/internal/handlers"
	"shop-api/internal/router"
)

// Пустые репозитории — роутеру всё равно, лишь бы интерфейсы удовлетворяли.
func TestRouter_RegistersAllRoutes(t *testing.T) {
	// используем nil-интерфейсы? нет — лучше пустые моки. Но роутер хранит конкретные типы.
	// Проще всего: создать хендлеры с nil-репозиториями — NewXxxHandler принимает интерфейс.
	// Репозиторий в хендлере не вызывается до запроса, так что nil безопасен для проверки матчинга.
	h := router.Handlers{
		Client:   handlers.NewClientHandler(nil),
		Product:  handlers.NewProductHandler(nil),
		Supplier: handlers.NewSupplierHandler(nil),
		Image:    handlers.NewImageHandler(nil),
	}
	r := router.New(h)

	cases := []struct {
		method string
		path   string
	}{
		{http.MethodPost, "/api/v1/clients"},
		{http.MethodGet, "/api/v1/clients"},
		{http.MethodDelete, "/api/v1/clients/00000000-0000-0000-0000-000000000000"},
		{http.MethodPatch, "/api/v1/clients/00000000-0000-0000-0000-000000000000/address"},
		{http.MethodPost, "/api/v1/products"},
		{http.MethodGet, "/api/v1/products"},
		{http.MethodGet, "/api/v1/products/00000000-0000-0000-0000-000000000000"},
		{http.MethodDelete, "/api/v1/products/00000000-0000-0000-0000-000000000000"},
		{http.MethodPatch, "/api/v1/products/00000000-0000-0000-0000-000000000000/stock"},
		{http.MethodPost, "/api/v1/suppliers"},
		{http.MethodGet, "/api/v1/suppliers"},
		{http.MethodGet, "/api/v1/suppliers/00000000-0000-0000-0000-000000000000"},
		{http.MethodDelete, "/api/v1/suppliers/00000000-0000-0000-0000-000000000000"},
		{http.MethodPatch, "/api/v1/suppliers/00000000-0000-0000-0000-000000000000/address"},
		{http.MethodPost, "/api/v1/images"},
		{http.MethodGet, "/api/v1/images/00000000-0000-0000-0000-000000000000"},
		{http.MethodPut, "/api/v1/images/00000000-0000-0000-0000-000000000000"},
		{http.MethodDelete, "/api/v1/images/00000000-0000-0000-0000-000000000000"},
		{http.MethodGet, "/api/v1/products/00000000-0000-0000-0000-000000000000/image"},
	}

	for _, c := range cases {
		t.Run(c.method+" "+c.path, func(t *testing.T) {
			req := httptest.NewRequest(c.method, c.path, nil)
			match := &mux.RouteMatch{}
			assert.True(t, r.(*mux.Router).Match(req, match),
				"route not registered: %s %s", c.method, c.path)
		})
	}
}
