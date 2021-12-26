package controller

import (
	"encoding/json"
	"net/http"

	"github.com/xiemenger/GolangMock/entity"
	"github.com/xiemenger/GolangMock/errors"
	"github.com/xiemenger/GolangMock/service"
)

var (
	postService service.PostService = service.NewPostService()
)

type controller struct{}

type PostController interface {
	GetPosts(resp http.ResponseWriter, request *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

func NewPostController() PostController {
	return &controller{}
}

func (*controller) GetPosts(resp http.ResponseWriter, request *http.Request) {
	resp.Header().Set("Content-Type", "Application/json")
	posts, err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "fetching the posts array"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func (*controller) AddPost(resp http.ResponseWriter, req *http.Request) {
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error unmarshalling the request"})
		return
	}
	err = postService.Validate(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	_, err = postService.Create(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)
}
