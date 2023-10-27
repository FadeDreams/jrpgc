package entity

type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type CreatePostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type IPostRepo interface {
	//createPost(*Post) (*Post, error)
	persistPost(*Post) (bool, error)
	getPost(int) (*Post, error)
	getPosts(int) ([]Post, error)
}

type UpdatePostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
