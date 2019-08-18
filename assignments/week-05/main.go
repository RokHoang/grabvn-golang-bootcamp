package main

import (
	"encoding/json"
	"log"
	"net/http"

	combinePostWithComments "./combine-posts-comments"
)

type PostWithCommentsResponse struct {
	Posts []combinePostWithComments.PostWithComments `json:"posts"`
}

//TODO: how to separate API logic, business logic and response format logic
func main() {
	combinePostWithCommentsService, combinePostWithCommentsErr := combinePostWithComments.New()

	if combinePostWithCommentsErr != nil {
		log.Println("unable to parse response: ", combinePostWithCommentsErr)
		return
	}

	http.HandleFunc("/postWithComments", func(writer http.ResponseWriter, request *http.Request) {

		// Combine and return response
		postWithComments, postWithCommentsErr := combinePostWithCommentsService.CombinePostWithComments()
		if postWithCommentsErr != nil {
			log.Println("unable to parse response: ", postWithCommentsErr)
			writer.WriteHeader(500)
		}
		resp := PostWithCommentsResponse{Posts: postWithComments}
		buf, err := json.Marshal(resp)
		if err != nil {
			log.Println("unable to parse response: ", err)
			writer.WriteHeader(500)
		}

		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(buf)
	})

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
