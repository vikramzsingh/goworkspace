// Hands on exercise 4 video 106

//Create and use an anonymous struct

package main

import "fmt"

func main() {
	s := struct {
		first    string
		city     map[string][]string
		number   []int
		age      int
		learning bool
	}{
		first: "Vikram",
		city: map[string][]string{
			"favcity":   {"Pathankot", "Jalander", "goa"}, // slic of string or
			"favmovies": []string{"Don", "Hum"},
		},
		number:   []int{9888709028, 9888855555},
		age:      23,
		learning: true,
	}

	fmt.Println(s.first)
	fmt.Println(s.city)
	fmt.Println(s.number)
	for k, v := range s.city {
		fmt.Println("\nKey is:", k, "and Value is:", v)
		for i, val := range v {
			fmt.Println(i, val)
		}
	}
	fmt.Println()
	for _, v := range s.number {
		fmt.Println("Number:", v)
	}
	fmt.Println()
	fmt.Println("age:", s.age)
	fmt.Println("Education:", s.learning)

}
