package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"../entity"
	"./entity"
	"./repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPosts(resp http.ResponseWriter, request *http.Request) {
	resp.Header().Set("Content-Type", "Application/json")
	posts, err := repo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error fetching the posts array"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func addPost(resp http.ResponseWriter, req *http.Request) {
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}
	post.Id = rand.Int()
	repo.Save(&post)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)
}
