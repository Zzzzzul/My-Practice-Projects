package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Address
}

func main() {
	user1 := User{
		ID:       1,
		UserName: "test1",
		Email:    "test123@test.com",
		Address: Address{
			Street:  "123 golang st.",
			City:    "Golang City",
			ZipCode: "12333",
		},
	}
	jsonData, err := json.MarshalIndent(user1, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println("✅ Marshaled JSON:")
	fmt.Println(string(jsonData))

	var user2 User
	err = json.Unmarshal(jsonData, &user2)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Println("\n✅ Unmarshaled Struct:")
	fmt.Printf("%+v\n", user2)
}
