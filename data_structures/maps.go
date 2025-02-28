package data_structures

import "fmt"

func Maps() {
	websites := map[string]string{
		"Google":   "www.google.com",
		"Facebook": "www.facebook.com",
		"Twitter":  "www.twitter.com",
	}

	fmt.Println(websites["Google"])

	websites["Youtube"] = "www.youtube.com"
	fmt.Println(websites)

	delete(websites, "Youtube")
	fmt.Println(websites)
}
