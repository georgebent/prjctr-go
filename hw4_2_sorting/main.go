package main

import (
	"fmt"
	"sort"
)

type User struct {
	Id int `json:"id"`
}

func (u User) String() string {
	return fmt.Sprintf("%d;", u.Id)
}

func main() {
	users := createUsers()
	fmt.Printf("Users: %s\n", users)

	uniqueUsers := getUniqueUsers(users)
	fmt.Printf("Unique: %s\n", uniqueUsers)

	sort.Slice(uniqueUsers, func(i, j int) bool {
		return uniqueUsers[i].Id < uniqueUsers[j].Id
	})

	fmt.Printf("Sorted: %s\n", uniqueUsers)
}

func isExist(users []User, user User) bool {
	for _, userItem := range users {
		if user.Id == userItem.Id {
			return true
		}
	}

	return false
}

func getUniqueUsers(users []User) []User {
	var uniqueUsers []User

	for _, user := range users {
		if !isExist(uniqueUsers, user) {
			uniqueUsers = append(uniqueUsers, user)
		}
	}

	return uniqueUsers
}

func createUsers() []User {
	return []User{
		{Id: 10},
		{Id: 1},
		{Id: 3},
		{Id: 5},
		{Id: 1},
		{Id: 1},
		{Id: 8},
		{Id: 10},
		{Id: 3},
		{Id: 2},
		{Id: 4},
		{Id: 1},
	}
}
