package decorate

import "fmt"

//定义一个抽象的组件
type Component interface {
	Operate()
}
//实现一个具体的组件
type Component1 struct {
	
}

func (c *Component1) Operate()  {
	fmt.Println("c1 operate")
}

//定义一个抽象的装饰者
type Decorate interface {
	Component
	Do() //这是一个额外的方法
}
//实现一个具体的装饰者
type Decorate1 struct {
	c Component
}

func (d1 *Decorate1) Do()  {
	fmt.Println("发生了装饰行为")
}

func (d1 *Decorate1) Operate()  {
	d1.Do()
	d1.c.Operate()
}