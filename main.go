package main

import "fmt"

type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

func main() {
	user1 := User{"vasya", 23, "m", 75, 105}

	fmt.Printf("%+v\n", user1.name)
}
