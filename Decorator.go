package main

import "fmt"

type Coffee interface {
	cost() float32
}

type ConcreteCoffee struct {
	Price float32
}

func (p *ConcreteCoffee) cost() float32 {
	return p.Price
}

//继承+组合
type Decorator struct {
	Coffee
	Price float32
}

func (d *Decorator) cost() float32 {
	return d.Coffee.cost() + d.Price
}

func NewDecorator(c Coffee, price float32) Coffee {
	return &Decorator{
		Coffee: c, Price: price,
	}
}

func NewConcreteCoffee(price float32) Coffee {
	return &ConcreteCoffee{
		Price: price,
	}
}

func main() {
    //一杯咖啡用两份配料来装饰
	var c Coffee = NewConcreteCoffee(1.99)
    c = NewDecorator(c, 0.99)
    c = NewDecorator(c,1.5)
    fmt.Println(c.cost())
}
