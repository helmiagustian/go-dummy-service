package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/helmiagustian/go-dummy-service/handler"
)

func main() {
	h := handler.NewHandler()

	router := chi.NewRouter()

	router.Get("/articles", h.GetArticles)
	router.Post("/articles", h.CreateArticle)
	router.Get("/articles/{id}", h.GetArticleByID)
	router.Put("/articles/{id}", h.UpdateArticleByID)
	router.Delete("/articles/{id}", h.DeleteArticleByID)

	fmt.Println("Serve at 1234")
	http.ListenAndServe(":1234", router)
}
