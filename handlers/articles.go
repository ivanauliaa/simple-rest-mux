package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"simple-rest-mux/data"
	"simple-rest-mux/utils"

	"github.com/gorilla/mux"
)

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: All Articles Request")

	response := utils.ResponseFormatter{
		Message: "Successfully get all articles",
		Data:    data.Articles,
		Status:  200,
	}
	json.NewEncoder(w).Encode(response)
}

func getArticleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Printf("Endpoint hit: Article with ID: %s request\n", key)

	for _, article := range *(data.Articles) {
		if article.ID == key {
			response := utils.ResponseFormatter{
				Message: fmt.Sprintf("Successfully get article with ID: %s", key),
				Data:    article,
				Status:  200,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	response := utils.ResponseFormatter{
		Message: "ID not found",
		Status:  404,
	}
	json.NewEncoder(w).Encode(response)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: Create New Article Request")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	// create article request curl example
	// curl -X POST -d '{"id": "3", "title": "New Title", "desc": "New Desc", "content": "New Content"}' localhost:8081/articles

	article := data.Article{}
	json.Unmarshal(reqBody, &article)
	*(data.Articles) = append(*(data.Articles), article)

	response := utils.ResponseFormatter{
		Message: "Successfully create new article",
		Data:    article,
		Status:  200,
	}
	json.NewEncoder(w).Encode(response)
}

func updateArticleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Printf("Endpoint hit: Update Article with ID: %s request\n", key)

	// update article request curl example
	// curl -X PUT -d '{"title": "Updated Title", "desc": "Updated Desc", "content": "Updated Content"}' localhost:8081/articles/1

	for index, article := range *(data.Articles) {
		if article.ID == key {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintln(w, err.Error())
			}

			article := data.Article{}
			article.ID = key

			json.Unmarshal(reqBody, &article)
			(*(data.Articles))[index] = article

			response := utils.ResponseFormatter{
				Message: fmt.Sprintf("Successfully update article with ID: %s", key),
				Data:    (*(data.Articles))[index],
				Status:  200,
			}
			json.NewEncoder(w).Encode(response)

			return
		}
	}

	response := utils.ResponseFormatter{
		Message: "ID not found",
		Status:  404,
	}
	json.NewEncoder(w).Encode(response)
}

func deleteArticleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Printf("Endpoint hit: Delete Article with ID: %s request\n", key)

	for index, article := range *(data.Articles) {
		if article.ID == key {
			*(data.Articles) = append((*(data.Articles))[:index], (*(data.Articles))[index+1:]...)

			response := utils.ResponseFormatter{
				Message: fmt.Sprintf("Succesfully delete article with ID: %s", key),
				Status:  200,
			}
			json.NewEncoder(w).Encode(response)

			return
		}
	}

	response := utils.ResponseFormatter{
		Message: "ID not found",
		Status:  404,
	}
	json.NewEncoder(w).Encode(response)
}

var ArticlesHandlerWrapper = struct {
	GetAllArticles    func(w http.ResponseWriter, r *http.Request)
	GetArticleByID    func(w http.ResponseWriter, r *http.Request)
	CreateNewArticle  func(w http.ResponseWriter, r *http.Request)
	UpdateArticleByID func(w http.ResponseWriter, r *http.Request)
	DeleteArticleByID func(w http.ResponseWriter, r *http.Request)
}{
	GetAllArticles:    getAllArticles,
	GetArticleByID:    getArticleByID,
	CreateNewArticle:  createNewArticle,
	UpdateArticleByID: updateArticleByID,
	DeleteArticleByID: deleteArticleByID,
}
