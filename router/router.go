package router

import (
	"simple-rest-mux/handlers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	articlesHandler := handlers.ArticlesHandlerWrapper

	router.HandleFunc("/articles", articlesHandler.GetAllArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", articlesHandler.GetArticleByID).Methods("GET")
	router.HandleFunc("/articles", articlesHandler.CreateNewArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", articlesHandler.UpdateArticleByID).Methods("PUT")
	router.HandleFunc("/articles/{id}", articlesHandler.DeleteArticleByID).Methods("DELETE")

	return router
}
