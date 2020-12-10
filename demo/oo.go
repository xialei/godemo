package demo

import (
	"fmt"
	"reflect"
)

// DemoObjectOriented : used for demo
// GC 三色标记法 mark and sweep的改进
// 小对象复用，局部变量尽量少声明，多个小对象可以放入到结构体，方便GC扫描；少用string的"+"
func DemoObjectOriented() {
	a := 111
	b := "roger"
	c := []string{"hello", "world"}
	d := 3.58
	demoParam(a, b, c, &d)
	fmt.Println(a, b, c, d)

	demoStruct()

	demoInterface()

	demoAssert()

	demoReflect()
}

/**
 * slice, map, channel, interface, func 引用类型，其他都是值类型
**/
func demoParam(i int, s string, sli []string, m *float64) {
	i = 222
	s = "xia"
	sli[0] = "你好"
	*m = 3.1415926
}

// 结构体是值类型，要使用引用传递地址，可以使用结构体指针
func demoStruct() {
	var peo People
	peo = People{Name: "roger", Age: 35, Energy: 100}
	fmt.Println(peo)

	peo2 := People{Age: 20, Name: "xia"}
	fmt.Println(peo2)

	peo2.Name = "Jim"
	peo2.Age = 30
	fmt.Println(peo2.Name, peo2.Age)

	peo3 := new(People)
	peo3.Name = "Jim"
	peo3.Age = 30
	fmt.Println(*peo3) //peo3 is address, *peo3 is value {Jim 30}

	peo4 := peo3
	peo4.Name = "changeme"
	fmt.Println(peo3.Name) // changeme

	peo5 := &People{Name: "Kevin", Age: 18}
	fmt.Println(peo5) // &{Kevin 18}

	peo.rest(10)
	fmt.Println(peo.Energy) // not changed
	peo.play(5)
	fmt.Println(peo.Energy) // changed by reference

}

func demoInterface() {
	var p Practice = &People{Name: "Roger", Age: 35, Energy: 100} // 因为play实现了*People，所以这里要加&
	restAndPlay(p)

	var a Practice = &Animal{"Panda", 100}
	restAndPlay(a)
}

// restAndPlay 通过接口实现多态
func restAndPlay(p Practice) {
	p.rest(12)
	fmt.Println(p.energy())
	p.play(8)
	fmt.Println(p.energy())
}

// People 结构体
type People struct {
	Name   string
	Age    int
	Energy int
}

// Animal 另外一个struct
type Animal struct {
	Name   string
	Energy int
}

// Practice 结构体重写接口的所有方法，就认为结构体属于接口类型，可以把结构体变量赋值给接口变量，实现多态
type Practice interface {
	rest(hour int)
	play(hour int)
	energy() int
}

/**
 * 函数属于包，方法属于结构体，通过结构体变量调用，方法是属于特定类型的函数
 * func (变量名 结构体类型) 方法名(参数列表) 返回值列表 {
 * 		//方法体
 * }
**/
func (p People) rest(hour int) {
	fmt.Println(p.Name, "is resting now.")
	p.Energy += hour

}
func (p *People) play(hour int) {
	fmt.Println(p.Name, "is playing now.")
	p.Energy -= hour

}
func (p People) energy() int {
	return p.Energy
}

func (p *Animal) rest(day int) {
	fmt.Println(p.Name, "is resting now.")
	p.Energy += day

}
func (p *Animal) play(day int) {
	fmt.Println(p.Name, "is playing now.")
	p.Energy -= day
}
func (p Animal) energy() int {
	return p.Energy
}

// 断言用于判断类型
func demoAssert() {
	var i interface{} = 123
	result, ok := i.(int)
	fmt.Println(result, ok) // 123 true
	if ok {
		fmt.Println("good....")
	}

	res, notok := i.(bool)
	fmt.Println(res, notok) // false false
}

func demoReflect() {
	fmt.Println("===> in demoReflect")
	peo := People{Name: "Roger", Age: 35, Energy: 100}
	v := reflect.ValueOf(peo) //反射时获取peo的地址，Elem()获取指针指向地址的封装
	fmt.Println(v.NumField(), v.FieldByIndex([]int{0}))
	fieldName := "Name"
	fmt.Println(v.FieldByName(fieldName))

	elem := reflect.ValueOf(&peo).Elem()
	if elem.FieldByName(fieldName).CanSet() {
		elem.FieldByName(fieldName).SetString("Rogers")
		fmt.Println(elem.FieldByName(fieldName))
	}
	fmt.Println("===> end demoReflect")
}
