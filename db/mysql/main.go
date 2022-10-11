package main

import (
	"fmt"

	gorm_ "github.com/xm-tech/go-demo/db/mysql/gorm_"
)

type User struct {
	Id   int64  `gorm:"id"`
	Name string `gorm:"name"`
	Sex  int    `gorm:"sex"`
}

func (user *User) initWithId(uid int) error {
	db := gorm_.GetDb()
	err := db.Where("id=?", uid).Limit(1).Find(&user).Error
	if err != nil {
		return fmt.Errorf("getUser err, id:%v,%v", uid, err)
	}
	return nil
}

func (user *User) Save() error {
	db := gorm_.GetDb()
	err := db.Save(&user).Error
	if err != nil {
		return fmt.Errorf("save err, id:%v,%v", user.Id, err)
	}
	return nil
}

func (user *User) del() error {
	db := gorm_.GetDb()
	err := db.Where("id=?", user.Id).Delete(&User{}).Error
	if err != nil {
		return fmt.Errorf("delUser err, user=%v", &user)
	}
	return nil
}

func delUsersByName(name string) error {
	db := gorm_.GetDb()
	err := db.Where("name=?", name).Delete(User{}).Error
	if err != nil {
		panic(err)
	}
	return nil
}

func getUsersByName(name string) ([]User, error) {
	db := gorm_.GetDb()
	var users []User
	err := db.Where("name=?", name).Find(&users).Error
	if err != nil {
		// panic(err)
		return nil, fmt.Errorf("getUsersByName err, name=%v", name)
	}
	return users, nil
}

func getUsersByNameWithRawSql(name string) (interface{}, error) {
	db := gorm_.GetDb()
	type UserProp struct {
		Name string `gorm:"name"`
		Sex  int    `gorm:"sex"`
	}

	var users []UserProp
	err := db.Raw("select name, sex from user where name =?", name).Scan(&users).Error
	if err != nil {
		return nil, fmt.Errorf("getUsersByNameWithRawSql err, name=%v", name)
	}
	return users, nil
}

func main() {
	// var db *gorm.DB = gorm_.GetDb()

	// if !db.HasTable(&User{}) {
	// 	log.Println("create table user")
	// 	db.AutoMigrate(&User{})
	// 	db.Model(&User{}).AddIndex("i_user_id", "user_id")
	// }

	// user0 := User{UserId: 1, Name: "maxm", Sex: 1}
	// user1 := User{UserId: 2, Name: "lfj", Sex: 0}
	// db.Create(&user0)
	// db.Create(&user1)

	// db.Model(&user1).Update("Name", "tanqin")

	// var tanqin User
	// db.Where("name=?", "tanqin").Find(&tanqin)
	// fmt.Printf("tanqin = %+v\n", tanqin)
	// up()
	// create()
	// delUser(2)
	// delUser(4)
	// delUser(6)
	// delUser(8)
	// delUser(10)
	// delUser(12)
	// users, err := getUsersByName("maxm")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("users = %+v\n", users)

	// delUsersByName("lqh")
	// create("tanqin", 1)
	// create("lqh", 1)
	// create("maxm", 0)
	users, err := getUsersByNameWithRawSql("lqh")
	if err != nil {
		panic(err)
	}
	fmt.Printf("users = %+v\n", users)
}

func delUser(id int) error {
	user := new(User)
	user.initWithId(id)
	user.del()
	return nil
}

func create(name string, sex int) error {
	user := new(User)
	user.Name = name
	user.Sex = sex
	user.Save()
	return nil
}

func up() error {
	user := new(User)
	user.initWithId(22)
	fmt.Printf("user.Id = %+v, user.Name=%+v \n", user.Id, user.Name)
	user.Sex = 3
	user.Save()
	return nil
}
