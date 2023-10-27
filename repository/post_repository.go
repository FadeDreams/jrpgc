// repository/your_concrete_post_repository.go
package repository

import (
	"fmt"
	"github.com/fadedreams/jrpgc/entity"
	"gorm.io/gorm"
)

type IPostRepository interface {
	CreatePost(post *entity.Post) (*entity.Post, error)
	GetPost(id int) (*entity.Post, error)
	GetPosts() ([]entity.Post, error)
	UpdatePost(post *entity.Post) error
	DeletePost(id int) error
	Migrate() bool
}

type ConcretePostRepository struct {
	storage *Storage
}

type Storage struct {
	db *gorm.DB
}

func NewConcretePostRepository(db *gorm.DB) *ConcretePostRepository {
	return &ConcretePostRepository{
		storage: &Storage{
			db: db,
		},
	}
}

func (c *ConcretePostRepository) Migrate() bool {
	fmt.Println("Migrating...")
	err := c.storage.db.AutoMigrate(&entity.Post{})
	if err != nil {
		return false
	}
	return true
	//s.db.AutoMigrate(&Post{})
}

func (c *ConcretePostRepository) CreatePost(post *entity.Post) (*entity.Post, error) {
	r := c.storage.db.Create(post)
	if r.Error != nil {
		return nil, r.Error
	}
	return post, nil
}

//func (r *ConcretePostRepository) CreatePost(post *entity.Post) (*entity.Post, error) {
//// Implement the logic to create a new post in your database.
//return post, nil
//}

func (c *ConcretePostRepository) GetPost(id int) (*entity.Post, error) {
	var post entity.Post
	r := c.storage.db.First(&post, id)
	//r.Row()
	if r.Error != nil {
		return nil, r.Error
	}
	return &post, nil
}

func (c *ConcretePostRepository) GetPosts() ([]entity.Post, error) {
	var posts []entity.Post
	r := c.storage.db.Find(&posts)
	if r.Error != nil {
		return nil, r.Error
	}
	return posts, nil // Return the actual posts slice
}

func (c *ConcretePostRepository) UpdatePost(post *entity.Post) error {
	r := c.storage.db.Save(post)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func (c *ConcretePostRepository) DeletePost(id int) error {
	r := c.storage.db.Delete(&entity.Post{}, id)
	if r.Error != nil {
		return r.Error
	}
	return nil
}
