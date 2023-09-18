package users

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"golang.org/x/exp/slices"
)

type User struct {
	Id         int
	First_name string
	Last_name  string
	Email      string
	Gender     string
	Ip_address string
}

func readUsers() [][]string {
	file, err := os.Open("users/users.csv")
	if err != nil {
		log.Fatalf("impossible to open file %s", err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	// Read header
	_, err = r.Read()

	if err != nil {
		log.Fatal(err)
	}

	rawUsers, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return rawUsers
}

func recordToUser(record []string) User {
	id, err := strconv.Atoi(record[0])

	if err != nil {
		log.Fatal(err)
	}

	return User{
		Id:         id,
		First_name: record[1],
		Last_name:  record[2],
		Email:      record[3],
		Gender:     record[4],
		Ip_address: record[5],
	}
}

func GetAllUsers() []User {
	users := make([]User, 6000)
	for _, user := range readUsers() {
		users = append(users, recordToUser(user))
	}

	return users
}

func GetAllUsersByEmail() map[string]User {
	users := make(map[string]User, 6000)
	for _, record := range readUsers() {
		user := recordToUser(record)

		users[user.Email] = user
	}

	return users
}

func SearchUserSlow(users []User, email string) (*User, error) {

	position := slices.IndexFunc(users, func(u User) bool {
		return u.Email == email
	})

	if position == -1 {
		return nil, fmt.Errorf("User with %s email not found", email)
	}

	return &users[position], nil
}

func SearchUserFast(users map[string]User, email string) (*User, error) {

	if user, ok := users[email]; ok {
		return &user, nil
	}

	return nil, fmt.Errorf("User with %s email not found", email)
}

func FastExample(email string) {
	byEmail := GetAllUsersByEmail()

	start := time.Now()

	found, err := SearchUserFast(byEmail, email)

	fmt.Println("Fast search in:", time.Since(start))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Found %s %s by %s email\n", found.First_name, found.Last_name, email)
}

func SlowExample(email string) {
	usersSlice := GetAllUsers()

	start := time.Now()

	found, err := SearchUserSlow(usersSlice, email)

	fmt.Println("Slow search in:", time.Since(start))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Found %s %s by %s email", found.First_name, found.Last_name, email)
}
