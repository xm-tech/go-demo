package main

import "encoding/json"
import "fmt"

type User struct {
	Name string `json:"name"`
	Sex  int    `json:"sex"`
	Age  int    `json:"age"`
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

func (self *User) getSex() int {
	return self.Sex
}

func main() {
	var user = NewUser("maxm", 0, 36)
	fmt.Printf("user = %+v\n", user)
	fmt.Printf("user.getName() = %+v\n", user.getName())
	fmt.Printf("user.getSex() = %+v", user.getSex())

	var ret, err = json.Marshal(&user)
	if err != nil {
		panic("json Marshal fail")
	}
	fmt.Printf("Marshal ret: %+v\n", ret)

	var ret2, err2 = json.Marshal(struct {
		*User
		City string `json:"city"`
	}{
		User: user,
		City: "hangzhou",
	})
	if err2 != nil {
		panic("json Marshal fail2")
	}
	fmt.Printf("Marshal ret2: = %+v\n", ret2)

	var user2 User
	var wrong = json.Unmarshal(ret, &user2)
	if wrong != nil {
		panic("json Unmarshal fail")
	}
	fmt.Printf("Unmarshal ret: %+v\n", user2)

	var users []*User
	var maxm = NewUser("maxm", 0, 35)
	var lfj = NewUser("lfj", 0, 35)
	users = append(users, maxm)
	users = append(users, lfj)
	var ret3, err3 = json.Marshal(users)
	if err3 != nil {
		panic("json Marshal users Fail")
	}
	fmt.Printf("ret3 = %+v\n", ret3)
	fmt.Printf("json(users) = %+v\n", string(ret3))
}
