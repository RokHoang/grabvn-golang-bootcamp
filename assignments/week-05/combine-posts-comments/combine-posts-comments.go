package service_getter

import (
	"log"

	service_getter "../service-getter"
)

type PostWithComments struct {
	ID       int64                    `json:"id"`
	Title    string                   `json:"string"`
	Comments []service_getter.Comment `json:"comments,omitempty"`
}

type CombinePostsComments interface {
	CombinePostWithComments() ([]PostWithComments, error)
}

type combinePostsCommentsImpl struct {
	serviceGetter service_getter.ServiceGetter
}

func (sg *combinePostsCommentsImpl) CombinePostWithComments() ([]PostWithComments, error) {
	// Get posts from api
	posts, err := sg.serviceGetter.GetPosts()
	if err != nil {
		log.Println("get posts failed with error: ", err)
		// writer.WriteHeader(500)
		return nil, err
	}

	// Get comments from api
	comments, err := sg.serviceGetter.GetComments()
	if err != nil {
		log.Println("get comments failed with error: ", err)
		// writer.WriteHeader(500)
		return nil, err
	}
	commentsByPostID := map[int64][]service_getter.Comment{}
	for _, comment := range comments {
		commentsByPostID[comment.PostID] = append(commentsByPostID[comment.PostID], comment)
	}

	result := make([]PostWithComments, 0, len(posts))
	for _, post := range posts {
		result = append(result, PostWithComments{
			ID:       post.ID,
			Title:    post.Title,
			Comments: commentsByPostID[post.ID],
		})
	}

	return result, nil
}

func New() (CombinePostsComments, error) {
	serviceGetter, serviceGetterErr := service_getter.New()
	if serviceGetterErr != nil {
		log.Println("ERROR: failed to init service getter")
		return nil, serviceGetterErr
	}
	return &combinePostsCommentsImpl{
		serviceGetter: serviceGetter,
	}, nil
}
