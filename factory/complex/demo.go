package complex

// 抽象的产品
type Product interface {
	SetName(name string)
	GetName() string
}
// 具体的产品1
type Product1 struct {
	name string
}

func (p *Product1)SetName(name string)  {
	p.name = name
}

func (p *Product1)GetName( ) string {
	return p.name
}
// 具体的产品2
type Product2 struct {
	name string
}
func (p *Product2)SetName(name string)  {
	p.name = name
}

func (p *Product2)GetName( ) string {
	return p.name
}

// 实现一个抽象的工厂
type Factory interface {
	Create() Product
}

// 实现一个具体的产品工厂1
type Product1Factory struct {

}

func (p1f *Product1Factory) Create() Product  {
	return &Product1{}
}


// 实现一个具体的产品工厂2
type Product2Factory struct {

}

func (p1f *Product2Factory) Create() Product  {
	return &Product2{}
}