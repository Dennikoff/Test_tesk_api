package model

func TestUser() *User {
	return &User{
		Email:    "user@example.com",
		Password: "password",
		NickName: "nickname",
	}
}
