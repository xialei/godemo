package user

import "fmt"

// DemoUser : entry for user
func DemoUser() {
	u := new(User)
	u.SetName("Roger")
	u.SetAge(35)
	fmt.Println(u) // &{Roger 35}

	vip := new(VipUser)
	vip.SetName("Ken")
	vip.SetAge(35)
	vip.SetCs("Judy")
	fmt.Println(vip.GetCs(), "is customer service to", vip.GetName())
}

// User : 相比于面向过程，面向对象的目的是更好的封装，更好的可复用
type User struct {
	name string
	age int
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetAge(age int) {
	u.age = age
}

func (u *User) GetAge() int {
	return u.age
}

// Go语言组合式继承，实现代码复用的手段
type VipUser struct {
	cs string
	User // 匿名
}

func (v *VipUser) SetCs(cs string) {
	v.cs = cs
}

func (v *VipUser) GetCs() string {
	return v.cs
}