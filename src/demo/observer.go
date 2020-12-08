package demo

import (
	"fmt"
	"time"
)

// Observer 一个对象的状态发生改变，所有依赖的对象都得到通知并自动更新
// 观察者接口定义
type Observer interface {
	Update(*Event)
}

// Event event driven
type Event struct {
	Msg string
}

// Subject object to be observed
type Subject interface {
	Regist(Observer)
	Deregist(Observer)
	Notify(*Event)
}

type ConcreateObserver struct {
	Id int
}

func (co *ConcreateObserver) Update(e *Event) {
	fmt.Printf("observer [%d] received msg: %s, \n", co.Id, e.Msg)
}

type ConcreateSubject struct {
	Observers map[Observer]struct{}
}

func (cs *ConcreateSubject) Regist(ob Observer) {
	cs.Observers[ob] = struct{}{}
}

func (cs *ConcreateSubject) Deregist(ob Observer) {
	delete(cs.Observers, ob)
}

func (cs *ConcreateSubject) Notify(e *Event) {
	for ob, _ := range cs.Observers {
		ob.Update(e)
	}
}

func DemoObserver() {
	cs := &ConcreateSubject{
		Observers: make(map[Observer]struct{}),
	}

	obs1 := &ConcreateObserver{1}
	obs2 := &ConcreateObserver{2}

	cs.Regist(obs1)
	cs.Regist(obs2)

	for i := 0; i < 5; i++ {
		e := &Event{fmt.Sprintf("msg [%d]", i)}
		cs.Notify(e)
		time.Sleep(time.Duration(1) * time.Second)
	}

}
