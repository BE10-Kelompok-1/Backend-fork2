package domain

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	ID           int
	Firstname    string
	Lastname     string
	Username     string
	Email        string
	Password     string
	Birthdate    string
	Photoprofile string
	Posts        []Post
}

// type UserPosting struct {
// 	PostID    int
// 	Photo     string
// 	Caption   string
// 	CreatedAt time.Time
// }

// type UserPostingComment struct {
// 	CommentID    int
// 	Firstname    string
// 	Lastname     string
// 	Photoprofile string
// 	Comment      string
// 	CreatedAt    time.Time
// }

type UserHandler interface {
	Register() echo.HandlerFunc
	Update() echo.HandlerFunc
	Search() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
}

type UserUseCase interface {
	RegisterUser(newuser User, cost int) int
	UpdateUser(newuser User, userid, cost int) int
	SearchUser(username string) (User, int)
	DeleteUser(userid int) int
	LoginUser(userdata User) (User, error)
	ProfileUser(userid int) (User, error)
}

type UserData interface {
	RegisterData(newuser User) User
	UpdateUserData(newuser User) User
	SearchUserData(username string) User
	DeleteUserData(userid int) bool
	LoginData(userdata User) User
	GetPasswordData(name string) string
	CheckDuplicate(newuser User) bool
	// SearchUserPostingData(username string) []User
	// SearchUserPostingCommentData(username string) []User
	ProfileUserData(userid int) User
}
