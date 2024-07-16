package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// 生成 bcrypt 哈希
	password := []byte("mypassword")
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Hashing error:", err)
		return
	}
	fmt.Println("Hashed password:", string(hashedPassword))

	// 比较哈希和明文密码
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		fmt.Println("Invalid password")
	} else {
		fmt.Println("Password is valid")
	}
}
