package api

import (
	// 	"encoding/json"
	"encoding/json"
	"net/http"

	// 	"github.com/go-chi/chi/v5"
	"github.com/YagoNigro123/url-shortener/internal/core"
	"github.com/go-chi/chi/v5"
)

type CreateLinkRequest struct {
	OriginalURL string `json:"original_url"`
}

type CreateLinkResponse struct {
	ID       string `json:"id"`
	ShortURL string `json:"short_url"`
}

type Handler struct {
	svc *core.Service
}

func NewHandler(svc *core.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateLink(w http.ResponseWriter, r *http.Request) {
	var req CreateLinkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	link, err := h.svc.Shorten(req.OriginalURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp := CreateLinkResponse{
		ID:       link.ID,
		ShortURL: "http://localhost:8080/" + link.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	link, err := h.svc.GetOriginal(id)
	if err != nil {
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, link.Original, http.StatusMovedPermanently)
}
