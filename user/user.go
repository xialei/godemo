package user

import (
	"fmt"
)

// DemoUser : entry for user
func DemoUser() {
	u := new(User)
	u.SetFirstName("Roger")
	u.SetAge(35)
	fmt.Println(u) // &{Roger 35}

	vip := new(VipUser)
	vip.SetFirstName("Ken")
	vip.SetAge(35)
	vip.SetCs("Judy")
	fmt.Println(vip.GetCs(), "is customer service to", vip.GetFirstName())
}

// User : 相比于面向过程，面向对象的目的是更好的封装，更好的可复用
type User struct {
	id string
	// name string
	age       int
	firstName string
	lastName  string
}

// SetFirstName set
func (u *User) SetFirstName(firstName string) {
	u.firstName = firstName
}

// GetFirstName get
func (u *User) GetFirstName() string {
	return u.firstName
}

// SetAge set
func (u *User) SetAge(age int) {
	u.age = age
}

// GetAge get
func (u *User) GetAge() int {
	return u.age
}

//VipUser Go语言组合式继承，实现代码复用的手段
type VipUser struct {
	cs   string
	User // 匿名
}

// SetCs set
func (v *VipUser) SetCs(cs string) {
	v.cs = cs
}

// GetCs get
func (v *VipUser) GetCs() string {
	return v.cs
}
