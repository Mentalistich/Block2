package main

import (
	"crypto/rand"
	"fmt"
	"strings"
)

type User struct {
	Username           string
	Password           string
	Role               string
	FailedLoginAttemps int
}

func (u *User) Authenticate(password string) {
	if u.Password == "" {
		u.ResetPassword()
	}
	if u.Password == password {
		fmt.Println("Success")
	} else {
		u.FailedLoginAttemps++
		fmt.Println("Failed")
	}
}

func (u *User) ChangePassword(newPassword string) {
	if u.Password == newPassword {
		fmt.Println("Mistake. Enter a new password")
	} else {
		u.Password = newPassword
		fmt.Println("Password has been changed")
	}
}

func (u *User) ResetPassword() {
	if u.FailedLoginAttemps > 1 || u.Password == "" {
		var input string
		fmt.Println("Enter a new password")
		fmt.Scanln(&input)
		u.ChangePassword(input)
	}
}

func (u *User) HasAcces(resource string) bool {
	if u.Role == "Admin" && (resource == "admin panel" || resource == "website") {
		return true
	}
	if u.Role == "User" && resource == "website" {
		return true
	}
	return false
}

func (u *User) BlockUser() {
	if u.FailedLoginAttemps > 3 {
		u.Role = "Locked in"
	}
}

func (u *User) SendOTP() {
	lenKey := 16
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	byteKey := make([]byte, lenKey)
	rand.Read(byteKey)

	var keyBuilder strings.Builder
	for i, b := range byteKey {
		if i > 0 && i%4 == 0 {
			// Добавляем дефис каждые 4 символа
			keyBuilder.WriteRune('-')
		}
		keyBuilder.WriteByte(charset[int(b)%len(charset)])
	}
	//или вот так
	//temp:=u.Password
	u.Password = keyBuilder.String()
	u.Authenticate(u.Password)
	//u.Password=temp
	u.Password = ""
}

func main() {

}
