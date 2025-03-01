package product

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yikuanzz/ecom/model"
	"github.com/yikuanzz/ecom/utils"
)

type Handler struct {
	store model.ProductStore
}

func NewHandler(store model.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodGet)
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodPost)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get products"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}
