package usecase

import (
	"fmt"

	"github.com/fadedreams/jrpgc/entity"
	"github.com/fadedreams/jrpgc/repository"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}
type PostUsecase struct {
	postRepository repository.IPostRepository
	storage        *Storage
}

func NewPostUsecase(postRepository repository.IPostRepository) *PostUsecase {
	return &PostUsecase{
		postRepository: postRepository,
	}
}

func (u *PostUsecase) TestPost() {
	fmt.Println("test post")
}

func (u *PostUsecase) Migrate() bool {
	status := u.postRepository.Migrate()
	return status
}

func (u *PostUsecase) CreatePost(title, content string) (*entity.Post, error) {
	post := &entity.Post{
		Title:   title,
		Content: content,
	}
	//u.storage.db.Create(post)
	_, err := u.postRepository.CreatePost(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (u *PostUsecase) GetPost(id int) (*entity.Post, error) {
	post, err := u.postRepository.GetPost(id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (u *PostUsecase) GetPosts() ([]entity.Post, error) {
	posts, err := u.postRepository.GetPosts()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (u *PostUsecase) UpdatePost(id int, title, content string) error {
	post, err := u.postRepository.GetPost(id)
	if err != nil {
		return err
	}
	// Check if the title or content need to be updated
	if title != "" {
		post.Title = title
	}
	if content != "" {
		post.Content = content
	}
	err = u.postRepository.UpdatePost(post)
	if err != nil {
		return err
	}
	return nil
}

func (u *PostUsecase) DeletePost(id int) error {
	err := u.postRepository.DeletePost(id)
	if err != nil {
		return err
	}
	return nil
}
