package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"shop-api/internal/dto"
	"shop-api/internal/mappers"
	"shop-api/internal/repository"
)

type ProductHandler struct {
	repo repository.ProductRepository
}

func NewProductHandler(repo repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

// CreateProduct godoc
// @Summary      Добавить товар
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product body dto.CreateProductRequest true "Товар"
// @Success      201 {object} dto.ProductResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /products [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductRequest
	if err := decodeAndValidate(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	model, err := mappers.CreateProductRequestToModel(req)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid supplier_id")
		return
	}
	created, err := h.repo.Create(r.Context(), &model, req.Category)
	if err != nil {
		handleRepoError(w, err)
		return
	}
	// доменный ответ
	full := mappers.ProductFullToResponse(structToFull(*created, req.Category))
	_ = full
	writeJSON(w, http.StatusCreated, mapProduct(*created, req.Category))
}

// ListProducts godoc
// @Summary      Получить все доступные товары
// @Tags         products
// @Produce      json
// @Param        limit  query int false "Лимит"
// @Param        offset query int false "Смещение"
// @Success      200 {array} dto.ProductResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /products [get]
func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	limit := parseIntQuery(r, "limit")
	offset := parseIntQuery(r, "offset")
	items, err := h.repo.ListAvailable(r.Context(), limit, offset)
	if err != nil {
		handleRepoError(w, err)
		return
	}
	resp := make([]dto.ProductResponse, 0, len(items))
	for _, p := range items {
		resp = append(resp, mappers.ProductFullToResponse(p))
	}
	writeJSON(w, http.StatusOK, resp)
}

// GetProduct godoc
// @Summary      Получить товар по id
// @Tags         products
// @Produce      json
// @Param        id path string true "UUID товара"
// @Success      200 {object} dto.ProductResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /products/{id} [get]
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	p, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		handleRepoError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, mappers.ProductFullToResponse(*p))
}

// DeleteProduct godoc
// @Summary      Удалить товар
// @Tags         products
// @Param        id path string true "UUID товара"
// @Success      204
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	if err := h.repo.Delete(r.Context(), id); err != nil {
		handleRepoError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// DecreaseStock godoc
// @Summary      Уменьшить количество товара
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id     path string                    true "UUID товара"
// @Param        amount body dto.DecreaseStockRequest  true "Сколько вычесть"
// @Success      200 {object} dto.ErrorResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /products/{id}/stock [patch]
func (h *ProductHandler) DecreaseStock(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	var req dto.DecreaseStockRequest
	if err := decodeAndValidate(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.repo.DecreaseStock(r.Context(), id, req.Amount); err != nil {
		if err.Error() == "not enough stock" {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		handleRepoError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "updated"})
}
