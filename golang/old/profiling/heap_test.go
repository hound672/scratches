package main

import "testing"

const (
	usersCount = 100
)

var

type User struct {
	Name string
	Email string
	Value int
}

////go:noinline
func Factory(i int) *User {
	return &User{Name: "sss", Email: "mail@mail.ru", Value: i}
}

func ServiceValue() []User {
	users := make([]User, 0, usersCount)
	for i := 0; i < usersCount; i++ {
		user := Factory(i)
		users = append(users, *user)
	}

	return users
}

func ServicePointer() []*User {
	users := make([]*User, 0, usersCount)
	for i := 0; i < usersCount; i++ {
		user := Factory(i)
		users = append(users, user)
	}

	return users
}

func BenchmarkUsersValues(b *testing.B) {
	for i := 0; i < b.N; i++ {
		users := ServiceValue()
		_ = users
	}
}

func BenchmarkUsersPointers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		users := ServicePointer()
		_ = users
	}
}

func modify(u *User) *User {
	u.Email = ""
	return u
}

func BenchmarkModify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := &User{}
		user = modify(user)
	}
}
