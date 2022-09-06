package main

import "fmt"

type User struct {
	id   int
	name string
}

type Manager struct {
	User  //匿名属性
	title string
}

func (s *User) toString() string {
	return fmt.Sprintf("User: %p, %v", s, s)
}

func (s *Manager) toString() string {
	return fmt.Sprintf("Manager: %p, %v", s, s)
}

func fun01() {
	m := Manager{
		User{
			id:   1,
			name: "Tom",
		},
		"Administrator",
	}

	fmt.Println(m.toString())

	fmt.Println(m.User.toString())
}

func main() {
	fun01()

}
