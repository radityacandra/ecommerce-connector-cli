package model

type User struct {
	Id             string
	Name           string `firestore:"name"`
	Email          string `firestore:"email"`
	ProfilePicture string `firestore:"profile_picture"`
}

func NewUser(id, name, email, profilePicture string) *User {
	return &User{
		Id:             id,
		Name:           name,
		Email:          email,
		ProfilePicture: profilePicture,
	}
}
