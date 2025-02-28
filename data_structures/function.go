package data_structures

import "fmt"

func Function() {
	usernames := make([]string, 2, 5)
	usernames[0] = "alice"
	usernames = append(usernames, "charlie")

	fmt.Println(usernames)
	for index, value := range usernames {
		fmt.Println(index, value)
	}

	courseRatings := make(map[string]float64, 3)
	courseRatings["Go"] = 5.0
	courseRatings["Java"] = 4.0
	courseRatings["Python"] = 3.0

	fmt.Println(courseRatings)

	for key, value := range courseRatings {
		fmt.Println(key, value)
	}
}
