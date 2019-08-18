package handler

import (
	"fmt"
	"github.com/helmiagustian/go-dummy-service/response"
	"net/http"
)

// Handler hold all API handler
type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// GetArticles handle get articles
func (h *Handler) GetArticles(w http.ResponseWriter, r *http.Request) error {
	return response.Write(w, response.SuccessDefault())
}

// CreateArticle handle create article
func (h *Handler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

// GetArticleByID handle get article by id
func (h *Handler) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

// UpdateArticleByID handle update article by id
func (h *Handler) UpdateArticleByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

// DeleteArticleByID handle delete article by id
func (h *Handler) DeleteArticleByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
