package main

import (
	"fmt"
	"time"
)

func main() {
	m := map[string]int{
		"x": 1,
		"y": 2,
		"z": 3,
		"t": time.Now().Day(),
	}
	resgetUser := getUser(m)
	for k, v := range resgetUser {
		fmt.Printf("Key: %s\tValue: %v\n", k, v)
	}
}

var m = map[string]int{
	"user_ID":      1,
	"user_age":     24,
	"user_created": time.Now().Day(),
}

func getUser(m map[string]int) map[string]int {
	setUsers := make(map[string]int)
	for k, v := range m {
		setUsers[k] = v * 2
	}
	return setUsers
}
