package complex

import (
	"fmt"
	"testing"
)

func TestProduct1Factory_Create(t *testing.T) {
	product1Factory := &Product1Factory{}
	product1 := product1Factory.Create()
	product1.SetName("product1")
	fmt.Println(product1.GetName())

	product2Factory := &Product2Factory{}
	product2 := product2Factory.Create()
	product2.SetName("product2")
	fmt.Println(product2.GetName())
}
