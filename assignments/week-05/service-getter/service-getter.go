package service_getter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	httpclient "../http-client"
)

const (
	getPostsEndpoint    = "https://my-json-server.typicode.com/typicode/demo/posts"
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

type ServiceGetter interface {
	GetPosts() ([]Post, error)
	GetComments() ([]Comment, error)
}

type Post struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}
type Comment struct {
	ID     int64  `json:"id"`
	Body   string `json:"body"`
	PostID int64  `json:"postId"`
}

type serviceGetterImpl struct {
	httpClient httpclient.HTTPClient
}

func (sg *serviceGetterImpl) GetComments() ([]Comment, error) {
	resp, err := sg.httpClient.Get(getCommentsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var comments []Comment
	if err = json.Unmarshal(body, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}

func (sg *serviceGetterImpl) GetPosts() ([]Post, error) {
	resp, err := sg.httpClient.Get(getPostsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var posts []Post
	if err = json.Unmarshal(body, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func New() (ServiceGetter, error) {
	httpClient := http.DefaultClient
	return &serviceGetterImpl{
		httpClient: httpClient,
	}, nil
}
