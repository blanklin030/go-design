package static

import (
	"fmt"
	"testing"
)

func TestProductFactory_Create(t *testing.T) {
	productFactory := &ProductFactory{}
	product1 := productFactory.Create(p1)
	product1.SetName("product1")
	fmt.Println(product1.GetName())

	product2 := productFactory.Create(p2)
	product2.SetName("product2")
	fmt.Println(product2.GetName())
}