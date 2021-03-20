package main

import "fmt"

type Observer interface {
	Update()
}

type Subjecter interface {
	Register(o Observer)
	Remove(o Observer)
	Notify()
}

type ConcreteObserver struct {
	Context string
	Subject Subjecter
}

type ConcreteSubject struct {
	Context   string
	Observers []Observer
}

func (s *ConcreteSubject) Register(o Observer) {
	s.Observers = append(s.Observers, o)
}

func (s *ConcreteSubject) Remove(o Observer) {
	for k, v := range s.Observers {
		if v == o {
			s.Observers = append(s.Observers[:k], s.Observers[k+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) SetValue(context string) {
	s.Context = context
	s.Notify()
}

func (s *ConcreteSubject) Notify() {
	for _, v := range s.Observers {
		v.Update()
	}
}

func (o *ConcreteObserver) Update() {
	o.Context = o.Subject.(*ConcreteSubject).Context
	fmt.Printf("now Context = %s\n", o.Context)
}

func NewConcreteObserver(s Subjecter) Observer {
	o := &ConcreteObserver{
		Subject: s,
	}
	s.Register(o)
	return o
}

func main() {
	s := new(ConcreteSubject)
	NewConcreteObserver(s)
	o2 := NewConcreteObserver(s)
	NewConcreteObserver(s)
	s.SetValue("123456")
	//有些人写成s.Remove(o2),难不成取消订阅还需要s去主动操作？无法理解。
	o2.(*ConcreteObserver).Subject.Remove(o2)
	s.SetValue("654321")
}
