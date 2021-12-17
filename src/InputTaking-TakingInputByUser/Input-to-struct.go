package main

import "fmt"

type User struct {
	Name string
	Age int
	Branch string
	EmailID string
}

func main() {
	u := User{}
	fmt.Println("Enter the Name:")
	fmt.Scanln(&u.Name)
	fmt.Println("Enter the Age:")
	fmt.Scan(&u.Age)
	fmt.Println("Enter the Branch:")
	fmt.Scanln(&u.Branch)
	fmt.Println("Enter the EmailID:")
	fmt.Scanln(&u.EmailID)
	fmt.Println(u)

	var data string
	fmt.Println("Enter any Data:")
	fmt.Scanln(&data)
	fmt.Println(data)
}
