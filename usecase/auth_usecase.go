package usecase

import (
	"fmt"

	"github.com/fadedreams/jrpgc/entity"
	"github.com/fadedreams/jrpgc/repository"
	//"github.com/golang-jwt/jwt"
	//"gorm.io/gorm"
	//"time"
)

type AuthUsecase struct {
	authRepository repository.IAuthRepository
	storage        *Storage
}

func NewAuthUsecase(authRepository repository.IAuthRepository) *AuthUsecase {
	return &AuthUsecase{
		authRepository: authRepository,
	}
}

func (u *AuthUsecase) Migrate() bool {
	status := u.authRepository.Migrate()
	return status
}

func (u *AuthUsecase) TestPost() {
	fmt.Println("test post")
}

func (u *AuthUsecase) SignUp(email, password string) (*entity.UserSignUp, error) {
	a := &entity.UserSignUp{
		Email:    email,
		Password: password,
	}
	_, err := u.authRepository.SignUp(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (u *AuthUsecase) SignIn(email, password string) (*entity.User, error) {
	// Call the repository to verify the user's credentials
	a := &entity.User{
		Email:    email,
		Password: password,
		Name:     "",
		Role:     "",
	}
	user, err := u.authRepository.SignIn(a)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//func (u *AuthUsecase) GetPost(id int) (*entity.Post, error) {
//post, err := u.authRepository.GetPost(id)
//if err != nil {
//return nil, err
//}

//return post, nil
//}

//func (u *AuthUsecase) GetPosts() ([]entity.Post, error) {
//posts, err := u.authRepository.GetPosts()
//if err != nil {
//return nil, err
//}

//return posts, nil
//}

//func (u *AuthUsecase) UpdatePost(id int, title, content string) error {
//post, err := u.authRepository.GetPost(id)
//if err != nil {
//return err
//}
//// Check if the title or content need to be updated
//if title != "" {
//post.Title = title
//}
//if content != "" {
//post.Content = content
//}
//err = u.authRepository.UpdatePost(post)
//if err != nil {
//return err
//}
//return nil
//}

//func (u *AuthUsecase) DeletePost(id int) error {
//err := u.authRepository.DeletePost(id)
//if err != nil {
//return err
//}
//return nil
//}
