package models

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type Role uint8

const (
	Undefined Role = iota
	Tutor
	Student
	Parent
)

type User struct {
	ID                int64  `json:"id"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Phone             string `json:"phone"`
	IsTutor           bool   `json:"isTutor"`
	IsStudent         bool   `json:"isStudent"`
	IsParent          bool   `json:"isParent"`
	ProfilePictureUrl string `json:"profilePictureUrl"`
}

func (u *User) GetID() int64 {
	return u.ID
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPasswordHash() string {
	return u.Password
}

func (u *User) GetPhone() string {
	return u.Phone
}

func (u *User) GetProfilePictureUrl() string {
	return u.ProfilePictureUrl
}

func (u *User) HasTutorRole() bool {
	return u.IsTutor
}

func (u *User) HasStudentRole() bool {
	return u.IsStudent
}

func (u *User) HasParentRole() bool {
	return u.IsParent
}

// This func needs to consume the current login status
// to compare the role that user has.
func (u *User) GenRole(sess *sessions.Sessions, ctx iris.Context) int {
	session := sessions.Get(ctx)

	// if not found then it returns just an empty string.
	currentRole, err := session.GetInt("currentRole")
	if err != nil {
		return 0 //Undefined
	}
	return currentRole
}

// genNullable(int64 userID): ?User
// genEnfore(int64 userID): User
func (u *User) GetUser(
	Username,
	Email,
	Password,
	Phone string,
	IsTutor bool,
	IsStudent bool,
	IsParent bool,
	ProfilePictureUrl string,
) User {
	return User{0, Username, Email, Password, Phone, IsTutor, IsStudent, IsParent, ProfilePictureUrl}
}
