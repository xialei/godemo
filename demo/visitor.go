package demo

import "fmt"

type IVisitor interface {
	Visit()
}

type CatVisitor struct {
}

func (v CatVisitor) Visit() {
	fmt.Printf("this is a cat\n")
}

type DogVisitor struct {
}

func (v DogVisitor) Visit() {
	fmt.Printf("this is a dog\n")
}

type IElement interface {
	Accept(visitor IVisitor)
}

type Element struct {
}

func (el Element) Accept(visitor IVisitor) {
	visitor.Visit()
}

type Pet struct {
	Element
}

func (e Pet) Print(visitor IVisitor) {
	fmt.Printf("Here comes a pet: ")
	e.Element.Accept(visitor)
}

// 平时我们定义完一个类之后，这个类所能执行的逻辑就是确定的了，但是我们经常会遇到一种场景: 根据外部环境更改这个类所能执行的行为。
// 而 访问者模式 就是在不更改这个类的前提下，更改这个类中方法所能执行的逻辑。
func DemoVisitor() {
	e := new(Element)
	e.Accept(new(CatVisitor))
	e.Accept(new(DogVisitor))

	m := new(Pet)
	m.Print(new(CatVisitor))
	m.Print(new(DogVisitor))
}
