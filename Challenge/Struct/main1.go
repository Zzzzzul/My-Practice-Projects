package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
	Email string
}

func PrintContacts(contacts []Contact) {
	fmt.Println("📇 Contact List:")
	fmt.Println("-----------------------")
	for _, c := range contacts {
		fmt.Printf("Name:  %s\n", c.Name)
		fmt.Printf("Phone: %s\n", c.Phone)
		fmt.Printf("Email: %s\n", c.Email)
		fmt.Println()
	}
}

func main() {
	var contacts []Contact

	contacts = append(contacts, Contact{
		Name:  "Alice Johnson",
		Phone: "0917-123-4567",
		Email: "alice@example.com",
	})

	contacts = append(contacts, Contact{
		Name:  "Bob Smith",
		Phone: "0917-123-1234",
		Email: "bob@example.com",
	})

	PrintContacts(contacts)
}
