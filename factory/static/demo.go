package static

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

type ProductType int

const (
	p1 ProductType = iota // 0
	p2 // 1
)
// 实现简单工厂类
type ProductFactory struct {

}

func (p *ProductFactory) Create(productType ProductType) Product {
	switch productType {
	case p1:
		return &Product1{}
	case p2:
		return &Product2{}
	}
	return nil
}