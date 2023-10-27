// repository/your_concrete_post_repository.go
package repository

import (
	//"encoding/json"
	"fmt"
	"github.com/fadedreams/jrpgc/entity"
	//"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	//"github.com/gorilla/mux"
	//"golang.org/x/crypto/bcrypt"
	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"log"
	//"net/http"
	//"time"
)

type IAuthRepository interface {
	SignUp(*entity.UserSignUp) (*entity.User, error)
	SignIn(*entity.User) (*entity.User, error)
	//isAuthorize() bool
	Migrate() bool
}

type ConcreteAuthRepository struct {
	storage *Storage
}

func NewConcreteAuthRepository(db *gorm.DB) *ConcreteAuthRepository {
	return &ConcreteAuthRepository{
		storage: &Storage{
			db: db,
		},
	}
}

func (c *ConcreteAuthRepository) Migrate() bool {
	fmt.Println("Migrating users...")
	err := c.storage.db.AutoMigrate(&entity.User{})
	if err != nil {
		return false
	}
	return true
	//s.db.AutoMigrate(&Post{})
}

//func (c *ConcreteAuthRepository) isAuthorize() bool {
//r := c.storage.db.Create(post)
//if r.Error != nil {
//return nil, r.Error
//}
//return post, nil
//}

//func (c *ConcreteAuthRepository) CheckPasswordHash(string, string) bool {
//var post entity.Post
//r := c.storage.db.First(&post, id)
////r.Row()
//if r.Error != nil {
//return nil, r.Error
//}
//return &post, nil
//}

//func (c *ConcreteAuthRepository) SignIn(*entity.User) (*entity.User, error) {
//var posts []entity.Post
//r := c.storage.db.Find(&posts)
//if r.Error != nil {
//return nil, r.Error
//}
//return posts, nil // Return the actual posts slice
//}

func (c *ConcreteAuthRepository) SignUp(userSignUp *entity.UserSignUp) (*entity.User, error) {
	// Check if a user with the same email already exists
	var existingUser entity.User
	result := c.storage.db.Where("email = ?", userSignUp.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		return nil, fmt.Errorf("User with email %s already exists", userSignUp.Email)
	}

	// Hash the password
	hashedPassword, err := GeneratehashPassword(userSignUp.Password)
	if err != nil {
		return nil, err
	}

	userSignUp.Password = hashedPassword

	// Create the user
	newUser := &entity.User{
		Email:    userSignUp.Email,
		Password: userSignUp.Password,
		Role:     "user", // Set the role to 'user'
	}

	result = c.storage.db.Create(newUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return newUser, nil
}

//func (c *ConcreteAuthRepository) GeneratehashPassword(password string) (string, error) {
//r := c.storage.db.Delete(&entity.Post{}, id)
//if r.Error != nil {
//return r.Error
//}
//return nil
//}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (c *ConcreteAuthRepository) SignIn(user *entity.User) (*entity.User, error) {
	// Find a user with the provided email
	var existingUser entity.User
	result := c.storage.db.Where("email = ?", user.Email).First(&existingUser)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("User with email %s does not exist", existingUser.Email)
	}

	// Check if the provided password matches the hashed password in the database
	if !CheckPasswordHash(user.Password, existingUser.Password) {
		return nil, fmt.Errorf("Invalid password")
	}

	// Password is correct, return the user
	return &existingUser, nil
}
