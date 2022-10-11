package main

import (
	"fmt"

	"github.com/xm-tech/newmath"
)

type User struct {
	Name string
	Sex  int
	Age  int
}

func NewUser(Name string, Sex int, Age int) *User {
	user := &User{
		Name: Name,
		Sex:  Sex,
		Age:  Age,
	}
	return user
}

func (self *User) getName() string {
	return self.Name
}

func main() {
	fmt.Printf("vim-go,Sqrt(2) = %v \n", newmath.Sqrt(2))
	i := 11
	fmt.Printf("i: %v", i)

	user := NewUser("maxm", 0, 34)
	fmt.Printf("user = %+v\n", user)
	fmt.Printf("user.name = %+v\n", user.getName())
}
